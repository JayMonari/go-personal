package main

func LevenshteinDistance(a, b string) int {
	dp := make([][]int, len(b)+1)
	for i := range dp {
		dp[i] = make([]int, len(a)+1)
		for j := range dp[i] {
			dp[i][j] = j
		}
	}
	for i := 1; i < len(b)+1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for i := 1; i < len(b)+1; i++ {
		for j := 1; j < len(a)+1; j++ {
			if b[i-1] == a[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			dp[i][j] = 1 + min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])
		}
	}
	return dp[len(b)][len(a)]
}

func min(args ...int) int {
	m := args[0]
	for _, n := range args {
		if n < m {
			m = n
		}
	}
	return m
}

// Test Case 1
// {
//   "str1": "abc",
//   "str2": "yabd"
// }
// Test Case 2
// {
//   "str1": "",
//   "str2": ""
// }
// Test Case 3
// {
//   "str1": "",
//   "str2": "abc"
// }
// Test Case 4
// {
//   "str1": "abc",
//   "str2": "abc"
// }
// Test Case 5
// {
//   "str1": "abc",
//   "str2": "abx"
// }
// Test Case 6
// {
//   "str1": "abc",
//   "str2": "abcx"
// }
// Test Case 7
// {
//   "str1": "abc",
//   "str2": "yabcx"
// }
// Test Case 8
// {
//   "str1": "algoexpert",
//   "str2": "algozexpert"
// }
// Test Case 9
// {
//   "str1": "abcdefghij",
//   "str2": "1234567890"
// }
// Test Case 10
// {
//   "str1": "abcdefghij",
//   "str2": "a234567890"
// }
// Test Case 11
// {
//   "str1": "biting",
//   "str2": "mitten"
// }
// Test Case 12
// {
//   "str1": "cereal",
//   "str2": "saturday"
// }
// Test Case 13
// {
//   "str1": "cereal",
//   "str2": "saturdzz"
// }
// Test Case 14
// {
//   "str1": "abbbbbbbbb",
//   "str2": "bbbbbbbbba"
// }
// Test Case 15
// {
//   "str1": "xabc",
//   "str2": "abcx"
// }
// Test Case 16
// {
//   "str1": "table",
//   "str2": "bal"
// }
// Test Case 17
// {
//   "str1": "gumbo",
//   "str2": "gambol"
// }
