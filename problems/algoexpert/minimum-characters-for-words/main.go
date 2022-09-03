package main

func MinimumCharactersForWords(words []string) []string {
	totalCounter := make(map[rune]int)
	for _, w := range words {
		for r, count := range countRunes(w) {
			if maxCount, found := totalCounter[r]; found {
				totalCounter[r] = max(maxCount, count)
				continue
			}
			totalCounter[r] = count
		}
	}

	var out []string
	for r, count := range totalCounter {
		for i := 0; i < count; i++ {
			out = append(out, string(r))
		}
	}
	return out
}

func countRunes(s string) (counter map[rune]int) {
	counter = make(map[rune]int)
	for _, r := range s {
		counter[r]++
	}
	return counter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test Case 1
// {
//   "words": ["this", "that", "did", "deed", "them!", "a"]
// }
// Test Case 2
// {
//   "words": ["a", "abc", "ab", "boo"]
// }
// Test Case 3
// {
//   "words": ["a"]
// }
// Test Case 4
// {
//   "words": ["abc", "ab", "b", "bac", "c"]
// }
// Test Case 5
// {
//   "words": ["!!!2", "234", "222", "432"]
// }
// Test Case 6
// {
//   "words": ["this", "that", "they", "them", "their", "there", "time", "is"]
// }
// Test Case 7
// {
//   "words": ["tim", "is", "great"]
// }
// Test Case 8
// {
//   "words": ["abc", "bavcc", "aaaa", "cde", "efg", "gead"]
// }
// Test Case 9
// {
//   "words": ["a", "a", "a"]
// }
// Test Case 10
// {
//   "words": ["them", "they", "that", "that", "yes", "yo", "no", "boo", "you", "okay", "too"]
// }
// Test Case 11
// {
//   "words": ["cta", "cat", "tca", "tac", "a", "c", "t"]
// }
// Test Case 12
// {
//   "words": ["my", "coding", "skills", "are", "great"]
// }
// Test Case 13
// {
//   "words": []
// }
// Test Case 14
// {
//   "words": ["168712hn3;nlsdjhahjdksaxa097918@#$RT%T^&*()_"]
// }
// Test Case 15
// {
//   "words": ["cat", "cAt", "tAc", "Act", "Cat"]
// }
// Test Case 16
// {
//   "words": ["Abc", "baVcc", "aaaa", "cdeE", "efg", "gead"]
// }
// Test Case 17
// {
//   "words": ["mississippi", "piper", "icing", "ice", "pickle", "piping", "pie", "pi", "sassy", "serpent", "python", "ascii", "sister", "mister"]
// }
