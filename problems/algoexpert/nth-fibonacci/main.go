package main

func GetNthFib(n int) int {
	if n <= 1 {
		return 0
	}
	fib1, fib2, next := 0, 1, 1
	for n > 2 {
		next = fib1 + fib2
		fib1, fib2 = fib2, next
		n--
	}
	return next
}

// Test Case 1
// {
//   "n": 6
// }
// Test Case 2
// {
//   "n": 1
// }
// Test Case 3
// {
//   "n": 2
// }
// Test Case 4
// {
//   "n": 3
// }
// Test Case 5
// {
//   "n": 4
// }
// Test Case 6
// {
//   "n": 5
// }
// Test Case 7
// {
//   "n": 7
// }
// Test Case 8
// {
//   "n": 8
// }
// Test Case 9
// {
//   "n": 9
// }
// Test Case 10
// {
//   "n": 10
// }
// Test Case 11
// {
//   "n": 11
// }
// Test Case 12
// {
//   "n": 12
// }
// Test Case 13
// {
//   "n": 13
// }
// Test Case 14
// {
//   "n": 14
// }
// Test Case 15
// {
//   "n": 15
// }
// Test Case 16
// {
//   "n": 16
// }
// Test Case 17
// {
//   "n": 17
// }
// Test Case 18
// {
//   "n": 18
// }
