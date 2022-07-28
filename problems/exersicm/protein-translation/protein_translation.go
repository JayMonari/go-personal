package protein

import (
	"errors"
	"regexp"
)

var (
	// ErrStop is the generated error when the base given is a stop codon
	ErrStop = errors.New("Given base was a stop codon")
	// ErrInvalidBase is the generated error when the base given is invalid
	ErrInvalidBase = errors.New("Given an invalid codon")
)

// FromCodon uses a codon value and maps it to its corresponding protein.
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA takes in a strand of RNA and maps each codon to its corresponding
// protein and returns them. If the passed in codon is a stop codon the parsing
// will terminate early and give the currently translated proteins. If an
// invalid base is found the proteins, along with an error will be returned.
func FromRNA(rna string) ([]string, error) {
	proteins := []string{}
	for _, codon := range regexp.MustCompile("...").FindAllString(rna, -1) {
		protein, err := FromCodon(codon)
		if err == ErrStop {
			return proteins, nil
		} else if err == ErrInvalidBase {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
