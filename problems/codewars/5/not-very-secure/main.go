package main

import "regexp"

// https://www.codewars.com/kata/526dbd6c8c0eb53254000110/train/go
func alphanumeric(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(s)
}
