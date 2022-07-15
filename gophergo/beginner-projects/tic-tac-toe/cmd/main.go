package main

import (
	"fmt"
	"xno"
)

func main() {
	var pick int
	p1, p2 := xno.PlayerD, xno.PlayerM
	b := xno.Board{
		xno.Row{},
		xno.Row{},
		xno.Row{},
	}
	var turn bool
	for pick < 10 {
		fmt.Scanln(&pick)
		if pick < 1 || pick > 9 {
			fmt.Println("SYKE")
			continue
		}
		row := (pick - 1) / 3
		col := (pick - 1) % 3
		if b[row][col] != "" {
			fmt.Println("SYKE")
			continue
		}
		if turn {
			b[row][col] = p1.String()
		} else {
			b[row][col] = p2.String()
		}
		// TODO(jaymonari): Check for win
		// b.CheckWin()
		turn = !turn
		fmt.Println(b)
	}
}
