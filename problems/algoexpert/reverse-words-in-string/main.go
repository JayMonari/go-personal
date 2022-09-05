package main

import (
	"strings"
)

func isWhitespace(b byte) bool { return b == ' ' || b == '\t' || b == '\n' }

func ReverseWordsInString(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		if b := s[i]; isWhitespace(b) {
			sb.WriteByte(b)
			continue
		}
		startIdx := i
		for ; startIdx >= 0 && !isWhitespace(s[startIdx]); startIdx-- {
		}
		switch {
		case startIdx == -1:
			startIdx = 0
			sb.WriteString(s[startIdx : i+1])
		case isWhitespace(s[startIdx]):
			startIdx++
			sb.WriteString(s[startIdx : i+1])
		}
		i = startIdx
	}
	return sb.String()
}

// Test Case 1
// {
//   "string": "AlgoExpert is the best!"
// }
// Test Case 2
// {
//   "string": "Reverse These Words"
// }
// Test Case 3
// {
//   "string": "..H,, hello 678"
// }
// Test Case 4
// {
//   "string": "this this words this this this words this"
// }
// Test Case 5
// {
//   "string": "1 12 23 34 56"
// }
// Test Case 6
// {
//   "string": "APPLE PEAR PLUM ORANGE"
// }
// Test Case 7
// {
//   "string": "this-is-one-word"
// }
// Test Case 8
// {
//   "string": "a"
// }
// Test Case 9
// {
//   "string": "ab"
// }
// Test Case 10
// {
//   "string": ""
// }
// Test Case 11
// {
//   "string": "algoexpert is the best platform to use to prepare for coding interviews!"
// }
// Test Case 12
// {
//   "string": "words, separated, by, commas"
// }
// Test Case 13
// {
//   "string": "this      string     has a     lot of   whitespace"
// }
// Test Case 14
// {
//   "string": "a ab a"
// }
// Test Case 15
// {
//   "string": "test        "
// }
// Test Case 16
// {
//   "string": " "
// }
