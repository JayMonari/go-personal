package main

import (
	"regexp"
)

// https://www.codewars.com/kata/51fc12de24a9d8cb0e000001/train/go
func ValidISBN10(isbn string) bool {
	switch {
	case !regexp.MustCompile(`\d{9}\d|[xX]`).MatchString(isbn):
		return false
	case len(isbn) != 10:
		return false
	}

	sum := 0
	for i, r := range isbn {
		if r == 'X' || r == 'x' {
			sum += 100
			break
		}
		sum += int(r-'0') * (i + 1)
	}
	return sum%11 == 0
}
