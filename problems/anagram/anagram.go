package anagram

import (
	"sort"
	"strings"
)

// Detect finds all anagrams of a subject, case insensitive. If there are none
// in the slice of words an empty string slice is returned.
func Detect(subject string, words []string) []string {
	anagrams := []string{}
	subject = strings.ToLower(subject)
	target := sortString(subject)
	for _, word := range words {
		lower := strings.ToLower(word)
		if lower == subject {
			continue
		}
		cand := sortString(lower)
		if cand == target {
			anagrams = append(anagrams, word)
		}
	}
	return anagrams
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
