package main

func SortStack(stack []int) []int {
	if len(stack) == 0 {
		return stack
	}
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	SortStack(stack)

	insert(&stack, top)
	return stack
}

func insert(stack *[]int, n int) {
	if len(*stack) == 0 || (*stack)[len(*stack)-1] <= n {
		*stack = append(*stack, n)
		return
	}
	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	insert(stack, n)
	*stack = append(*stack, top)
}

// Test Case 1
// {
//   "stack": [-5, 2, -2, 4, 3, 1]
// }
// Test Case 2
// {
//   "stack": [3, 4, 5, 1, 2]
// }
// Test Case 3
// {
//   "stack": [0, -2, 3, 4, 1, -9, 8]
// }
// Test Case 4
// {
//   "stack": [2, 4, 22, 1, -9, 0, 6, 23, -2, 1]
// }
// Test Case 5
// {
//   "stack": [3, 4, 5, 1, 2]
// }
// Test Case 6
// {
//   "stack": [-1, 0, 2, 3, 4, 1, 1, 1]
// }
// Test Case 7
// {
//   "stack": []
// }
// Test Case 8
// {
//   "stack": [1]
// }
// Test Case 9
// {
//   "stack": [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]
// }
// Test Case 10
// {
//   "stack": [9, 2, 8, 1]
// }
// Test Case 11
// {
//   "stack": [2, 33, 44, 2, -9, -7, -5, -2, -2, -2, 0]
// }
// Test Case 12
// {
//   "stack": [3, 3, 3, 3, 3, 3]
// }
// Test Case 13
// {
//   "stack": [0, 0]
// }
// Test Case 14
// {
//   "stack": [2, 22, 222, 3, 33, 33, 9, 2, 3, 312, -9, -2, 3]
// }
// Test Case 15
// {
//   "stack": [3, 4, 5, 1, 2, 2, 2, 1, 3, 4, 5, 3, 1, 3, -1, 2, 3]
// }
