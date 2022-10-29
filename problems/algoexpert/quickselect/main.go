package main

func Quickselect(nums []int, k int) int {
	return helper(nums, 0, len(nums)-1, k-1)
}

func helper(nums []int, start, end, position int) int {
	for {
		if start > end {
			panic("This should never happen!")
		}
		pivot, hi := start, end
		for lo := start + 1; lo <= hi; {
			if nums[lo] > nums[hi] && nums[hi] < nums[pivot] {
				nums[lo], nums[hi] = nums[hi], nums[lo]
			}
			if nums[lo] <= nums[pivot] {
				lo++
			}
			if nums[hi] >= nums[pivot] {
				hi--
			}
		}
		nums[pivot], nums[hi] = nums[hi], nums[pivot]

		switch {
		case hi < position:
			start = hi + 1
		case hi > position:
			end = hi - 1
		default:
			return nums[hi]
		}
	}
}

// Test Case 1
//
// {
//   "array": [8, 5, 2, 9, 7, 6, 3],
//   "k": 3
// }
//
// Test Case 2
//
// {
//   "array": [1],
//   "k": 1
// }
//
// Test Case 3
//
// {
//   "array": [43, 24, 37],
//   "k": 1
// }
//
// Test Case 4
//
// {
//   "array": [43, 24, 37],
//   "k": 2
// }
//
// Test Case 5
//
// {
//   "array": [43, 24, 37],
//   "k": 3
// }
//
// Test Case 6
//
// {
//   "array": [8, 3, 2, 5, 1, 7, 4, 6],
//   "k": 1
// }
//
// Test Case 7
//
// {
//   "array": [8, 3, 2, 5, 1, 7, 4, 6],
//   "k": 2
// }
//
// Test Case 8
//
// {
//   "array": [8, 3, 2, 5, 1, 7, 4, 6],
//   "k": 3
// }
//
// Test Case 9
//
// {
//   "array": [8, 3, 2, 5, 1, 7, 4, 6],
//   "k": 4
// }
//
// Test Case 10
//
// {
//   "array": [8, 3, 2, 5, 1, 7, 4, 6],
//   "k": 5
// }
//
// Test Case 11
//
// {
//   "array": [8, 3, 2, 5, 1, 7, 4, 6],
//   "k": 6
// }
//
// Test Case 12
//
// {
//   "array": [8, 3, 2, 5, 1, 7, 4, 6],
//   "k": 7
// }
//
// Test Case 13
//
// {
//   "array": [8, 3, 2, 5, 1, 7, 4, 6],
//   "k": 8
// }
//
// Test Case 14
//
// {
//   "array": [102, 41, 58, 81, 2, -5, 1000, 10021, 181, -14515, 25, 15],
//   "k": 5
// }
//
// Test Case 15
//
// {
//   "array": [102, 41, 58, 81, 2, -5, 1000, 10021, 181, -14515, 25, 15],
//   "k": 4
// }
//
// Test Case 16
//
// {
//   "array": [102, 41, 58, 81, 2, -5, 1000, 10021, 181, -14515, 25, 15],
//   "k": 9
// }
//
// Test Case 17
//
// {
//   "array": [1, 3, 71, 123, 124, 156, 814, 1294, 10024, 110000, 985181, 55516151],
//   "k": 12
// }
//
// Test Case 18
//
// {
//   "array": [1, 3, 71, 123, 124, 156, 814, 1294, 10024, 110000, 985181, 55516151],
//   "k": 4
// }
