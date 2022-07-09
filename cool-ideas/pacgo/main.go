package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/danicat/simpleansi"
)

// config holds the emoji configurations
type config struct {
	Player           string        `json:"player"`
	Ghost            string        `json:"ghost"`
	GhostBlue        string        `json:"ghost_blue"`
	Wall             string        `json:"wall"`
	Dot              string        `json:"dot"`
	Pill             string        `json:"pill"`
	Death            string        `json:"death"`
	Space            string        `json:"space"`
	UseEmoji         bool          `json:"use_emoji"`
	PillDurationSecs time.Duration `json:"pill_duration_secs"`
}

type point struct{ row, col int }
type sprite struct {
	point              point
	startRow, startCol int
}
type GhostStatus uint8

const (
	GhostStatusNormal GhostStatus = iota
	GhostStatusBlue
)

type ghost struct {
	sprite
	status GhostStatus
}

type direction uint8

const (
	noop direction = iota
	esc
	up
	down
	left
	right
)

var (
	configFile = flag.String("config-file", "config.json", "path to custom configuration file")
	levelFile  = flag.String("level-file", "level01.txt", "path to a custom level file")

	cfg config

	maze []string

	player sprite

	score   int
	numDots int
	lives   = 3

	ghosts   []*ghost
	ghostsMu sync.RWMutex
	move     = map[int]direction{
		0: up,
		1: down,
		2: right,
		3: left,
	}
	pillTimer *time.Timer
	pillMu    sync.Mutex
)

func main() {
	flag.Parse()

	initialize()
	defer cleanup()
	if err := loadMaze(*levelFile); err != nil {
		log.Println("failed to load maze:", err)
		return
	}
	if err := loadConfig(*configFile); err != nil {
		log.Println("failed to load configuration:", err)
		return
	}
	input := make(chan direction)
	go func(ch chan<- direction) {
		for {
			input, err := readInput()
			if err != nil {
				log.Println("error reading input:", err)
				ch <- esc
			}
			ch <- input
		}
	}(input)
	for {
		printScreen()
		select {
		case i := <-input:
			if i == esc {
				lives = 0
			}
			movePlayer(i)
		default:
		}
		moveGhosts()
		for _, g := range ghosts {
			if player.point == g.point {
				lives--
				if lives != 0 {
					moveCursor(player.point.row, player.point.col)
					fmt.Print(cfg.Death)
					moveCursor(len(maze)+2, 0)
					time.Sleep(1 * time.Second)
					player.point.row, player.point.col = player.startRow, player.startCol
				}
			}
		}
		if numDots == 0 || lives == 0 {
			if lives == 0 {
				moveCursor(player.point.row, player.point.col)
				fmt.Print(cfg.Death)
				moveCursor(len(maze)+2, 0)
			}
			break
		}
		time.Sleep(250 * time.Millisecond)
	}
}

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
				player = sprite{
					point:    point{row: row, col: col},
					startRow: row, startCol: col,
				}
			case 'G':
				ghosts = append(ghosts, &ghost{
					sprite: sprite{
						point:    point{row: row, col: col},
						startRow: row, startCol: col,
					},
					status: GhostStatusNormal,
				})
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
			case 'X':
				fmt.Print(cfg.Pill)
			default:
				fmt.Print(cfg.Space)
			}
		}
		fmt.Println()
	}
	moveCursor(player.point.row, player.point.col)
	fmt.Print(cfg.Player)

	ghostsMu.RLock()
	for _, g := range ghosts {
		moveCursor(g.point.row, g.point.col)
		if g.status == GhostStatusNormal {
			fmt.Print(cfg.Ghost)
		} else {
			fmt.Print(cfg.GhostBlue)
		}
	}
	ghostsMu.RUnlock()

	var livesRemaning string
	if cfg.UseEmoji {
		livesRemaning = getLivesAsEmoji()
	} else {
		strconv.Itoa(lives)
	}
	moveCursor(len(maze)+1, 0)
	fmt.Println("Score:", score, "\tLives:", livesRemaning)
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

func readInput() (direction, error) {
	buf := make([]byte, 100)
	cnt, err := os.Stdin.Read(buf)
	if err != nil {
		return noop, err
	}
	if cnt == 1 && buf[0] == '' {
		return esc, nil
	} else if cnt >= 3 {
		if buf[0] == '' && buf[1] == '[' {
			switch buf[2] {
			case 'A':
				return up, nil
			case 'B':
				return down, nil
			case 'C':
				return right, nil
			case 'D':
				return left, nil
			}
		}
	}
	return noop, nil
}

func makeMove(oldRow, oldCol int, d direction) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol
	switch d {
	case up:
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case down:
		newRow = newRow + 1
		if newRow == len(maze) {
			newRow = 0
		}
	case right:
		newCol = newCol + 1
		if newCol == len(maze[0]) {
			newCol = 0
		}
	case left:
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

func movePlayer(d direction) {
	player.point.row, player.point.col = makeMove(player.point.row, player.point.col, d)

	removeDot := func(r, c int) {
		maze[r] = maze[r][0:c] + " " + maze[r][c+1:]
	}
	switch maze[player.point.row][player.point.col] {
	case '.':
		numDots--
		score += 10
		removeDot(player.point.row, player.point.col)
	case 'X':
		score += 100
		removeDot(player.point.row, player.point.col)
		go processPill()
	}
}

func drawDirection() direction { return move[rand.Intn(4)] }

func moveGhosts() {
	for _, g := range ghosts {
		dir := drawDirection()
		g.point.row, g.point.col = makeMove(g.point.row, g.point.col, dir)
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

func getLivesAsEmoji() string {
	sb := strings.Builder{}
	for i := lives; i > 0; i-- {
		sb.WriteString(cfg.Player)
	}
	return sb.String()
}

func processPill() {
	pillMu.Lock()
	updateGhosts(ghosts, GhostStatusBlue)
	if pillTimer != nil {
		pillTimer.Stop()
	}
	pillTimer = time.NewTimer(time.Second * cfg.PillDurationSecs)
	pillMu.Unlock()
	<-pillTimer.C
	pillMu.Lock()
	pillTimer.Stop()
	updateGhosts(ghosts, GhostStatusNormal)
	pillMu.Unlock()
}

func updateGhosts(gg []*ghost, s GhostStatus) {
	ghostsMu.Lock()
	defer ghostsMu.Unlock()
	for _, g := range gg {
		g.status = s
	}
}
