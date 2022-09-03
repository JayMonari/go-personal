package main

import (
	"sort"
)

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)
	n := len(queries) - 1
	total := 0
	for i, q := range queries[:n] {
		total += (n - i) * q
	}
	return total
}

// Test Case 1
// {
//   "queries": [3, 2, 1, 2, 6]
// }
// Test Case 2
// {
//   "queries": [2, 1, 1, 1]
// }
// Test Case 3
// {
//   "queries": [1, 2, 4, 5, 2, 1]
// }
// Test Case 4
// {
//   "queries": [25, 30, 2, 1]
// }
// Test Case 5
// {
//   "queries": [1, 1, 1, 1, 1]
// }
// Test Case 6
// {
//   "queries": [7, 9, 2, 3]
// }
// Test Case 7
// {
//   "queries": [4, 3, 1, 1, 3, 2, 1, 8]
// }
// Test Case 8
// {
//   "queries": [2]
// }
// Test Case 9
// {
//   "queries": [7]
// }
// Test Case 10
// {
//   "queries": [8, 9]
// }
// Test Case 11
// {
//   "queries": [1, 9]
// }
// Test Case 12
// {
//   "queries": [5, 4, 3, 2, 1]
// }
// Test Case 13
// {
//   "queries": [1, 2, 3, 4, 5]
// }
// Test Case 14
// {
//   "queries": [1, 1, 1, 4, 5, 6, 8, 1, 1, 2, 1]
// }
// Test Case 15
// {
//   "queries": [17, 4, 3]
// }
