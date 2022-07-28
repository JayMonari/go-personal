package piglatin

import (
	"regexp"
	"strings"
)

var (
	// Rule 1
	vowel = regexp.MustCompile("^[aeiou]|^yt|^xr")
	// Rule 2 and 3
	consonant   = regexp.MustCompile("^([bcdfghjklmnpqrstuvwxyz]+)([aeiou].*)$")
	// Rule 4
	consonant_y = regexp.MustCompile("^([bcdfghjklmnpqrstuvwxyz]+)(y.*)$")
)

const (
	pigend = "ay"
)

// Sentence translates a text of english words into pig latin.
func Sentence(s string) string {
	latin := []string{}
	for _, word := range strings.Split(s, " ") {
		latin = append(latin, translate(word))
	}
	return strings.Join(latin, " ")
}

// translate takes a word s and translates it to pig latin. If a word cannot be
// translated the function panics.
func translate(s string) string {
	if match := vowel.MatchString(s); match {
		return s + pigend
	}

	if sl := consonant.FindStringSubmatch(s); len(sl) != 0 {
		return sl[2] + sl[1] + pigend
	} else if sl = consonant_y.FindStringSubmatch(s); len(sl) != 0 {
		return sl[2] + sl[1] + pigend
	}

	panic("this should be unreachable")
}
