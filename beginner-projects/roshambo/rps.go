package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// picks is the byte representation of rock, paper, scissors
var picks = []byte{'r', 'p', 's'}

// rdr is used to read in what the user has typed 'r', 'p', or 's'
var rdr = bufio.NewReader(os.Stdin)

// rnd is used to generate random guesses from the computer / AI
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	aiScore := 0
	userScore := 0
	fmt.Println("First to 3 wins!\nPick (r)ock, (p)aper, or (s)cissors")
	for {
		ai := picks[rnd.Intn(3)]
		user := input()
		fmt.Println("AI chose:", emojify(ai))
		fmt.Println("you chose:", emojify(user))
		switch {
		case !bytes.Contains(picks, []byte{user}):
			fmt.Println("That's not a choice! 😡")
		case user == ai:
			fmt.Println("👔 game!")
			continue
		case hasWon(user, ai):
			userScore++
			result(user)
			fmt.Println("You win!")
		default:
			aiScore++
			result(ai)
			fmt.Println("You lose!")
		}

		switch {
		case aiScore == 3:
			fmt.Println("🤖 I win human\n🤖 I am superior")
			return
		case userScore == 3:
			fmt.Println("🥳 You won the first to 3! 🎉🎉")
			return
		}
	}
}

// input grabs the input from the user from os.Stdin (standard in).
func input() (userPick byte) {
	pick, err := rdr.ReadByte()
	if err != nil {
		fmt.Println("How did you manage this?")
	}
	// If we don't discard the remaning the next byte we will read is a newline
	// character (\n) because we pressed the Enter/Return key.
	rdr.Discard(rdr.Buffered())
	return pick
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
func result(pick byte) {
	switch pick {
	case 'r':
		fmt.Println("🪨 beats ✂️ ")
	case 'p':
		fmt.Println("📰 covers 🪨")
	case 's':
		fmt.Println("✂️ cuts 📰")
	}
}
