package main

import "fmt"

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
			fmt.Scan(&noun1)
		case 1:
			fmt.Scan(&noun2)
		case 2:
			fmt.Scan(&noun3)
		case 3:
			fmt.Scan(&location)
		case 4:
			fmt.Scan(&percent)
		case 5:
			fmt.Scan(&verb1)
		case 6:
			fmt.Scan(&verb2)
		}
	}
	fmt.Printf(madlib, noun1, noun2, noun3, location, percent, verb1, verb2)
}
