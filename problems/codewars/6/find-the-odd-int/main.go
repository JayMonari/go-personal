package main

// https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go
func FindOdd(seq []int) (odd int) {
	counter := map[int]uint{}
	for _, n := range seq {
		counter[n]++
	}
	var val uint
	for odd, val = range counter {
		if val%2 != 0 {
			break
		}
	}
	return odd
}
