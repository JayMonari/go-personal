package main

func NextGreaterElementLeft(nums []int) []int {
	out := make([]int, len(nums))
	for i := range nums {
		out[i] = -1
	}

	var stack []int
	for i := 0; i < 2*len(nums); i++ {
		ringIdx := i % len(nums)
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[ringIdx] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			out[top] = nums[ringIdx]
		}
		stack = append(stack, ringIdx)
	}
	return out
}

func NextGreaterElementRight(nums []int) []int {
	out := make([]int, len(nums))
	for i := range nums {
		out[i] = -1
	}
	var stack []int
	for i := 2*len(nums) - 1; i >= 0; i-- {
		ringIdx := i % len(nums)
		for len(stack) > 0 {
			if stack[len(stack)-1] > nums[ringIdx] {
				out[ringIdx] = stack[len(stack)-1]
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[ringIdx])
	}
	return out
}

// Test Case 1
// {
//   "array": [2, 5, -3, -4, 6, 7, 2]
// }
// Test Case 2
// {
//   "array": [0, 1, 2, 3, 4]
// }
// Test Case 3
// {
//   "array": [6, 4, 5, 7, 2, 1, 3]
// }
// Test Case 4
// {
//   "array": [1, 0, 1, 0, 1, 0, 1]
// }
// Test Case 5
// {
//   "array": [5, 6, 1, 3, 1, -2, -1, 3, 4, 5]
// }
// Test Case 6
// {
//   "array": [7, 6, 5, 4, 3, 2, 1]
// }
// Test Case 7
// {
//   "array": [5, 6, 1, 2, 3, 4]
// }
// Test Case 8
// {
//   "array": [1, 1, 1, 1, 1, 1, 1, 1]
// }
// Test Case 9
// {
//   "array": [12]
// }
// Test Case 10
// {
//   "array": [12, 4]
// }
// Test Case 11
// {
//   "array": [-9, 0, -5, 1, 3, -2, 18, 2, 5, 18]
// }
// Test Case 12
// {
//   "array": [2, 6, 7, 2, 2, 2]
// }
// Test Case 13
// {
//   "array": [1, 2, 3, 4, 1, 2, 3, 4, -8, -7, 6, 2, 17, 2, -8, 9, 0, 2]
// }
// Test Case 14
// {
//   "array": [-8, -1, -1, -2, -4, -5, -6, 0, -9, -91, -2, 8]
// }
// Test Case 15
// {
//   "array": []
// }
