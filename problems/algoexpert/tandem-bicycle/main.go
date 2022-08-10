package main

import (
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TandemBicycle(speeds1, speeds2 []int, fastest bool) int {
	sort.Ints(speeds1)
	sort.Ints(speeds2)
	if !fastest {
		sort.Sort(sort.Reverse(sort.IntSlice(speeds1)))
	}
	total := 0
	for i := range speeds1 {
		total += max(speeds1[i], speeds2[len(speeds2)-i-1])
	}
	return total
}

// Test Case 1
// {
//   "redShirtSpeeds": [5, 5, 3, 9, 2],
//   "blueShirtSpeeds": [3, 6, 7, 2, 1],
//   "fastest": true
// }
// Test Case 2
// {
//   "redShirtSpeeds": [5, 5, 3, 9, 2],
//   "blueShirtSpeeds": [3, 6, 7, 2, 1],
//   "fastest": false
// }
// Test Case 3
// {
//   "redShirtSpeeds": [1, 2, 1, 9, 12, 3],
//   "blueShirtSpeeds": [3, 3, 4, 6, 1, 2],
//   "fastest": false
// }
// Test Case 4
// {
//   "redShirtSpeeds": [1, 2, 1, 9, 12, 3],
//   "blueShirtSpeeds": [3, 3, 4, 6, 1, 2],
//   "fastest": true
// }
// Test Case 5
// {
//   "redShirtSpeeds": [1, 2, 1, 9, 12, 3, 1, 54, 21, 231, 32, 1],
//   "blueShirtSpeeds": [3, 3, 4, 6, 1, 2, 5, 6, 34, 256, 123, 32],
//   "fastest": true
// }
// Test Case 6
// {
//   "redShirtSpeeds": [1, 2, 1, 9, 12, 3, 1, 54, 21, 231, 32, 1],
//   "blueShirtSpeeds": [3, 3, 4, 6, 1, 2, 5, 6, 34, 256, 123, 32],
//   "fastest": false
// }
// Test Case 7
// {
//   "redShirtSpeeds": [1],
//   "blueShirtSpeeds": [5],
//   "fastest": true
// }
// Test Case 8
// {
//   "redShirtSpeeds": [1],
//   "blueShirtSpeeds": [5],
//   "fastest": false
// }
// Test Case 9
// {
//   "redShirtSpeeds": [1, 1, 1, 1],
//   "blueShirtSpeeds": [1, 1, 1, 1],
//   "fastest": true
// }
// Test Case 10
// {
//   "redShirtSpeeds": [1, 1, 1, 1],
//   "blueShirtSpeeds": [1, 1, 1, 1],
//   "fastest": false
// }
// Test Case 11
// {
//   "redShirtSpeeds": [1, 1, 1, 1, 2, 2, 2, 2, 9, 11],
//   "blueShirtSpeeds": [1, 1, 1, 1, 3, 3, 3, 3, 5, 7],
//   "fastest": true
// }
// Test Case 12
// {
//   "redShirtSpeeds": [1, 1, 1, 1, 2, 2, 2, 2, 9, 11],
//   "blueShirtSpeeds": [1, 1, 1, 1, 3, 3, 3, 3, 5, 7],
//   "fastest": false
// }
// Test Case 13
// {
//   "redShirtSpeeds": [9, 8, 2, 2, 3, 5, 6],
//   "blueShirtSpeeds": [3, 4, 4, 1, 1, 8, 9],
//   "fastest": true
// }
// Test Case 14
// {
//   "redShirtSpeeds": [9, 8, 2, 2, 3, 5, 6],
//   "blueShirtSpeeds": [3, 4, 4, 1, 1, 8, 9],
//   "fastest": false
// }
// Test Case 15
// {
//   "redShirtSpeeds": [5, 4, 3, 2, 1],
//   "blueShirtSpeeds": [1, 2, 3, 4, 5],
//   "fastest": false
// }
// Test Case 16
// {
//   "redShirtSpeeds": [5, 4, 3, 2, 1],
//   "blueShirtSpeeds": [1, 2, 3, 4, 5],
//   "fastest": true
// }
// Test Case 17
// {
//   "redShirtSpeeds": [],
//   "blueShirtSpeeds": [],
//   "fastest": true
// }
