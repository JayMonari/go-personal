package main

import (
	"math"
)

func MaximizeExpression(nums []int) int {
	if len(nums) < 4 {
		return 0
	}

	maxOfA := []int{nums[0]}
	maxOfAMinusB := []int{math.MinInt32}
	maxOfAMinusBPlusC := []int{math.MinInt32, math.MinInt32}
	maxOfAMinusBPlusCMinusD := []int{math.MinInt32, math.MinInt32, math.MinInt32}

	for i := 1; i < len(nums); i++ {
		curr := max(maxOfA[i-1], nums[i])
		maxOfA = append(maxOfA, curr)
	}

	for i := 1; i < len(nums); i++ {
		curr := max(maxOfAMinusB[i-1], maxOfA[i-1]-nums[i])
		maxOfAMinusB = append(maxOfAMinusB, curr)
	}

	for i := 2; i < len(nums); i++ {
		curr := max(
			maxOfAMinusBPlusC[i-1], maxOfAMinusB[i-1]+nums[i])
		maxOfAMinusBPlusC = append(maxOfAMinusBPlusC, curr)
	}

	for i := 3; i < len(nums); i++ {
		curr := max(
			maxOfAMinusBPlusCMinusD[i-1], maxOfAMinusBPlusC[i-1]-nums[i])
		maxOfAMinusBPlusCMinusD = append(maxOfAMinusBPlusCMinusD, curr)
	}

	return maxOfAMinusBPlusCMinusD[len(maxOfAMinusBPlusCMinusD)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test Case 1
// {
//   "array": [3, 6, 1, -3, 2, 7]
// }
// Test Case 2
// {
//   "array": [3, 9, 10, 1, 30, 40]
// }
// Test Case 3
// {
//   "array": [40, 30, 1, 10, 9, 3]
// }
// Test Case 4
// {
//   "array": [-1, 2, -1, -2, -2, 1, -1]
// }
// Test Case 5
// {
//   "array": [10, 5, 10, 5]
// }
// Test Case 6
// {
//   "array": [0, 0, 0, 0, 0, 0, 0, 1, 1, 0]
// }
// Test Case 7
// {
//   "array": [34, 21, 22, 0, -98, -72, 100, 23]
// }
// Test Case 8
// {
//   "array": [5, 2, 2, 1, -2, 2, -9, 0]
// }
// Test Case 9
// {
//   "array": [1, 1, 1, 1]
// }
// Test Case 10
// {
//   "array": [1, -1, 1, -1]
// }
// Test Case 11
// {
//   "array": [3, 6, 1, 2, -9, -2, 1, 3, 4, -3, 2]
// }
// Test Case 12
// {
//   "array": [1, -1, 1, -1, -2]
// }
// Test Case 13
// {
//   "array": [3, -1, 1, -1, -2, 4, 5, -4]
// }
// Test Case 14
// {
//   "array": [-1, 2, -3, -3, 2, -1]
// }
// Test Case 15
// {
//   "array": [6, 2, 3, 4, 1, -1, -2, 3, 1, 7, 8, -8, 2, 5, 1]
// }
// Test Case 16
// {
//   "array": [5, 10, 5, 10]
// }
// Test Case 17
// {
//   "array": [2, 3]
// }
// Test Case 18
// {
//   "array": [2, 3, 4]
// }
// Test Case 19
// {
//   "array": [1]
// }
// Test Case 20
// {
//  // }
