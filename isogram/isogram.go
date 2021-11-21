package isogram

import "strings"

func IsIsogram(word string) bool {
	found := map[rune]bool{}
	for _, c := range strings.ToLower(word) {
		if c == '-' || c == ' ' {
			continue
		}
		if found[c] {
			return false
		}
		found[c] = true
	}
	return true
}
