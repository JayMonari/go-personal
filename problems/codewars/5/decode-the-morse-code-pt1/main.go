package main

import (
	"fmt"
	"strings"
)

var MORSE_CODE = map[string]string{
	"....":      "H",
	"..":        "I",
	"---..":     "8",
	"..--..":    "?",
	"----.":     "9",
	"-.":        "N",
	"--.-":      "Q",
	"...":       "S",
	"-.--":      "Y",
	".-..-.":    "\"",
	".----.":    "'",
	"-..-.":     "/",
	"-.-.":      "C",
	".--.":      "P",
	"..---":     "2",
	".....":     "5",
	".--.-.":    "@",
	".":         "E",
	".-.":       "R",
	"-.-.--":    "!",
	"...-..-":   "$",
	".-.-.-":    ".",
	"...---...": "SOS",
	".-.-.":     "+",
	"---...":    ":",
	"---":       "O",
	"..--.-":    "_",
	".-...":     "&",
	"-.--.-":    ")",
	"-.-.-.":    ";",
	".-":        "A",
	"-.-":       "K",
	"-":         "T",
	"...-":      "V",
	"..-.":      "F",
	"--..--":    ",",
	"-....":     "6",
	"-...-":     "=",
	".---":      "J",
	".--":       "W",
	"-..-":      "X",
	"...--":     "3",
	"--.":       "G",
	"-....-":    "-",
	".----":     "1",
	"-----":     "0",
	"....-":     "4",
	"--...":     "7",
	".-..":      "L",
	"--..":      "Z",
	"..-":       "U",
	"-.--.":     "(",
	"-...":      "B",
	"-..":       "D",
	"--":        "M",
}

// https://www.codewars.com/kata/54b724efac3d5402db00065e/train/go
func DecodeMorse(morseCode string) string {
	var useSpace bool
	var sb strings.Builder
	for _, s := range strings.Split(morseCode, " ") {
		if s == "" {
			if useSpace = !useSpace; useSpace {
				sb.WriteRune(' ')
			}
			continue
		}
		sb.WriteString(MORSE_CODE[s])
	}
	return strings.TrimSpace(sb.String())
}

func main() {
	fmt.Println(DecodeMorse(".... . -.--   .--- ..- -.. ."))
}
