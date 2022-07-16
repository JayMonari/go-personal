package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://www.codewars.com/kata/525f50e3b73515a6db000b83/train/go
func CreatePhoneNumber(numbers [10]uint) string {
	var sb strings.Builder
	for _, n := range numbers {
		sb.WriteString(strconv.Itoa(int(n)))
	}
	s := sb.String()
	return fmt.Sprintf("(%s) %s-%s", s[:3], s[3:6], s[6:])
}
