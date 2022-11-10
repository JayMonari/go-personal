package main

func RadixSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	digit := 0
	for max(nums)/pow(10, digit) > 0 {
		countingSort(nums, digit)
		digit++
	}
	return nums
}

var digitCounts = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func countingSort(nums []int, digit int) {
	sortedArray := make([]int, len(nums))

	digitColumn := pow(10, digit)
	for _, n := range nums {
		digitCounts[(n/digitColumn)%10]++
	}

	for i := 1; i < 10; i++ {
		digitCounts[i] += digitCounts[i-1]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		idx := (nums[i] / digitColumn) % 10
		digitCounts[idx]--
		sortedArray[digitCounts[idx]] = nums[i]
	}

	for i := range nums {
		nums[i] = sortedArray[i]
	}

	for i := range digitCounts {
		digitCounts[i] = 0
	}
}

func max(nums []int) (mx int) {
	mx = nums[0]
	for _, n := range nums {
		if mx < n {
			mx = n
		}
	}
	return mx
}

func pow(a, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result *= a
	}
	return result
}

// Test Case 1
//
// {
//   "array": [8762, 654, 3008, 345, 87, 65, 234, 12, 2]
// }
//
// Test Case 2
//
// {
//   "array": [2, 12, 65, 87, 234, 345, 654, 3008, 8762]
// }
//
// Test Case 3
//
// {
//   "array": [111, 11, 11, 1, 0]
// }
//
// Test Case 4
//
// {
//   "array": [12, 123, 456, 986, 2, 3, 34, 543, 97654, 34200]
// }
//
// Test Case 5
//
// {
//   "array": [4, 44, 444, 888, 88, 33, 3, 22, 2222, 1111, 1234]
// }
//
// Test Case 6
//
// {
//   "array": [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]
// }
//
// Test Case 7
//
// {
//   "array": []
// }
//
// Test Case 8
//
// {
//   "array": [100]
// }
//
// Test Case 9
//
// {
//   "array": [10000, 100001, 10001, 10101, 101, 11, 100, 10, 1, 0]
// }
//
// Test Case 10
//
// {
//   "array": [34, 56, 7373, 2321, 3421, 8272, 232, 23892831, 230983, 2312, 7878, 87, 234, 23, 332, 4556]
// }
//
// Test Case 11
//
// {
//   "array": [10, 87, 2321, 3221, 2312, 7632, 6542, 3223, 231, 2342, 321, 9, 1, 323, 421, 325, 65, 789, 4002]
// }
//
// Test Case 12
//
// {
//   "array": [0, 1, 2, 22, 11, 3, 33, 0, 0, 0, 21, 21, 21, 1, 11, 111]
// }
//
// Test Case 13
//
// {
//   "array": [8, 4, 5, 34, 234, 987, 444, 23, 21, 8, 1, 0]
// }
//
// Test Case 14
//
// {
//   "array": [1, 11]
// }
//
// Test Case 15
//
// {
//   "array": [1, 11, 1, 11, 101, 9, 99, 432, 441]
// }
//
// Test Case 16
//
// {
//   "array": [1000, 100, 10, 1, 10, 100, 1000, 10001, 10201, 1001, 0, 1, 1]
// }
//
// Test Case 17
//
// {
//   "array": [9, 109, 908, 876, 1099, 190, 290, 999, 9999]
// }
//
// Test Case 18
//
// {
//   "array": [0, 999999, 234892, 10000009, 89892, 782731, 891932, 92012, 1892193, 181730, 785239, 2314451, 1231421, 812723]
// }
