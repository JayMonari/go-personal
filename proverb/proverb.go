package proverb

import "fmt"

const proverbFmt = "For want of a %s the %s was lost."
const endFmt = "And all for the want of a %s."

// Proverb returns a slice of strings in the format:
//
// For want of a words[0..n-1] the words[1..n] was lost.
// And all for the want of a words[0]
// where n is the length of words.
//
// If the length is zero an empty slice is returned.
func Proverb(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	lines := []string{}
	for i := 1; i < len(words); i++ {
		line := fmt.Sprintf(proverbFmt, words[i-1], words[i])
		lines = append(lines, line)
	}

	return append(lines, fmt.Sprintf(endFmt, words[0]))
}
