package main

import (
	"blkjk/blackjack"
	"fmt"
)

func main() {
	opts := blackjack.Options{}
	game := blackjack.New(opts)
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
