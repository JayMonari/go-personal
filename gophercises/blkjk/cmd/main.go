package main

import (
	"blkjk/blackjack"
	"fmt"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
