package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/danicat/simpleansi"
)

// Config holds the emoji configurations
type Config struct {
	Player   string `json:"player"`
	Ghost    string `json:"ghost"`
	Wall     string `json:"wall"`
	Dot      string `json:"dot"`
	Pill     string `json:"pill"`
	Death    string `json:"death"`
	Space    string `json:"space"`
	UseEmoji bool   `json:"use_emoji"`
}

var cfg Config

type sprite struct{ row, col int }

var player sprite

var score int
var numDots int
var lives = 1

var ghosts []*sprite
var move = map[int]string{
	0: "UP",
	1: "DOWN",
	2: "RIGHT",
	3: "LEFT",
}

func main() {
	initialize()
	defer cleanup()
	if err := loadMaze("level01.txt"); err != nil {
		log.Println("failed to load maze:", err)
		return
	}
	if err := loadConfig("config.json"); err != nil {
		log.Println("failed to load configuration:", err)
		return
	}
	input := make(chan string)
	go func(ch chan<- string) {
		for {
			input, err := readInput()
			if err != nil {
				log.Println("error reading input:", err)
				ch <- "ESC"
			}
			ch <- input
		}
	}(input)
	for {
		printScreen()
		select {
		case i := <-input:
			if i == "ESC" {
				lives = 0
			}
			movePlayer(i)
		default:
		}
		moveGhosts()
		for _, g := range ghosts {
			if player == *g {
				lives = 0
			}
		}
		if numDots == 0 || lives == 0 {
			if lives == 0 {
				moveCursor(player.row, player.col)
				fmt.Print(cfg.Death)
				moveCursor(len(maze)+2, 0)
			}
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

var maze []string

func loadMaze(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		maze = append(maze, s.Text())
	}
	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = sprite{row: row, col: col}
			case 'G':
				ghosts = append(ghosts, &sprite{row, col})
			case '.':
				numDots++
			}
		}
	}
	return nil
}

func printScreen() {
	simpleansi.ClearScreen()
	for _, line := range maze {
		for _, r := range line {
			switch r {
			case '#':
				fmt.Print(simpleansi.WithBlueBackground(cfg.Wall))
			case '.':
				fmt.Print(cfg.Dot)
			default:
				fmt.Print(cfg.Space)
			}
		}
		fmt.Println()
	}
	moveCursor(player.row, player.col)
	fmt.Print(cfg.Player)
	for _, g := range ghosts {
		moveCursor(g.row, g.col)
		fmt.Print(cfg.Ghost)
	}
	moveCursor(len(maze)+1, 0)
	fmt.Println("Score:", score, "\tLives:", lives)
}

func initialize() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin
	if err := cbTerm.Run(); err != nil {
		log.Fatalln("unable to activate cbreak mode:", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin
	if err := cookedTerm.Run(); err != nil {
		log.Fatalln("unable to restore cooked mode:", err)
	}
}

func readInput() (string, error) {
	buf := make([]byte, 100)
	cnt, err := os.Stdin.Read(buf)
	if err != nil {
		return "", err
	}
	if cnt == 1 && buf[0] == '' {
		return "ESC", nil
	} else if cnt >= 3 {
		if buf[0] == '' && buf[1] == '[' {
			switch buf[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}
	return "", nil
}

func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol
	switch dir {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(maze) {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(maze) - 1
		}
	}
	if maze[newRow][newCol] == '#' {
		newRow = oldRow
		newCol = oldCol
	}
	return
}

func movePlayer(dir string) {
	player.row, player.col = makeMove(player.row, player.col, dir)

	removeDot := func(r, c int) {
		maze[r] = maze[r][0:c] + " " + maze[r][c+1:]
	}
	switch maze[player.row][player.col] {
	case '.':
		numDots--
		score += 10
		removeDot(player.row, player.col)
	case 'X':
		score += 100
		removeDot(player.row, player.col)
	}
}

func drawDirection() string { return move[rand.Intn(4)] }

func moveGhosts() {
	for _, g := range ghosts {
		dir := drawDirection()
		g.row, g.col = makeMove(g.row, g.col, dir)
	}
}

func loadConfig(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = json.NewDecoder(f).Decode(&cfg); err != nil {
		return err
	}
	return nil
}

func moveCursor(row, col int) {
	if cfg.UseEmoji {
		simpleansi.MoveCursor(row, col*2)
	} else {
		simpleansi.MoveCursor(row, col)
	}
}
