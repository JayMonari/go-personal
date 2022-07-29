package main

func BinarySearch(array []int, target int) int {
	for lo, hi := 0, len(array)-1; lo <= hi; {
		mid := (hi + lo) / 2
		val := array[mid]
		switch {
		case val > target:
			hi = mid - 1
		case val < target:
			lo = mid + 1
		case val == target:
			return mid
		}
	}
	return -1
}

// Test Case 1
// {
//   "array": [0, 1, 21, 33, 45, 45, 61, 71, 72, 73],
//   "target": 33
// }
// Test Case 2
// {
//   "array": [1, 5, 23, 111],
//   "target": 111
// }
// Test Case 3
// {
//   "array": [1, 5, 23, 111],
//   "target": 5
// }
// Test Case 4
// {
//   "array": [1, 5, 23, 111],
//   "target": 35
// }
// Test Case 5
// {
//   "array": [0, 1, 21, 33, 45, 45, 61, 71, 72, 73],
//   "target": 0
// }
// Test Case 6
// {
//   "array": [0, 1, 21, 33, 45, 45, 61, 71, 72, 73],
//   "target": 1
// }
// Test Case 7
// {
//   "array": [0, 1, 21, 33, 45, 45, 61, 71, 72, 73],
//   "target": 21
// }
// Test Case 8
// {
//   "array": [0, 1, 21, 33, 45, 45, 61, 71, 72, 73],
//   "target": 45
// }
// Test Case 9
// {
//   "array": [0, 1, 21, 33, 45, 45, 61, 71, 72, 73],
//   "target": 61
// }
// Test Case 10
// {
//   "array": [0, 1, 21, 33, 45, 45, 61, 71, 72, 73],
//   "target": 71
// }
// Test Case 11
// {
//   "array": [0, 1, 21, 33, 45, 45, 61, 71, 72, 73],
//   "target": 72
// }
// Test Case 12
// {
//   "array": [0, 1, 21, 33, 45, 45, 61, 71, 72, 73],
//   "target": 73
// }
// Test Case 13
// {
//   "array": [0, 1, 21, 33, 45, 45, 61, 71, 72, 73],
//   "target": 70
// }
// Test Case 14
// {
//   "array": [0, 1, 21, 33, 45, 45, 61, 71, 72, 73, 355],
//   "target": 355
// }
// Test Case 15
// {
//   "array": [0, 1, 21, 33, 45, 45, 61, 71, 72, 73, 355],
//   "target": 354
// }
// Test Case 16
// {
//   "array": [1, 5, 23, 111],
//   "target": 120
// }
// Test Case 17
// {
//   "array": [5, 23, 111],
//   "target": 3
// }
