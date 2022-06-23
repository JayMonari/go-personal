package wordcount

import (
	"regexp"
	"strings"
)

// Frequency gives the count of the time a word appears in the phrase of
// WordCount
type Frequency map[string]int

// WordCount counts the occurences of each word in the given phrase and will
// return a case insensitive, unordered Frequency.
func WordCount(phrase string) Frequency {
	freq := Frequency{}
	phrase = strings.ToLower(phrase)
	words := regexp.MustCompile("[a-z0-9']+").FindAllString(phrase, -1)
	for _, word := range words {
		if strings.HasPrefix(word, "'") && strings.HasSuffix(word, "'") {
			word = word[1:len(word)-1]
		}
		freq[word]++
	}
	return freq
}
