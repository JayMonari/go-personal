package main

import (
	"flag"
	"guess"
)

func main() {
	ai := flag.Bool("ai", false,
		"If you want the AI to guess your number add this flag.")
	n := flag.Int("range", 100, "The max number you want to guess to.")
	flag.Parse()
	if *ai {
		guess.AIGuess(*n)
	} else {
		guess.Guess(*n)
	}
}
