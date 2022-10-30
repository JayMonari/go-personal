package main

func ShiftedBinarySearch(nums []int, target int) int {
	return helper(nums, target, 0, len(nums)-1)
}

func helper(nums []int, target, lo, hi int) int {
	for lo <= hi {
		mid := (lo + hi) / 2
		cand := nums[mid]
		switch {
		case target == cand:
			return mid
		case nums[lo] <= cand:
			if target < cand && target >= nums[lo] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		default:
			if target > cand && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

// Test Case 1
//
// {
//   "array": [45, 61, 71, 72, 73, 0, 1, 21, 33, 37],
//   "target": 33
// }
//
// Test Case 2
//
// {
//   "array": [5, 23, 111, 1],
//   "target": 111
// }
//
// Test Case 3
//
// {
//   "array": [111, 1, 5, 23],
//   "target": 5
// }
//
// Test Case 4
//
// {
//   "array": [23, 111, 1, 5],
//   "target": 35
// }
//
// Test Case 5
//
// {
//   "array": [61, 71, 72, 73, 0, 1, 21, 33, 37, 45],
//   "target": 33
// }
//
// Test Case 6
//
// {
//   "array": [72, 73, 0, 1, 21, 33, 37, 45, 61, 71],
//   "target": 72
// }
//
// Test Case 7
//
// {
//   "array": [71, 72, 73, 0, 1, 21, 33, 37, 45, 61],
//   "target": 73
// }
//
// Test Case 8
//
// {
//   "array": [73, 0, 1, 21, 33, 37, 45, 61, 71, 72],
//   "target": 70
// }
//
// Test Case 9
//
// {
//   "array": [33, 37, 45, 61, 71, 72, 73, 355, 0, 1, 21],
//   "target": 355
// }
//
// Test Case 10
//
// {
//   "array": [33, 37, 45, 61, 71, 72, 73, 355, 0, 1, 21],
//   "target": 354
// }
//
// Test Case 11
//
// {
//   "array": [45, 61, 71, 72, 73, 0, 1, 21, 33, 37],
//   "target": 45
// }
//
// Test Case 12
//
// {
//   "array": [45, 61, 71, 72, 73, 0, 1, 21, 33, 37],
//   "target": 61
// }
//
// Test Case 13
//
// {
//   "array": [45, 61, 71, 72, 73, 0, 1, 21, 33, 37],
//   "target": 71
// }
//
// Test Case 14
//
// {
//   "array": [45, 61, 71, 72, 73, 0, 1, 21, 33, 37],
//   "target": 72
// }
//
// Test Case 15
//
// {
//   "array": [45, 61, 71, 72, 73, 0, 1, 21, 33, 37],
//   "target": 73
// }
//
// Test Case 16
//
// {
//   "array": [45, 61, 71, 72, 73, 0, 1, 21, 33, 37],
//   "target": 0
// }
//
// Test Case 17
//
// {
//   "array": [45, 61, 71, 72, 73, 0, 1, 21, 33, 37],
//   "target": 1
// }
//
// Test Case 18
//
// {
//   "array": [45, 61, 71, 72, 73, 0, 1, 21, 33, 37],
//   "target": 21
// }
//
// Test Case 19
//
// {
//   "array": [45, 61, 71, 72, 73, 0, 1, 21, 33, 37],
//   "target": 37
// }
//
// Test Case 20
//
// {
//   "array": [45, 61, 71, 72, 73, 0, 1, 21, 33, 37],
//   "target": 38
// }
//
// Test Case 21
//
// {
//   "array": [0, 1, 21, 33, 37, 45, 61, 71, 72, 73],
//   "target": 0
// }
//
// Test Case 22
//
// {
//   "array": [0, 1, 21, 33, 37, 45, 61, 71, 72, 73],
//   "target": 1
// }
//
// Test Case 23
//
// {
//   "array": [0, 1, 21, 33, 37, 45, 61, 71, 72, 73],
//   "target": 21
// }
//
// Test Case 24
//
// {
//   "array": [0, 1, 21, 33, 37, 45, 61, 71, 72, 73],
//   "target": 33
// }
//
// Test Case 25
//
// {
//   "array": [0, 1, 21, 33, 37, 45, 61, 71, 72, 73],
//   "target": 37
// }
//
// Test Case 26
//
// {
//   "array": [0, 1, 21, 33, 37, 45, 61, 71, 72, 73],
//   "target": 45
// }
//
// Test Case 27
//
// {
//   "array": [0, 1, 21, 33, 37, 45, 61, 71, 72, 73],
//   "target": 61
// }
//
// Test Case 28
//
// {
//   "array": [0, 1, 21, 33, 37, 45, 61, 71, 72, 73],
//   "target": 71
// }
//
// Test Case 29
//
// {
//   "array": [0, 1, 21, 33, 37, 45, 61, 71, 72, 73],
//   "target": 72
// }
//
// Test Case 30
//
// {
//   "array": [0, 1, 21, 33, 37, 45, 61, 71, 72, 73],
//   "target": 73
// }
//
// Test Case 31
//
// {
//   "array": [0, 1, 21, 33, 37, 45, 61, 71, 72, 73],
//   "target": 38
// }
