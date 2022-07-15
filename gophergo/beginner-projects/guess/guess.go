package guess

import (
	"fmt"
	"math/rand"
	"time"
)

func Guess(n int) {
	// This is how to make a random number generator in Go.
	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	randN := rnd.Intn(n)
	var guess int
	for guess != randN {
		fmt.Println("Guess a number between 1 and", n)
		fmt.Scanln(&guess)
		if guess < randN {
			fmt.Println("🤔 Looks like you're too low, guess higher ⬆️ ")
		} else if guess > randN {
			fmt.Println("🤔 Looks like you're too high, guess lower ⬇️ ")
		}
	}
	fmt.Println("🥳🥳Congratulations🎉🎉 you guessed the number correctly!")
}
