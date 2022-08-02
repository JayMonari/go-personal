package main

func LongestPeak(nums []int) int {
	max := 0
	n := len(nums)
	for i := 1; i < n-1; {
		if !(nums[i-1] < nums[i] && nums[i] > nums[i+1]) {
			i++
			continue
		}
		left := i - 2
		for left >= 0 && nums[left] < nums[left+1] {
			left--
		}
		right := i + 2
		for right < n && nums[right] < nums[right-1] {
			right++
		}
		if curr := right - left - 1; curr > max {
			max = curr
		}
		i = right
	}
	return max
}

// Test Case 1
// {
//   "array": [1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3]
// }
// Test Case 2
// {
//   "array": []
// }
// Test Case 3
// {
//   "array": [1, 3, 2]
// }
// Test Case 4
// {
//   "array": [1, 2, 3, 4, 5, 1]
// }
// Test Case 5
// {
//   "array": [5, 4, 3, 2, 1, 2, 1]
// }
// Test Case 6
// {
//   "array": [5, 4, 3, 2, 1, 2, 10, 12, -3, 5, 6, 7, 10]
// }
// Test Case 7
// {
//   "array": [5, 4, 3, 2, 1, 2, 10, 12]
// }
// Test Case 8
// {
//   "array": [1, 2, 3, 4, 5, 6, 10, 100, 1000]
// }
// Test Case 9
// {
//   "array": [1, 2, 3, 3, 2, 1]
// }
// Test Case 10
// {
//   "array": [1, 1, 3, 2, 1]
// }
// Test Case 11
// {
//   "array": [1, 2, 3, 2, 1, 1]
// }
// Test Case 12
// {
//   "array": [1, 1, 1, 2, 3, 10, 12, -3, -3, 2, 3, 45, 800, 99, 98, 0, -1, -1, 2, 3, 4, 5, 0, -1, -1]
// }
// Test Case 13
// {
//   "array": [1, 2, 3, 3, 4, 0, 10]
// }
