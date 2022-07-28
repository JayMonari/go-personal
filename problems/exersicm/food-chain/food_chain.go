package foodchain

import (
	"fmt"
	"strings"
)

const (
	// outlierFmt is made for fly and horse
	outlierFmt = "I know an old lady who swallowed a %s.\n"
	startFmt   = "I know an old lady who swallowed a %s.\n%s\n"
	lyricFmt   = "She swallowed the %s to catch the %s.\n"
	spiderFmt  = "She swallowed the %s to catch the %s that wriggled and jiggled and tickled inside her.\n"
	ending     = "I don't know why she swallowed the fly. Perhaps she'll die."
)

// animals are all of the animals the old lady swallows.
var animals = [8]string{
	"fly",
	"spider",
	"bird",
	"cat",
	"dog",
	"goat",
	"cow",
	"horse",
}

// rhymes is an array of text that comes after the startFmt.
var rhymes = [8]string{
	"",
	"It wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!",
	"Imagine that, to swallow a cat!",
	"What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!",
	"I don't know how she swallowed a cow!",
	"She's dead, of course!",
}

// Verse returns the nth verse from the nursery them I know an Old Lady Who
// Swallowed a Fly.
func Verse(n int) string {
	sb := strings.Builder{}
	if n == 1 || n == 8 {
		sb.WriteString(fmt.Sprintf(outlierFmt, animals[n-1]))
	} else {
		sb.WriteString(fmt.Sprintf(startFmt, animals[n-1], rhymes[n-1]))
	}

	if n == 8 {
		sb.WriteString(rhymes[n-1])
		return sb.String()
	}

	for i := n - 1; i >= 1; i-- {
		if i == 2 {
			sb.WriteString(fmt.Sprintf(spiderFmt, animals[i], animals[i-1]))
		} else {
			sb.WriteString(fmt.Sprintf(lyricFmt, animals[i], animals[i-1]))
		}
	}

	sb.WriteString(ending)
	return sb.String()
}

// Verses returns the verses from start to end.
func Verses(start, end int) string {
	vs := []string{}
	for n := start; n <= end; n++ {
		vs = append(vs, Verse(n))
	}
	return strings.Join(vs, "\n\n")
}

// Song returns the entire nursery rhyme.
func Song() string {
	return Verses(1, 8)
}
