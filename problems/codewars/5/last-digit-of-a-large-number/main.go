package main

import "strconv"

// https://www.codewars.com/kata/5511b2f550906349a70004e1/train/go
func LastDigit(n1, n2 string) int {
	max := func(a, b int) int {
		if b > a {
			return b
		} else {
			return a
		}
	}
	pow := func(n, e int) int {
		res := 1
		for i := 0; i < e; i++ {
			res *= n
		}
		return res
	}
	if n2 == "0" {
		return 1
	}
	a, _ := strconv.Atoi(n1[len(n1)-1:])
	b, _ := strconv.Atoi(n2[max(0, len(n2)-2):])
	return pow(a, (b%4)+4) % 10
}
