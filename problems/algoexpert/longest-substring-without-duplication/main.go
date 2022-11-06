package main

type substring struct {
	left  int
	right int
}

func (ss substring) length() int { return ss.right - ss.left }

func LongestSubstringWithoutDuplication(s string) string {
	lastSeen := map[rune]int{}
	ss := substring{0, 1}
	start := 0
	for i, r := range s {
		if seen, found := lastSeen[r]; found && start < seen+1 {
			start = seen + 1
		}
		if ss.length() < i+1-start {
			ss = substring{start, i + 1}
		}
		lastSeen[r] = i
	}
	return s[ss.left:ss.right]
}

// Test Case 1
//
// {
//   "string": "clementisacap"
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
//   "string": "abc"
// }
//
// Test Case 4
//
// {
//   "string": "abcb"
// }
//
// Test Case 5
//
// {
//   "string": "abcdeabcdefc"
// }
//
// Test Case 6
//
// {
//   "string": "abccdeaabbcddef"
// }
//
// Test Case 7
//
// {
//   "string": "abacacacaaabacaaaeaaafa"
// }
//
// Test Case 8
//
// {
//   "string": "abcdabcef"
// }
//
// Test Case 9
//
// {
//   "string": "abcbde"
// }
//
// Test Case 10
//
// {
//   "string": "clementisanarm"
// }
