package main

type irange [2]int

func SearchForRange(nums []int, target int) []int {
	finalRange := irange{-1, -1}
	alteredBinarySearch(nums, target, 0, len(nums)-1, &finalRange, true)
	alteredBinarySearch(nums, target, 0, len(nums)-1, &finalRange, false)
	return []int{finalRange[0], finalRange[1]}
}

func alteredBinarySearch(nums []int, target, lo, hi int, r *irange, goLeft bool) {
	for lo <= hi {
		mid := (lo + hi) / 2
		switch {
		case nums[mid] < target:
			lo = mid + 1
			continue
		case nums[mid] > target:
			hi = mid - 1
			continue
		}

		if goLeft {
			if mid != 0 && nums[mid-1] == target {
				hi = mid - 1
				continue
			}
			r[0] = mid
			return
		}

		if mid != len(nums)-1 && nums[mid+1] == target {
			lo = mid + 1
			continue
		}
		r[1] = mid
		return
	}
}

// Test Case 1
//
// {
//   "array": [0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73],
//   "target": 45
// }
//
// Test Case 2
//
// {
//   "array": [5, 7, 7, 8, 8, 10],
//   "target": 5
// }
//
// Test Case 3
//
// {
//   "array": [5, 7, 7, 8, 8, 10],
//   "target": 7
// }
//
// Test Case 4
//
// {
//   "array": [5, 7, 7, 8, 8, 10],
//   "target": 8
// }
//
// Test Case 5
//
// {
//   "array": [5, 7, 7, 8, 8, 10],
//   "target": 10
// }
//
// Test Case 6
//
// {
//   "array": [5, 7, 7, 8, 8, 10],
//   "target": 9
// }
//
// Test Case 7
//
// {
//   "array": [0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73],
//   "target": 47
// }
//
// Test Case 8
//
// {
//   "array": [0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73],
//   "target": -1
// }
//
// Test Case 9
//
// {
//   "array": [0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 45, 45, 45],
//   "target": 45
// }
