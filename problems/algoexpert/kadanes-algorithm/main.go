package main

func KadanesAlgorithm(nums []int) int {
	finalMax := nums[0]
	for i, currMax := 1, nums[0]; i < len(nums); i++ {
		currMax = max(nums[i], currMax+nums[i])
		if finalMax < currMax {
			finalMax = currMax
		}
	}
	return finalMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test Case 1
// {
//   "array": [3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4]
// }
// Test Case 2
// {
//   "array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
// }
// Test Case 3
// {
//   "array": [-1, -2, -3, -4, -5, -6, -7, -8, -9, -10]
// }
// Test Case 4
// {
//   "array": [-10, -2, -9, -4, -8, -6, -7, -1, -3, -5]
// }
// Test Case 5
// {
//   "array": [1, 2, 3, 4, 5, 6, -20, 7, 8, 9, 10]
// }
// Test Case 6
// {
//   "array": [1, 2, 3, 4, 5, 6, -22, 7, 8, 9, 10]
// }
// Test Case 7
// {
//   "array": [1, 2, -4, 3, 5, -9, 8, 1, 2]
// }
// Test Case 8
// {
//   "array": [3, 4, -6, 7, 8]
// }
// Test Case 9
// {
//   "array": [3, 4, -6, 7, 8, -18, 100]
// }
// Test Case 10
// {
//   "array": [3, 4, -6, 7, 8, -15, 100]
// }
// Test Case 11
// {
//   "array": [8, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4]
// }
// Test Case 12
// {
//   "array": [8, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 6]
// }
// Test Case 13
// {
//   "array": [8, 5, -9, 1, 3, -2, 3, 4, 7, 2, -18, 6, 3, 1, -5, 6]
// }
// Test Case 14
// {
//   "array": [8, 5, -9, 1, 3, -2, 3, 4, 7, 2, -18, 6, 3, 1, -5, 6, 20, -23, 15, 1, -3, 4]
// }
// Test Case 15
// {
//   "array": [100, 8, 5, -9, 1, 3, -2, 3, 4, 7, 2, -18, 6, 3, 1, -5, 6, 20, -23, 15, 1, -3, 4]
// }
// Test Case 16
// {
//   "array": [-1000, -1000, 2, 4, -5, -6, -7, -8, -2, -100]
// }
// Test Case 17
// {
//   "array": [-2, -1]
// }
// Test Case 18
// {
//   "array": [-2, 1]
// }
// Test Case 19
// {
//   "array": [-10]
// }
