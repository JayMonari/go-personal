package guess

import (
	"fmt"
	"math/rand"
)

func Guess(n int) {
	randN := rand.Intn(n)
	var guess int
	for guess != randN {
		fmt.Println("Guess a number between 1 and", n)
		fmt.Scan(&guess)
		if guess < randN {
			fmt.Println("🤔 Looks like you're too low, guess higher ⬆️ ")
		} else if guess > randN {
			fmt.Println("🤔 Looks like you're too high, guess lower ⬇️ ")
		}
	}
	fmt.Println("🥳🥳Congratulations🎉🎉 you guessed the number correctly!")
}
