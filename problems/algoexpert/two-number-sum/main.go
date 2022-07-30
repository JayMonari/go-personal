package main

func TwoNumberSum(array []int, target int) []int {
	// Save all previous numbers to get target
	seen := map[int]struct{}{}
	for _, n := range array {
		pair := target - n
		if _, ok := seen[pair]; ok {
			return []int{pair, n}
		}
		seen[n] = struct{}{}
	}
	return []int{}
}

//  Test Case 1
// {
//   "array": [3, 5, -4, 8, 11, 1, -1, 6],
//   "targetSum": 10
// }
// Test Case 2
// {
//   "array": [4, 6],
//   "targetSum": 10
// }
// Test Case 3
// {
//   "array": [4, 6, 1],
//   "targetSum": 5
// }
// Test Case 4
// {
//   "array": [4, 6, 1, -3],
//   "targetSum": 3
// }
// Test Case 5
// {
//   "array": [1, 2, 3, 4, 5, 6, 7, 8, 9],
//   "targetSum": 17
// }
// Test Case 6
// {
//   "array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 15],
//   "targetSum": 18
// }
// Test Case 7
// {
//   "array": [-7, -5, -3, -1, 0, 1, 3, 5, 7],
//   "targetSum": -5
// }
// Test Case 8
// {
//   "array": [-21, 301, 12, 4, 65, 56, 210, 356, 9, -47],
//   "targetSum": 163
// }
// Test Case 9
// {
//   "array": [-21, 301, 12, 4, 65, 56, 210, 356, 9, -47],
//   "targetSum": 164
// }
// Test Case 10
// {
//   "array": [3, 5, -4, 8, 11, 1, -1, 6],
//   "targetSum": 15
// }
// Test Case 11
// {
//   "array": [14],
//   "targetSum": 15
// }
// Test Case 12
// {
//   "array": [15],
//   "targetSum": 15
// }
