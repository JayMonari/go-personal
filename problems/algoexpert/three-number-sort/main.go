package main

func ThreeNumberSort(nums []int, order []int) []int {
	changeIdx := 0
	for i := range nums {
		if nums[i] != order[0] {
			continue
		}
		nums[i], nums[changeIdx] = nums[changeIdx], nums[i]
		changeIdx++
	}
	changeIdx = len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != order[2] {
			continue
		}
		nums[i], nums[changeIdx] = nums[changeIdx], nums[i]
		changeIdx--
	}
	return nums
}

func ThreeNumberSortOpt(nums []int, order []int) []int {
	idx1, idx2, idx3 := 0, 0, len(nums)-1
	for idx2 <= idx3 {
		switch {
		case nums[idx2] == order[0]:
			nums[idx1], nums[idx2] = nums[idx2], nums[idx1]
			idx1++
			idx2++
		case nums[idx2] == order[1]:
			idx2++
		default:
			nums[idx2], nums[idx3] = nums[idx3], nums[idx2]
			idx3--
		}
	}
	return nums
}

// Test Case 1
// {
//   "array": [1, 0, 0, -1, -1, 0, 1, 1],
//   "order": [0, 1, -1]
// }
// Test Case 2
// {
//   "array": [7, 8, 9, 7, 8, 9, 9, 9, 9, 9, 9, 9],
//   "order": [8, 7, 9]
// }
// Test Case 3
// {
//   "array": [],
//   "order": [0, 7, 9]
// }
// Test Case 4
// {
//   "array": [-2, -3, -3, -3, -3, -3, -2, -2, -3],
//   "order": [-2, -3, 0]
// }
// Test Case 5
// {
//   "array": [0, 10, 10, 10, 10, 10, 25, 25, 25, 25, 25],
//   "order": [25, 10, 0]
// }
// Test Case 6
// {
//   "array": [6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6],
//   "order": [4, 5, 6]
// }
// Test Case 7
// {
//   "array": [1, 3, 4, 4, 4, 4, 3, 3, 3, 4, 1, 1, 1, 4, 4, 1, 3, 1, 4, 4],
//   "order": [1, 4, 3]
// }
// Test Case 8
// {
//   "array": [1, 2, 3],
//   "order": [3, 1, 2]
// }
// Test Case 9
// {
//   "array": [0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 1, 2],
//   "order": [1, 2, 0]
// }
// Test Case 10
// {
//   "array": [7, 7, 7, 11, 11, 7, 11, 7],
//   "order": [11, 7, 9]
// }
// Test Case 11
// {
//   "array": [9, 9, 9, 7, 9, 7, 9, 9, 7, 9],
//   "order": [11, 7, 9]
// }
// Test Case 12
// {
//   "array": [9, 9, 9, 7, 9, 7, 9, 9, 7, 9],
//   "order": [7, 11, 9]
// }
// Test Case 13
// {
//   "array": [1],
//   "order": [0, 1, 2]
// }
// Test Case 14
// {
//   "array": [0, 1],
//   "order": [1, 2, 0]
// }
