package strand

import (
	"regexp"
	"strings"
)

// ToRNA takes in a DNA strand and transposes it to a strand of RNA. If an
// improper strand of DNA is passed in an empty string is returned.
func ToRNA(dna string) string {
	if match, _ := regexp.MatchString("[^CATG]", dna); match {
		return ""
	}

	rna := strings.Builder{}
	for _, nuc := range dna {
		switch nuc {
		case 'C':
			rna.WriteByte('G')
		case 'A':
			rna.WriteByte('U')
		case 'T':
			rna.WriteByte('A')
		case 'G':
			rna.WriteByte('C')
		}
	}
	return rna.String()
}
