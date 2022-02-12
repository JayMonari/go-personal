package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var aiPicks = [3]byte{'r', 'p', 's'}
var r = bufio.NewScanner(os.Stdin)
var rnd = rand.New(rand.NewSource(time.Now().UnixMilli()))

func main() {
	for {
		r.Scan()
		ai := aiPicks[rnd.Intn(3)]
		u := r.Text()[0]
		fmt.Println("AI chose:", emojify(ai))
		fmt.Println("you chose:", emojify(u))
		switch {
		case u != 'p' && u != 'r' && u != 's':
			fmt.Println("That's not a choice! 😡")
		case u == ai:
			fmt.Println("👔 game!")
			continue
		case hasWon(u, ai):
			fmt.Println("You win!")
			result(u)
		default:
			fmt.Println("You lose!")
			result(ai)
		}
	}
}

// emojify takes a byte and turns it into a single rune emoji string.
func emojify(b byte) string {
	switch b {
	case 'r':
		return "🪨"
	case 's':
		return "✂️ "
	case 'p':
		return "📰"
	default:
		return "❌"
	}
}

// hasWon returns true if the user won the rock paper scissors match.
func hasWon(u, ai byte) bool {
	return u == 'r' && ai == 's' || u == 's' && ai == 'p' || u == 'p' && ai == 'r'
}

// result prints out the final results of the match.
func result(p byte) {
	switch p {
	case 'r':
		fmt.Println("🪨 beats ✂️ ")
	case 'p':
		fmt.Println("📰 covers 🪨")
	case 's':
		fmt.Println("✂️ cuts 📰")
	}
}
