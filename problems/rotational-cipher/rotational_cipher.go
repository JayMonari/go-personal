package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher
func RotationalCipher(text string, shift int) string {
	rot := strings.Builder{}
	amt := rune(shift)
	for _, r := range text {
		var cp rune
		switch {
		case unicode.IsUpper(r):
			cp = rotate(r, amt, 'A', 'Z')
		case unicode.IsLower(r):
			cp = rotate(r, amt, 'a', 'z')
		default:
			cp = r
		}
		rot.WriteRune(cp)
	}
	return rot.String()
}

// rotate returns a rune that satisfies: start <= r + amt <= end
func rotate(r, amt, start, end rune) rune {
	cp := r + amt
	switch {
	case cp < start:
		return cp + end - start + 1
	case cp > end:
		return cp%end + start - 1
	default:
		return cp
	}
}
