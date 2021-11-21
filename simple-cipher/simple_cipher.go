package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

type Encrypt []rune

// NewCaesar returns an Encrypt with Caeser's go to shift value.
func NewCaesar() Encrypt {
	return []rune{3}
}

// NewShift returns an Encrypt with the provided shift value.
func NewShift(s int) Encrypt {
	if s < -25 || s == 0 || s > 25 {
		return nil
	}
	return []rune{rune(s)}
}

// NewVigenere returns an Encrypt with a key to shift values by.
func NewVigenere(key string) Encrypt {
	if match, _ := regexp.MatchString("[^a-z]+", key); match {
		return nil
	} else if strings.Count(key, "a") == len(key) {
		return nil
	}
	shift := []rune{}
	for _, r := range key {
		shift = append(shift, r-'a')
	}
	return shift
}

// Encode returns a string ciphered by the key of e.
func (e Encrypt) Encode(text string) string {
	enc := strings.Builder{}
	i := 0
	for _, r := range strings.ToLower(text) {
		if !unicode.IsLetter(r) {
			continue
		}
		cp := rotate(r, e[i%len(e)])
		enc.WriteRune(cp)
		i++
	}
	return enc.String()
}

// Decode returns a string with the key provided in e.
func (e Encrypt) Decode(text string) string {
	dec := strings.Builder{}
	for i, r := range text {
		cp := rotate(r, -e[i%len(e)])
		dec.WriteRune(cp)
	}
	return dec.String()
}

// rotate shifts a rune by the given amount and leaves it within the lowercase
// letters a-z.
func rotate(r rune, shift rune) rune {
	cp := r + shift
	switch {
	case cp < 'a':
		return cp + 'z' - 'a' + 1
	case cp > 'z':
		return cp%'z' + 'a' - 1
	default:
		return cp
	}
}
