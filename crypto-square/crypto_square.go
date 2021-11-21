package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode creates a secret message, known as a square code, from a string of
// text. The returned ciphered text will be lowercase without punctuation and
// grouped according to the length of the text.
func Encode(text string) string {
	n := normalize(text)
	// Provided by tleen's solution, because who, besides a genius, looks at this
	// and says, "Yes, this is obvious."
	cols := int(math.Ceil(math.Sqrt(float64(len(n)))))
	ciphered := make([]string, cols)
	for i, r := range n {
		ciphered[i%cols] += string(r)
	}

	if cols > 0 {
		padLength := len(ciphered[0])
		for i := 1; i < cols; i++ {
			ciphered[i] += strings.Repeat(" ", padLength-len(ciphered[i]))
		}
	}
	return strings.Join(ciphered, " ")
}

// normalize removes any punctuation, space, or symbols from a string and
// transforms all characters to lowercase.
func normalize(s string) string {
	n := strings.Builder{}
	for _, r := range strings.ToLower(s) {
		if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) {
			continue
		}
		n.WriteRune(r)
	}
	return n.String()
}
