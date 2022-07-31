package main

import "math"

func FindThreeLargestNumbers(array []int) []int {
	large3 := make([]int, 3)
	for i := range large3 {
		large3[i] = math.MinInt
	}
	for _, n := range array {
		switch {
		case large3[2] < n:
			large3[0], large3[1], large3[2] = large3[1], large3[2], n
		case large3[1] < n:
			large3[0], large3[1] = large3[1], n
		case large3[0] < n:
			large3[0] = n
		}
	}
	return large3
}

// Test Case 1
// {
//   "array": [141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7]
// }
// Test Case 2
// {
//   "array": [55, 7, 8]
// }
// Test Case 3
// {
//   "array": [55, 43, 11, 3, -3, 10]
// }
// Test Case 4
// {
//   "array": [7, 8, 3, 11, 43, 55]
// }
// Test Case 5
// {
//   "array": [55, 7, 8, 3, 43, 11]
// }
// Test Case 6
// {
//   "array": [7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7]
// }
// Test Case 7
// {
//   "array": [7, 7, 7, 7, 7, 7, 8, 7, 7, 7, 7]
// }
// Test Case 8
// {
//   "array": [-1, -2, -3, -7, -17, -27, -18, -541, -8, -7, 7]
// }
