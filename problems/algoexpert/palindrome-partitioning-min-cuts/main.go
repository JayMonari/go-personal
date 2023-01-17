package main

import "math"

func PalindromePartitioningMinCuts(s string) int {
	palindromes := make([][]bool, len(s))
	for i := range palindromes {
		palindromes[i] = make([]bool, len(s))
	}
	for i := range s {
		palindromes[i][i] = true
	}
	for length := 2; length < len(s)+1; length++ {
		for i := 0; i < len(s)-length+1; i++ {
			j := i + length - 1
			if length == 2 {
				palindromes[i][j] = (s[i] == s[j])
				continue
			}
			palindromes[i][j] = (s[i] == s[j] && palindromes[i+1][j-1])
		}
	}

	cuts := make([]int, len(s))
	for i := range cuts {
		cuts[i] = math.MaxInt32
	}
	for i := range s {
		if palindromes[0][i] {
			cuts[i] = 0
			continue
		}
		cuts[i] = cuts[i-1] + 1
		for j := 1; j < i; j++ {
			if palindromes[j][i] && cuts[j-1]+1 < cuts[i] {
				cuts[i] = cuts[j-1] + 1
			}
		}
	}
	return cuts[len(cuts)-1]
}

// Test Case 1
//
// {
//   "string": "noonabbad"
// }
//
// Test Case 2
//
// {
//   "string": "a"
// }
//
// Test Case 3
//
// {
//   "string": "abba"
// }
//
// Test Case 4
//
// {
//   "string": "abbba"
// }
//
// Test Case 5
//
// {
//   "string": "abb"
// }
//
// Test Case 6
//
// {
//   "string": "abbb"
// }
//
// Test Case 7
//
// {
//   "string": "abcbm"
// }
//
// Test Case 8
//
// {
//   "string": "ababbbabbababa"
// }
//
// Test Case 9
//
// {
//   "string": "abbbacecffgbgffab"
// }
//
// Test Case 10
//
// {
//   "string": "abcdefghijklmonpqrstuvwxyz"
// }
//
// Test Case 11
//
// {
//   "string": "abcdefghijklmracecaronpqrstuvwxyz"
// }
