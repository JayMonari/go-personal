package main

func IndexEqualsValue(nums []int) int {
	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		target := nums[mid]
		if (target == mid && mid == 0) || (target == mid && nums[mid-1] < mid-1) {
			return mid
		}

		switch {
		case target < mid:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

// Test Case 1
//
// {
//   "array": [-5, -3, 0, 3, 4, 5, 9]
// }
//
// Test Case 2
//
// {
//   "array": [-12, 1, 2, 3, 12]
// }
//
// Test Case 3
//
// {
//   "array": [-5, -4, -3, -2, -1, 0, 1, 3, 5, 6, 7, 11, 12, 14, 19, 20]
// }
//
// Test Case 4
//
// {
//   "array": [-3, -1, 1, 3, 5, 7, 9]
// }
//
// Test Case 5
//
// {
//   "array": [-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5]
// }
//
// Test Case 6
//
// {
//   "array": [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
// }
//
// Test Case 7
//
// {
//   "array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
// }
//
// Test Case 8
//
// {
//   "array": [0]
// }
//
// Test Case 9
//
// {
//   "array": [0, 1]
// }
//
// Test Case 10
//
// {
//   "array": [1, 2]
// }
//
// Test Case 11
//
// {
//   "array": [-50, 1, 2, 3, 4]
// }
//
// Test Case 12
//
// {
//   "array": [-40, -20, 0, 2, 4, 6, 8, 10]
// }
//
// Test Case 13
//
// {
//   "array": [-1000, 0, 1, 5000, 5001, 5002, 5005]
// }
//
// Test Case 14
//
// {
//   "array": [-1000, 0, 1, 2, 3, 4, 6, 5000, 5001, 5002, 5005]
// }
//
// Test Case 15
//
// {
//   "array": [-999, 0, 2, 500, 1000, 1500, 2000, 2500, 3000, 3500]
// }
//
// Test Case 16
//
// {
//   "array": [-9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 18]
// }
//
// Test Case 17
//
// {
//   "array": []
// }
