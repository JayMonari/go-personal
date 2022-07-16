package main

import (
	"fmt"
	"strings"
)

// https://www.codewars.com/kata/52223df9e8f98c7aa7000062/train/go
func Rot13(msg string) string {
	var sb strings.Builder
	for _, r := range msg {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') {
			sb.WriteRune(r)
			continue
		}
		switch {
		case r >= 'a' && r <= 'z':
			r = rotate(r, 13, 'a', 'z')
		case r >= 'A' && r <= 'Z':
			r = rotate(r, 13, 'A', 'Z')
		}
		sb.WriteRune(r)
	}
	return sb.String()
}

func rotate(r, amt, start, end rune) rune {
	cp := r + amt
	switch {
	case cp < start:
		return cp + end - start + 1
	case cp > end:
		return cp%end + start - 1
	default:
		return cp
	}
}

func main() {
	fmt.Println(Rot13("EBG13 rknzcyr."))
}
