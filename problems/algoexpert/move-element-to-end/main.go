package main

func MoveElementToEnd(nums []int, toMove int) []int {
	for i, moveIdx := 0, len(nums)-1; i < moveIdx; i++ {
		for i < moveIdx && nums[moveIdx] == toMove {
			moveIdx--
		}
		if nums[i] == toMove {
			nums[i], nums[moveIdx] = nums[moveIdx], nums[i]
		}
	}
	return nums
}

// Test Case 1
// {
//   "array": [2, 1, 2, 2, 2, 3, 4, 2],
//   "toMove": 2
// }
// Test Case 2
// {
//   "array": [],
//   "toMove": 3
// }
// Test Case 3
// {
//   "array": [1, 2, 4, 5, 6],
//   "toMove": 3
// }
// Test Case 4
// {
//   "array": [3, 3, 3, 3, 3],
//   "toMove": 3
// }
// Test Case 5
// {
//   "array": [3, 1, 2, 4, 5],
//   "toMove": 3
// }
// Test Case 6
// {
//   "array": [1, 2, 4, 5, 3],
//   "toMove": 3
// }
// Test Case 7
// {
//   "array": [1, 2, 3, 4, 5],
//   "toMove": 3
// }
// Test Case 8
// {
//   "array": [2, 4, 2, 5, 6, 2, 2],
//   "toMove": 2
// }
// Test Case 9
// {
//   "array": [5, 5, 5, 5, 5, 5, 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12],
//   "toMove": 5
// }
// Test Case 10
// {
//   "array": [1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 5, 5, 5, 5, 5, 5],
//   "toMove": 5
// }
// Test Case 11
// {
//   "array": [5, 1, 2, 5, 5, 3, 4, 6, 7, 5, 8, 9, 10, 11, 5, 5, 12],
//   "toMove": 5
// }
