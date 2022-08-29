package main

import "math"

func SubarraySort(nums []int) []int {
	minDisordered, maxDisordered := math.MaxInt32, math.MinInt32
	for i, n := range nums {
		if isOutOfOrder(i, n, nums) {
			minDisordered = min(minDisordered, n)
			maxDisordered = max(maxDisordered, n)
		}
	}
	if minDisordered == math.MaxInt32 {
		return []int{-1, -1}
	}

	subLeftIdx := 0
	for ; minDisordered >= nums[subLeftIdx]; subLeftIdx++ {
	}
	subRightIdx := len(nums) - 1
	for ; maxDisordered <= nums[subRightIdx]; subRightIdx-- {
	}
	return []int{subLeftIdx, subRightIdx}
}

func isOutOfOrder(i, n int, nums []int) bool {
	switch i {
	case 0:
		return n > nums[i+1]
	case len(nums) - 1:
		return n < nums[i-1]
	default:
		return n > nums[i+1] || n < nums[i-1]
	}
}

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

// Test Case 1
// {
//   "array": [1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19]
// }
// Test Case 2
// {
//   "array": [1, 2]
// }
// Test Case 3
// {
//   "array": [2, 1]
// }
// Test Case 4
// {
//   "array": [1, 2, 4, 7, 10, 11, 7, 12, 7, 7, 16, 18, 19]
// }
// Test Case 5
// {
//   "array": [1, 2, 4, 7, 10, 11, 7, 12, 13, 14, 16, 18, 19]
// }
// Test Case 6
// {
//   "array": [1, 2, 8, 4, 5]
// }
// Test Case 7
// {
//   "array": [4, 8, 7, 12, 11, 9, -1, 3, 9, 16, -15, 51, 7]
// }
// Test Case 8
// {
//   "array": [4, 8, 7, 12, 11, 9, -1, 3, 9, 16, -15, 11, 57]
// }
// Test Case 9
// {
//   "array": [-41, 8, 7, 12, 11, 9, -1, 3, 9, 16, -15, 11, 57]
// }
// Test Case 10
// {
//   "array": [-41, 8, 7, 12, 11, 9, -1, 3, 9, 16, -15, 51, 7]
// }
// Test Case 11
// {
//   "array": [1, 2, 3, 4, 5, 6, 8, 7, 9, 10, 11]
// }
// Test Case 12
// {
//   "array": [1, 2, 3, 4, 5, 6, 18, 7, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19]
// }
// Test Case 13
// {
//   "array": [1, 2, 3, 4, 5, 6, 18, 21, 22, 7, 14, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 4, 14, 11, 6, 33, 35, 41, 55]
// }
// Test Case 14
// {
//   "array": [1, 2, 20, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]
// }
// Test Case 15
// {
//   "array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 2]
// }
// Test Case 16
// {
//   "array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20]
// }
// Test Case 17
// {
//   "array": [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89]
// }
