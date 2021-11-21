package isbn

import (
	"regexp"
	"strings"
)

// IsValidISBN validates a given ISBN string. The input can have '-' characters
// and the number of digits must add up to 10 or 9 + the 'X' checksum character
// at the end of the string. If these conditions are not met, a valid ISBN will
// be treated as invalid.
func IsValidISBN(cand string) bool {
	cand = regexp.MustCompile("[^0-9X]").ReplaceAllString(cand, "")
  xCheck := strings.IndexByte(cand, 'X')
	if len(cand) != 10 || !(xCheck == -1 || xCheck == 9) {
		return false
	}

	sum := 0
	for i, r := range cand {
		if r == 'X' {
			sum += 10
			break
		}
		sum += int(r - '0') * (10 - i)
	}
	return sum%11 == 0
}
