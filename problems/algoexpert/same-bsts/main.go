package main

import "math"

func SameBsts(a, b []int) bool {
	return same(a, b, 0, 0, math.MinInt32, math.MaxInt32)
}

func same(a, b []int, i1, i2, min, max int) bool {
	switch {
	case i1 == -1 || i2 == -1:
		return i1 == i2
	case a[i1] != b[i2]:
		return false
	}

	currentValue := a[i1]
	return same(
		a, b, firstSmaller(a, i1, min), firstSmaller(b, i2, min), min, currentValue) &&
		same(
			a, b, firstBigger(a, i1, max), firstBigger(b, i2, max), currentValue, max)
}

func firstSmaller(a []int, start, min int) (idx int) {
	for i := start + 1; i < len(a); i++ {
		if a[i] < a[start] && a[i] >= min {
			return i
		}
	}
	return -1
}

func firstBigger(a []int, start, max int) (idx int) {
	for i := start + 1; i < len(a); i++ {
		if a[i] >= a[start] && a[i] < max {
			return i
		}
	}
	return -1
}

func SameBsts2(a, b []int) bool {
	switch {
	case len(a) == 0 && len(b) == 0:
		return true
	case len(a) != len(b):
		return false
	case a[0] != b[0]:
		return false
	}

	return SameBsts(smallSide(a), smallSide(b)) && SameBsts(bigSide(a), bigSide(b))
}

func smallSide(a []int) (small []int) {
	for _, n := range a[1:] {
		if n < a[0] {
			small = append(small, n)
		}
	}
	return small
}

func bigSide(a []int) (big []int) {
	for _, n := range a[1:] {
		if n >= a[0] {
			big = append(big, n)
		}
	}
	return big
}

// Test Case 1
// {
//   "arrayOne": [10, 15, 8, 12, 94, 81, 5, 2, 11],
//   "arrayTwo": [10, 8, 5, 15, 2, 12, 11, 94, 81]
// }
// Test Case 2
// {
//   "arrayOne": [1, 2, 3, 4, 5, 6, 7],
//   "arrayTwo": [1, 2, 3, 4, 5, 6, 7]
// }
// Test Case 3
// {
//   "arrayOne": [7, 6, 5, 4, 3, 2, 1],
//   "arrayTwo": [7, 6, 5, 4, 3, 2, 1]
// }
// Test Case 4
// {
//   "arrayOne": [10, 15, 8, 12, 94, 81, 5, 2],
//   "arrayTwo": [10, 8, 5, 15, 2, 12, 94, 81]
// }
// Test Case 5
// {
//   "arrayOne": [10, 15, 8, 12, 94, 81, 5, 2],
//   "arrayTwo": [11, 8, 5, 15, 2, 12, 94, 81]
// }
// Test Case 6
// {
//   "arrayOne": [10, 15, 8, 12, 94, 81, 5, 2, -1, 100, 45, 12, 8, -1, 8, 2, -34],
//   "arrayTwo": [10, 8, 5, 15, 2, 12, 94, 81, -1, -1, -34, 8, 2, 8, 12, 45, 100]
// }
// Test Case 7
// {
//   "arrayOne": [10, 15, 8, 12, 94, 81, 5, 2, -1, 101, 45, 12, 8, -1, 8, 2, -34],
//   "arrayTwo": [10, 8, 5, 15, 2, 12, 94, 81, -1, -1, -34, 8, 2, 8, 12, 45, 100]
// }
// Test Case 8
// {
//   "arrayOne": [5, 2, -1, 100, 45, 12, 8, -1, 8, 10, 15, 8, 12, 94, 81, 2, -34],
//   "arrayTwo": [5, 8, 10, 15, 2, 8, 12, 45, 100, 2, 12, 94, 81, -1, -1, -34, 8]
// }
// Test Case 9
// {
//   "arrayOne": [10, 15, 8, 12, 94, 81, 5, 2, -1, 100, 45, 12, 9, -1, 8, 2, -34],
//   "arrayTwo": [10, 8, 5, 15, 2, 12, 94, 81, -1, -1, -34, 8, 2, 9, 12, 45, 100]
// }
// Test Case 10
// {
//   "arrayOne": [1, 2, 3, 4, 5, 6, 7, 8],
//   "arrayTwo": [1, 2, 3, 4, 5, 6, 7]
// }
// Test Case 11
// {
//   "arrayOne": [7, 6, 5, 4, 3, 2, 1],
//   "arrayTwo": [7, 6, 5, 4, 3, 2, 1, 0]
// }
// Test Case 12
// {
//   "arrayOne": [10, 15, 8, 12, 94, 81, 5, 2, 10],
//   "arrayTwo": [10, 8, 5, 15, 2, 10, 12, 94, 81]
// }
// Test Case 13
// {
//   "arrayOne": [50, 76, 81, 23, 23, 23, 100, 56, 12, -1, 3],
//   "arrayTwo": [50, 23, 76, 23, 23, 12, 56, 81, -1, 3, 100]
// }
