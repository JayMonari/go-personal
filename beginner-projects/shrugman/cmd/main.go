package main

import (
	"bufio"
	"fmt"
	"hangman"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const alphabetLen = 26
const shrugman = `¯\_(ツ)_/¯`

var (
	rnd = rand.New(rand.NewSource(time.Now().UnixMilli()))
	rdr = bufio.NewReader(os.Stdin)
)

type set map[rune]struct{}

func newSet(str string) set {
	s := make(map[rune]struct{}, len(str))
	for _, r := range str {
		s[r] = struct{}{}
	}
	return s
}

func (s set) String() string {
	var sb strings.Builder
	for k := range s {
		sb.WriteRune(k)
		sb.WriteRune(' ')
	}
	// We cut the last ' ' character from the string `[:sb.Len()-1]`
	return sb.String()[:sb.Len()-1]
}

func main() {
	// Get random word
	word := hangman.Wordlist[rnd.Intn(len(hangman.Wordlist))]
	// Setup
	lives := 0
	letters := newSet(word)
	guessed := make(set, alphabetLen)
	display := make([]rune, len(word))
	for i := range word {
		display[i] = '_'
	}
	fmt.Println("runes guessed:", "\ndisplay:", string(display))
	for len(letters) > 0 {
		guess := getGuess(rdr)
		if _, found := guessed[guess]; found {
			fmt.Println("You already tried that!")
			continue
		}
		guessed[guess] = struct{}{}

		_, found := letters[guess]
		switch found {
		case true:
			for i, letter := range word {
				if letter == guess {
					display[i] = letter
				}
			}
			delete(letters, guess)
		case false:
			printShrugMan(&lives)
			if lives == len(shrugman) {
				fmt.Println("Looks like you lost!")
				return
			}
		}
		fmt.Println("runes guessed:", guessed, "\ndisplay:", string(display))
	}
	fmt.Println("🥳🎉🎉🎊  YOU WIN!  🎊🎉🎉🥳")
}

func printShrugMan(lives *int) {
	_, w := utf8.DecodeRuneInString(shrugman[*lives:])
	*lives += w

	var sb strings.Builder
	for i, r := range shrugman {
		if i == *lives {
			break
		}
		sb.WriteRune(r)
	}
	fmt.Println(sb.String())
}

func getGuess(r *bufio.Reader) rune {
	guess, _, err := r.ReadRune()
	if err != nil {
		fmt.Println("error trying to read rune:", err)
	}
	r.Discard(r.Buffered())
	guess = unicode.ToUpper(guess)
	return guess
}
