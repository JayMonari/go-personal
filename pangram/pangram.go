package pangram

import "strings"

// IsPangram takes in a string of characters and returns true if the count of
// letters a through z is greater than zero, else false.
func IsPangram(phrase string) bool {
	alphabet := make(map[rune]bool, 26)
	for _, letter := range strings.ToLower(phrase) {
		if letter < 'a' || letter > 'z' {
			continue
		}
		alphabet[letter] = true
		if len(alphabet) == 26 {
			break
		}
	}
	return len(alphabet) == 26
}
