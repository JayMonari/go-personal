package main

import "fmt"

// When outside of a function this is known as a global variable. It is best
// practice to keep global variables to a minimum and make them `const`.
const madlib = `He decided to fake his %s to avoid %s.
If any %s asks you where you were, just say you were in %s.
He was %d%% into %sing with her until he understood that he couldn't %s.
`

// nWords is the number of words to fill in for the madlib.
const nWords = 7

func main() {
	var noun1 string
	noun2 := ""
	var noun3 string = ""
	var location = ""
	var percent int
	verb1 := ""
	verb2 := ""
	for i := 0; i < nWords; i++ {
		if i < 3 {
			fmt.Println("Noun:")
		} else if i == 3 {
			fmt.Println("Location:")
		} else if i == 4 {
			fmt.Println("Number:")
		} else {
			fmt.Println("Verb:")
		}
		switch i {
		case 0:
			fmt.Scanln(&noun1)
		case 1:
			fmt.Scanln(&noun2)
		case 2:
			fmt.Scanln(&noun3)
		case 3:
			fmt.Scanln(&location)
		case 4:
			fmt.Scanln(&percent)
		case 5:
			fmt.Scanln(&verb1)
		case 6:
			fmt.Scanln(&verb2)
		}
	}
	fmt.Printf(madlib, noun1, noun2, noun3, location, percent, verb1, verb2)
}
