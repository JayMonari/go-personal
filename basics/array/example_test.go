package array_test

import "basics/array"

func ExampleArrays() {
	array.Arrays()
	// Output:
	// Length of arr: 4
	// Empty: [0 0 0 0]
	// Filled: [0 1 2 3]
	// Grab one value: 3
	// Can be declared inline [2 3 4 5]
	// Just an array [5]int
	// Values:[8 9 10 11 15]
}

func ExampleMatrix() {
	array.Matrix()
	// Output:
	// Matrix: [[0 1 2 3 4] [1 2 3 4 5] [2 3 4 5 6] [3 4 5 6 7]]
}
