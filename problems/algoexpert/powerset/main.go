package main

func Powerset(nums []int) [][]int {
	subsets := [][]int{{}}
	for _, v := range nums {
		for _, s := range subsets {
			next := append([]int{}, s...)
			subsets = append(subsets, append(next, v))
		}
	}
	return subsets
}

// Test Case 1
// {
//   "array": [1, 2, 3]
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
//   "array": [1, 2, 3, 4]
// }
// Test Case 6
// {
//   "array": [1, 2, 3, 4, 5]
// }
// Test Case 7
// {
//   "array": [1, 2, 3, 4, 5, 6]
// }
