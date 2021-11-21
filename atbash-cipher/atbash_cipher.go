package atbash

import (
	"strings"
	"unicode"
)

const (
	plain  = "abcdefghijklmnopqrstuvwxyz"
	cipher = "zyxwvutsrqponmlkjihgfedcba"
)

// Atbash encodes text with a simple substitution cipher that relies on
// transposing all the letters in the alphabet such that the resulting alphabet
// is backwards. The return string is grouped in fixed lengths of five
// without punctuation and in lowercase. If the text to encode has other
// characters they will be included in the return string.
func Atbash(toEncode string) string {
	encoded := strings.Builder{}
	for _, r := range strings.ToLower(toEncode) {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			continue
		} else if !unicode.IsLower(r) {
			encoded.WriteRune(r)
			continue
		}
		pIdx := strings.IndexRune(plain, r)
		encoded.WriteByte(cipher[pIdx])
	}

	return splitN(encoded.String(), ' ', 5)
}

// splitN takes a string and splits it by n with the appropriate separator
func splitN(s string, separator rune, n int) string {
	rns := []rune(s)
	grouped := strings.Builder{}
	grouped.WriteRune(rns[0])

	for i := 1; i < len(rns); i++ {
		if i%n == 0 {
			grouped.WriteRune(separator)
		}
		grouped.WriteRune(rns[i])
	}
	return grouped.String()
}
