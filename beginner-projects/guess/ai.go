package guess

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func AIGuess(high int) {
	fmt.Println("🤖I will now guess a number between 1 and", high)
	// Reader r will allow us to read in a rune 'h', 'l', 'c' from user input.
	r := bufio.NewReader(os.Stdin)
	low := 1
	var feedback rune
	for feedback != 'c' {
		// To guess a range in Go, we can subtract low from high and add low back
		// high=25 low=5 -> 25-5=20 -> random number 0-19 -> 18 -> 18 + 5 = 23!
		guess := rand.Intn(high-low) + low
		fmt.Println("Is", guess, "too (h)igh, too (l)ow, or (c)orrect?")

		// We read in the rune to provide feedback for the AI.
		feedback, _, _ = r.ReadRune()
		// We need to discard the newline character here or else feedback will be
		// wrong on the next loop.
		r.Discard(r.Buffered())

		switch feedback {
		case 'h':
			high = guess - 1
		case 'l':
			low = guess + 1
		case 'c':
			break
		default:
			fmt.Println("🤖That is not a choice, human!")
		}
	}
	fmt.Println("🦾🤖I have guessed your number correctly, human. I am the superior being.")
}
