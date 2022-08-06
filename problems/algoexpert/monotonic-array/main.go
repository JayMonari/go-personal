package main

func IsMonotonic(array []int) bool {
	onlyUp, onlyDown := true, true
	for i := 0; i < len(array)-1; i++ {
		switch {
		case array[i] < array[i+1]:
			onlyUp = false
		case array[i] > array[i+1]:
			onlyDown = false
		}
	}
	return onlyUp || onlyDown
}

// Test Case 1
// {
//   "array": [-1, -5, -10, -1100, -1100, -1101, -1102, -9001]
// }
// Test Case 2
// {
//   "array": []
// }
// Test Case 3
// {
//   "array": [1]
// }
// Test Case 4
// {
//   "array": [1, 2]
// }
// Test Case 5
// {
//   "array": [2, 1]
// }
// Test Case 6
// {
//   "array": [1, 5, 10, 1100, 1101, 1102, 9001]
// }
// Test Case 7
// {
//   "array": [-1, -5, -10, -1100, -1101, -1102, -9001]
// }
// Test Case 8
// {
//   "array": [-1, -5, -10, -1100, -900, -1101, -1102, -9001]
// }
// Test Case 9
// {
//   "array": [1, 2, 0]
// }
// Test Case 10
// {
//   "array": [1, 1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 7, 9, 10, 11]
// }
// Test Case 11
// {
//   "array": [1, 1, 2, 3, 4, 5, 5, 5, 6, 7, 8, 8, 9, 10, 11]
// }
// Test Case 12
// {
//   "array": [-1, -1, -2, -3, -4, -5, -5, -5, -6, -7, -8, -7, -9, -10, -11]
// }
// Test Case 13
// {
//   "array": [-1, -1, -2, -3, -4, -5, -5, -5, -6, -7, -8, -8, -9, -10, -11]
// }
// Test Case 14
// {
//   "array": [-1, -1, -1, -1, -1, -1, -1, -1]
// }
// Test Case 15
// {
//   "array": [1, 2, -1, -2, -5]
// }
// Test Case 16
// {
//   "array": [-1, -5, 10]
// }
// Test Case 17
// {
//   "array": [2, 2, 2, 1, 4, 5]
// }
// Test Case 18
// {
//   "array": [1, 1, 1, 2, 3, 4, 1]
// }
// Test Case 19
// {
//   "array": [1, 2, 3, 3, 2, 1]
// }
