package main

func IsPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

// Test Case 1
// {
//   "string": "abcdcba"
// }
// Test Case 2
// {
//   "string": "a"
// }
// Test Case 3
// {
//   "string": "ab"
// }
// Test Case 4
// {
//   "string": "aba"
// }
// Test Case 5
// {
//   "string": "abb"
// }
// Test Case 6
// {
//   "string": "abba"
// }
// Test Case 7
// {
//   "string": "abcdefghhgfedcba"
// }
// Test Case 8
// {
//   "string": "abcdefghihgfedcba"
// }
// Test Case 9
// {
//   "string": "abcdefghihgfeddcba"
// }
