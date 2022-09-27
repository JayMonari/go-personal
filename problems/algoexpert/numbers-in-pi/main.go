package main

import "math"

func NumbersInPi(pi string, numbers []string) int {
	nums := map[string]struct{}{}
	for _, n := range numbers {
		nums[n] = struct{}{}
	}
	minSpaces := getMinSpaces(pi, nums, map[int]int{}, 0)
	if minSpaces == math.MaxInt32 {
		return -1
	}
	return minSpaces
}

func getMinSpaces(
	pi string, nums map[string]struct{}, cache map[int]int, idx int,
) int {
	if idx == len(pi) {
		return -1
	}
	if val, found := cache[idx]; found {
		return val
	}

	minSpaces := math.MaxInt32
	for i := idx; i < len(pi); i++ {
		prefix := pi[idx : i+1]
		if _, found := nums[prefix]; found {
			minSpacesInSuffix := getMinSpaces(pi, nums, cache, i+1)
			minSpaces = min(minSpaces, minSpacesInSuffix+1)
		}
	}
	cache[idx] = minSpaces
	return cache[idx]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Test Case 1
// {
//   "pi": "3141592653589793238462643383279",
//   "numbers": ["314159265358979323846", "26433", "8", "3279", "314159265", "35897932384626433832", "79"]
// }
// Test Case 2
// {
//   "pi": "3141592653589793238462643383279",
//   "numbers": ["314159265358979323846264338327", "9"]
// }
// Test Case 3
// {
//   "pi": "3141592653589793238462643383279",
//   "numbers": ["3", "314", "49", "9001", "15926535897", "14", "9323", "8462643383279", "4", "793"]
// }
// Test Case 4
// {
//   "pi": "3141592653589793238462643383279",
//   "numbers": ["3141592653589793238462643383279"]
// }
// Test Case 5
// {
//   "pi": "3141592653589793238462643383279",
//   "numbers": ["3141", "1512", "159", "793", "12412451", "8462643383279"]
// }
// Test Case 6
// {
//   "pi": "3141592653589793238462643383279",
//   "numbers": ["314159265358979323846", "327", "26433", "8", "3279", "9", "314159265", "35897932384626433832", "79"]
// }
// Test Case 7
// {
//   "pi": "3141592653589793238462643383279",
//   "numbers": ["141592653589793238462643383279", "314159265358979323846", "327", "26433", "8", "3279", "9", "314159265", "35897932384626433832", "79", "3"]
// }
// Test Case 8
// {
//   "pi": "3141592653589793238462643383279",
//   "numbers": ["3", "1", "4", "592", "65", "55", "35", "8", "9793", "2384626", "83279"]
// }
// Test Case 9
// {
//   "pi": "3141592653589793238462643383279",
//   "numbers": ["3", "1", "4", "592", "65", "55", "35", "8", "9793", "2384626", "383279"]
// }
// Test Case 10
// {
//   "pi": "3141592653589793238462643383279",
//   "numbers": ["3", "141", "592", "65", "55", "35", "8", "9793", "2384626", "383279"]
// }
// Test Case 11
// {
//   "pi": "3141592653589793238462643383279",
//   "numbers": ["3", "141", "592", "65", "55", "35", "8", "9793", "23846264", "383279"]
// }
// Test Case 12
// {
//   "pi": "3141592653589793238462643383279",
//   "numbers": ["3", "141", "592", "65", "55", "35", "8", "9793", "23846264", "3832798"]
// }
