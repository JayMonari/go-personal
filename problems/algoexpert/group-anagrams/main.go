package main

import (
	"sort"
)

func GroupAnagrams(words []string) [][]string {
	sorted := make([]string, len(words))
	for i := range sorted {
		word := []byte(words[i])
		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		sorted[i] = string(word)
	}

	groups := map[string][]string{}
	for i, w := range words {
		groups[sorted[i]] = append(groups[sorted[i]], w)
	}

	anagrams := make([][]string, 0, len(groups))
	for _, g := range groups {
		anagrams = append(anagrams, g)
	}
	return anagrams
}

// Test Case 1
// {
//   "words": ["yo", "act", "flop", "tac", "foo", "cat", "oy", "olfp"]
// }
// Test Case 2
// {
//   "words": []
// }
// Test Case 3
// {
//   "words": ["test"]
// }
// Test Case 4
// {
//   "words": ["abc", "dabd", "bca", "cab", "ddba"]
// }
// Test Case 5
// {
//   "words": ["abc", "cba", "bca"]
// }
// Test Case 6
// {
//   "words": ["zxc", "asd", "weq", "sda", "qwe", "xcz"]
// }
// Test Case 7
// {
//   "words": ["cinema", "a", "flop", "iceman", "meacyne", "lofp", "olfp"]
// }
// Test Case 8
// {
//   "words": ["abc", "abe", "abf", "abg"]
// }
// Test Case 9
// {
//   "words": ["aaa", "a"]
// }
// Test Case 10
// {
//   "words": ["bob", "boo"]
// }
// Test Case 11
// {
//   "words": ["ill", "duh"]
// }
// Test Case 12
// {
//   "words": ["yo", "oy", "zn"]
// }
// Test Case 13
// {
//   "words": ["yyo", "yo"]
// }
// Test Case 14
// {
//   "words": ["aca", "bba"]
// }
