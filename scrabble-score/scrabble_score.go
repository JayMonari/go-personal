package scrabble

import (
	"unicode"
)

// Score converts a word into a sum of points based on the value assigned to
// each letter in Scrabble.
func Score(word string) int {
	sum := 0
	for _, r := range word {
		sum += convertToValue(unicode.ToUpper(r))
	}
	return sum
}

func convertToValue(letter rune) int {
	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}
