package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/danicat/simpleansi"
)

type sprite struct{ row, col int }

var player sprite

func main() {
	initialize()
	defer cleanup()
	err := loadMaze("maze01.txt")
	if err != nil {
		log.Println("failed to load maze:", err)
		return
	}
	for {
		printScreen()
		input, err := readInput()
		if err != nil {
			log.Print("error reading input:", err)
			break
		}
		movePlayer(input)

		if input == "ESC" {
			break
		}
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
				fmt.Printf("%c", r)
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	simpleansi.MoveCursor(player.row, player.col)
	fmt.Print("P")
	simpleansi.MoveCursor(len(maze)+1, 0)
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
}
