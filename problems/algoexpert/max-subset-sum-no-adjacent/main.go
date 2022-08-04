package main

func MaxSubsetSumNoAdjacent(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	one := max(nums[0], nums[1])
	two := nums[0]
	for i := 2; i < len(nums); i++ {
		next := max(one, two+nums[i])
		two = one
		one = next
	}
	return one
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test Case 1
// {
//   "array": [75, 105, 120, 75, 90, 135]
// }
// Test Case 2
// {
//   "array": []
// }
// Test Case 3
// {
//   "array": [1]
// }
// Test Case 4
// {
//   "array": [1, 2]
// }
// Test Case 5
// {
//   "array": [1, 2, 3]
// }
// Test Case 6
// {
//   "array": [1, 15, 3]
// }
// Test Case 7
// {
//   "array": [7, 10, 12, 7, 9, 14]
// }
// Test Case 8
// {
//   "array": [4, 3, 5, 200, 5, 3]
// }
// Test Case 9
// {
//   "array": [10, 5, 20, 25, 15, 5, 5, 15]
// }
// Test Case 10
// {
//   "array": [10, 5, 20, 25, 15, 5, 5, 15, 3, 15, 5, 5, 15]
// }
// Test Case 11
// {
//   "array": [125, 210, 250, 120, 150, 300]
// }
// Test Case 12
// {
//   "array": [30, 25, 50, 55, 100]
// }
// Test Case 13
// {
//   "array": [30, 25, 50, 55, 100, 120]
// }
// Test Case 14
// {
//   "array": [7, 10, 12, 7, 9, 14, 15, 16, 25, 20, 4]
// }
