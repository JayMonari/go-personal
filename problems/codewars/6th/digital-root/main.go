package main

import (
	"fmt"
	"strconv"
)

// https://www.codewars.com/kata/541c8630095125aba6000c00/train/go
func DigitalRoot(n int) (root int) {
	for ; root == 0 || len(strconv.Itoa(root)) != 1; n = root {
		root = 0
		for _, v := range strconv.Itoa(n) {
			root += int(v - '0')
		}
	}
	return root
}

func main() {
	fmt.Println(DigitalRoot(16))
	fmt.Println(DigitalRoot(195))
	fmt.Println(DigitalRoot(167346))
}
