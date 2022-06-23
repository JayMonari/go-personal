// Package listops provides basic list operations for IntList: Length, Reverse,
// Append, Concat, Foldl, Foldr, Filter, Map
package listops

type IntList []int

type binFunc func(x, y int) int
type unaryFunc func(x int) int
type predFunc func(n int) bool

// Length returns the count of elements within the IntList.
func (il *IntList) Length() int {
	cnt := 0
	for range *il {
		cnt++
	}
	return cnt
}

// Reverse returns a new IntList reversed.
func (il IntList) Reverse() IntList {
	for i, j := 0, il.Length()-1; i < j; i, j = i+1, j-1 {
		il[i], il[j] = il[j], il[i]
	}
	return il
}

// Append returns a new IntList of all the values of il with the values of nums
// appended to the end.
func (il IntList) Append(nums []int) IntList {
	il = append(il, nums...)
	return il
}

// Concat returns a new IntList of the original values in il and all values
// from lists flattened out into a single list.
func (il IntList) Concat(lists []IntList) IntList {
	for _, ol := range lists {
		il = il.Append(ol)
	}
	return il
}

// Foldl accumulates all values of the IntList with the supplied binFunc and
// the start value, starting from the left side.
func (il *IntList) Foldl(fn binFunc, start int) int {
	for _, n := range *il {
		start = fn(start, n)
	}
	return start
}

// Foldr accumulates all values of the IntList with the supplied binFunc and
// the start value, starting from the right side.
func (il *IntList) Foldr(fn binFunc, start int) int {
	for _, n := range il.Reverse() {
		start = fn(n, start)
	}
	return start
}

// Filter returns a new IntList with the results that return true from the
// supplied predFunc.
func (il IntList) Filter(fn predFunc) IntList {
	res := make(IntList, 0, il.Length())
	for _, n := range il {
		if fn(n) {
			res = append(res, n)
		}
	}
	return res
}

// Map returns a new IntList with the supplied unaryFunc applied to each
// integer.
func (il IntList) Map(fn unaryFunc) IntList {
	res := make(IntList, 0, il.Length())
	for _, n := range il {
		res = append(res, fn(n))
	}
	return res
}
