package main

import "math"

func MinNumberOfCoinsForChange(n int, coins []int) int {
	mins := make([]int, n+1)
	for i := range mins {
		mins[i] = math.MaxInt32
	}
	mins[0] = 0
	for _, c := range coins {
		for amt := 0; amt < n+1; amt++ {
			if c <= amt {
				mins[amt] = min(mins[amt-c]+1, mins[amt])
			}
		}
	}
	if mins[len(mins)-1] == math.MaxInt32 {
		return -1
	}
	return mins[len(mins)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Test Case 1
// {
//   "n": 7,
//   "denoms": [1, 5, 10]
// }
// Test Case 2
// {
//   "n": 7,
//   "denoms": [1, 10, 5]
// }
// Test Case 3
// {
//   "n": 7,
//   "denoms": [10, 1, 5]
// }
// Test Case 4
// {
//   "n": 0,
//   "denoms": [1, 2, 3]
// }
// Test Case 5
// {
//   "n": 3,
//   "denoms": [2, 1]
// }
// Test Case 6
// {
//   "n": 4,
//   "denoms": [1, 5, 10]
// }
// Test Case 7
// {
//   "n": 10,
//   "denoms": [1, 5, 10]
// }
// Test Case 8
// {
//   "n": 11,
//   "denoms": [1, 5, 10]
// }
// Test Case 9
// {
//   "n": 24,
//   "denoms": [1, 5, 10]
// }
// Test Case 10
// {
//   "n": 25,
//   "denoms": [1, 5, 10]
// }
// Test Case 11
// {
//   "n": 7,
//   "denoms": [2, 4]
// }
// Test Case 12
// {
//   "n": 7,
//   "denoms": [3, 7]
// }
// Test Case 13
// {
//   "n": 9,
//   "denoms": [3, 5]
// }
// Test Case 14
// {
//   "n": 9,
//   "denoms": [3, 4, 5]
// }
// Test Case 15
// {
//   "n": 135,
//   "denoms": [39, 45, 130, 40, 4, 1]
// }
// Test Case 16
// {
//   "n": 135,
//   "denoms": [39, 45, 130, 40, 4, 1, 60, 75]
// }
// Test Case 17
// {
//   "n": 10,
//   "denoms": [1, 3, 4]
// }
