package main

type substring struct{ left, right int }

func (ss substring) length() int { return ss.right - ss.left }

func LongestPalindromicSubstring(s string) string {
	ss := substring{0, 1}
	for i := 1; i < len(s); i++ {
		odd := findPalindrome(s, i-1, i+1)
		longest := findPalindrome(s, i-1, i)
		if odd.length() > longest.length() {
			longest = odd
		}
		if longest.length() > ss.length() {
			ss = longest
		}
	}
	return s[ss.left:ss.right]
}

func findPalindrome(s string, leftIdx, rightIdx int) substring {
	for leftIdx >= 0 && rightIdx < len(s) {
		if s[leftIdx] != s[rightIdx] {
			break
		}
		leftIdx--
		rightIdx++
	}
	return substring{leftIdx + 1, rightIdx}
}

// Test Case 1
// {
//   "string": "abaxyzzyxf"
// }
// Test Case 2
// {
//   "string": "a"
// }
// Test Case 3
// {
//   "string": "it's highnoon"
// }
// Test Case 4
// {
//   "string": "noon high it is"
// }
// Test Case 5
// {
//   "string": "abccbait's highnoon"
// }
// Test Case 6
// {
//   "string": "abcdefgfedcbazzzzzzzzzzzzzzzzzzzz"
// }
// Test Case 7
// {
//   "string": "abcdefgfedcba"
// }
// Test Case 8
// {
//   "string": "abcdefghfedcbaa"
// }
// Test Case 9
// {
//   "string": "abcdefggfedcba"
// }
// Test Case 10
// {
//   "string": "zzzzzzz2345abbbba5432zzbbababa"
// }
// Test Case 11
// {
//   "string": "z234a5abbbba54a32z"
// }
// Test Case 12
// {
//   "string": "z234a5abbba54a32z"
// }
// Test Case 13
// {
//   "string": "ab12365456321bb"
// }
// Test Case 14
// {
//   "string": "aca"
// }
