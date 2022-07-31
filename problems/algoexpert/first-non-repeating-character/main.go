package main

func FirstNonRepeatingCharacter(str string) int {
	counter := map[rune]int{}
	for _, r := range str {
		counter[r]++
	}
	for i, r := range str {
		if counter[r] == 1 {
			return i
		}
	}
	return -1
}

// Test Case 1
// {
//   "string": "abcdcaf"
// }
// Test Case 2
// {
//   "string": "faadabcbbebdf"
// }
// Test Case 3
// {
//   "string": "a"
// }
// Test Case 4
// {
//   "string": "ab"
// }
// Test Case 5
// {
//   "string": "abc"
// }
// Test Case 6
// {
//   "string": "abac"
// }
// Test Case 7
// {
//   "string": "ababac"
// }
// Test Case 8
// {
//   "string": "ababacc"
// }
// Test Case 9
// {
//   "string": "lmnopqldsafmnopqsa"
// }
// Test Case 10
// {
//   "string": "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy"
// }
// Test Case 11
// {
//   "string": "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
// }
// Test Case 12
// {
//   "string": ""
// }
// Test Case 13
// {
//   "string": "ggyllaylacrhdzedddjsc"
// }
// Test Case 14
// {
//   "string": "aaaaaaaaaaaaaaaaaaaabbbbbbbbbbcccccccccccdddddddddddeeeeeeeeffghgh"
// }
// Test Case 15
// {
//   "string": "aabbccddeeff"
// }
