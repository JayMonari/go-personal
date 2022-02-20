package guess

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode"
)

func AIGuess(high int) {
	// This is how to make a random number generator in Go.
	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	// Reader rdr will allow us to read in a rune 'h', 'l', 'c' from user input.
	var rdr = bufio.NewReader(os.Stdin)

	fmt.Println("🤖I WILL NOW GUESS A NUMBER BETWEEN 1 AND", high)
	// Notice we declare all of these outside of the for loop scope.
	low := 1
	var feedback rune
	var err error
	for feedback != 'c' {
		if high-low == 0 {
			fmt.Println("🤖YOU'RE NUMBER IS", high)
			break
		}
		// To guess a range in Go, we can subtract low from high and add low back
		// high=25 low=5 -> 25-5=20 -> random number 0-19 -> 18 -> 18 + 5 = 23!
		guess := rnd.Intn(high-low) + low
		fmt.Println("IS", guess, "TOO (H)IGH, TOO (L)OW, OR (C)ORRECT?")

		// We read in the rune to provide feedback for the AI and check for errors.
		// Notice we do not use `:=` here.
		feedback, _, err = rdr.ReadRune()
		if err != nil {
			fmt.Println("There was an error reading your input!", err)
			continue
		}
		// We need to discard all other runes after the first or else feedback will
		// be wrong on the next loop.
		rdr.Discard(rdr.Buffered())

		switch unicode.ToUpper(feedback) {
		case 'H':
			high = guess - 1
		case 'L':
			low = guess + 1
		case 'C':
			fmt.Println("🤖YOU'RE NUMBER IS", guess)
		default:
			fmt.Println("🤖THAT IS NOT A CHOICE, HUMAN!")
		}
	}
	fmt.Println("🦾🤖I HAVE GUESSED YOUR NUMBER CORRECTLY, HUMAN. I AM THE SUPERIOR BEING.")
}
