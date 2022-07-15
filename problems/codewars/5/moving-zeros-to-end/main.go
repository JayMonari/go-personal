package main

import "fmt"

// https://www.codewars.com/kata/52597aa56021e91c93000cb0/train/go
func MoveZeros(nums []int) []int {
	i := 0
	for _, n := range nums {
		fmt.Println(n, n != 0)
		if n != 0 {
			nums[i] = n
			i++
		}
	}
	for j := len(nums) - i; j < len(nums); j++ {
		nums[j] = 0
	}
	return nums
}

func main() {
	// []
	// [9, 9, 1, 2, 1, 1, 3, 1, 9, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]

	// [5, 9, 9, 6, 7, 8, 1, 1, 3, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	// [5, 6, 7, 8, 1, 1, 3, 1, 9, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]

	fmt.Println(MoveZeros([]int{9, 0, 9, 1, 2, 1, 1, 3, 1, 9, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0}))
}
