package main

import (
	"blkjk/blackjack"
	"blkjk/deck"
	"fmt"
)

type basicAI struct{}

func (ai basicAI) Bet(_ bool) int {
	// noop
	return 1
}

func (ai basicAI) Play(hand []deck.Card, dealer deck.Card) blackjack.Move {
	panic("not implemented")
}

func (ai basicAI) Results(_ [][]deck.Card, _ []deck.Card) {
	// noop
}

func main() {
	opts := blackjack.Options{}
	game := blackjack.New(opts)
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
