package main

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	sums := [][]int{}
	sort.Ints(array)
	for i, curr := range array {
		for j, k := i+1, len(array)-1; j < k; {
			sum := curr + array[j] + array[k]
			switch {
			case sum < target:
				j++
			case sum > target:
				k--
			case sum == target:
				sums = append(sums, []int{curr, array[j], array[k]})
				j, k = j+1, k-1
			}
		}
	}
	return sums
}

// Test Case 1
// {
//   "array": [12, 3, 1, 2, -6, 5, -8, 6],
//   "targetSum": 0
// }
// Test Case 2
// {
//   "array": [1, 2, 3],
//   "targetSum": 6
// }
// Test Case 3
// {
//   "array": [1, 2, 3],
//   "targetSum": 7
// }
// Test Case 4
// {
//   "array": [8, 10, -2, 49, 14],
//   "targetSum": 57
// }
// Test Case 5
// {
//   "array": [12, 3, 1, 2, -6, 5, 0, -8, -1],
//   "targetSum": 0
// }
// Test Case 6
// {
//   "array": [12, 3, 1, 2, -6, 5, 0, -8, -1, 6],
//   "targetSum": 0
// }
// Test Case 7
// {
//   "array": [12, 3, 1, 2, -6, 5, 0, -8, -1, 6, -5],
//   "targetSum": 0
// }
// Test Case 8
// {
//   "array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 15],
//   "targetSum": 18
// }
// Test Case 9
// {
//   "array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 15],
//   "targetSum": 32
// }
// Test Case 10
// {
//   "array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 15],
//   "targetSum": 33
// }
// Test Case 11
// {
//   "array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 15],
//   "targetSum": 5
// }
