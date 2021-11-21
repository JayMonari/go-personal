package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate takes in a phrase and turns it into an acronym
func Abbreviate(s string) string {
	s = clean(s)
	words := strings.Split(s, " ")
	acronym := strings.Builder{}
	for _, word := range words {
		if word == "" {
			continue
		}
		acronym.WriteRune(unicode.ToUpper(rune(word[0])))
	}
	return acronym.String()
}

func clean(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	return s
}
