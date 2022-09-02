package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func MinHeightBST(nums []int) *BST {
	return newBST(nums, 0, len(nums)-1)
}

func newBST(nums []int, start, end int) *BST {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	bst := &BST{Value: nums[mid]}
	bst.Left = newBST(nums, start, mid-1)
	bst.Right = newBST(nums, mid+1, end)
	return bst
}

// Test Case 1
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22]
// }
// Test Case 2
// {
//   "array": [1]
// }
// Test Case 3
// {
//   "array": [1, 2]
// }
// Test Case 4
// {
//   "array": [1, 2, 5]
// }
// Test Case 5
// {
//   "array": [1, 2, 5, 7]
// }
// Test Case 6
// {
//   "array": [1, 2, 5, 7, 10]
// }
// Test Case 7
// {
//   "array": [1, 2, 5, 7, 10, 13]
// }
// Test Case 8
// {
//   "array": [1, 2, 5, 7, 10, 13, 14]
// }
// Test Case 9
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15]
// }
// Test Case 10
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22]
// }
// Test Case 11
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28]
// }
// Test Case 12
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32]
// }
// Test Case 13
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36]
// }
// Test Case 14
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36, 89]
// }
// Test Case 15
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36, 89, 92]
// }
// Test Case 16
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36, 89, 92, 9000]
// }
// Test Case 17
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36, 89, 92, 9000, 9001]
// }
