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
