package main

import "math/bits"

// https://www.codewars.com/kata/526571aae218b8ee490006f4/train/go
func CountBits(n uint) int {
	return bits.OnesCount(n)
}
