package main

import "fmt"

type state int8

const (
	inProgress state = iota - 1
	catsGame
	xWon
	oWon
)

type player uint8

const (
	empty player = iota
	x
	o
)

// https://www.codewars.com/kata/525caa5c1bf619d28c000335/train/go
func IsSolved(board [3][3]int) int {
	crownWinner := func(n player) state {
		if n == x {
			return xWon
		}
		return oWon
	}
	r1, r2, r3 := board[0], board[1], board[2]
	switch {
	case r1[0] == r1[1] && r1[1] == r1[2] && r1[0] != int(empty): // left row
		fmt.Println("r1:", r1)
		return int(crownWinner(player(r1[0])))
	case r2[0] == r2[1] && r2[1] == r2[2] && r2[0] != int(empty): // middle row
		fmt.Println("r2:", r2)
		return int(crownWinner(player(r2[0])))
	case r3[0] == r3[1] && r3[1] == r3[2] && r3[0] != int(empty): // right row
		fmt.Println("r3:", r3)
		return int(crownWinner(player(r3[0])))
	case r1[0] == r2[0] && r2[0] == r3[0] && r1[0] != int(empty): // left column
		return int(crownWinner(player(r1[0])))
	case r1[1] == r2[1] && r2[1] == r3[1] && r1[1] != int(empty): // middle column
		return int(crownWinner(player(r1[1])))
	case r1[2] == r2[2] && r2[2] == r3[2] && r1[2] != int(empty): // right column
		return int(crownWinner(player(r1[2])))
	case r1[0] == r2[1] && r2[1] == r3[2] && r1[0] != int(empty): // left diagonal
		return int(crownWinner(player(r1[0])))
	case r1[2] == r2[1] && r2[1] == r3[0] && r1[2] != int(empty): // right diagonal
		return int(crownWinner(player(r1[2])))
	case noZeroes(board):
		return int(catsGame)
	}
	return int(inProgress)
}

func noZeroes(b [3][3]int) bool {
	for _, r := range b {
		for _, n := range r {
			if n == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(IsSolved([3][3]int{{1, 1, 1}, {0, 2, 2}, {0, 0, 0}}))
}
