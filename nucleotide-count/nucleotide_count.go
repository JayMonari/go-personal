package dna

import (
	"bytes"
	"errors"
	"regexp"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	if match, _ := regexp.MatchString("[^AGCT]", string(d)); match {
		return Histogram{}, errors.New("Invalid nucleotides in strand")
	}

	h := Histogram{}
	h['A'] = bytes.Count([]byte(d), []byte{'A'})
	h['G'] = bytes.Count([]byte(d), []byte{'G'})
	h['C'] = bytes.Count([]byte(d), []byte{'C'})
	h['T'] = bytes.Count([]byte(d), []byte{'T'})
	return h, nil
}
