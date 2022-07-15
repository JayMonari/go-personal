package main

import (
	"fmt"
	"math"
)

// https://www.codewars.com/kata/5526fc09a1bbd946250002dc/train/go
func FindOutlier(nums []int) (o int) {
	var lastBit uint8
	if (nums[0]&1)+(nums[1]&1)+(nums[2]&1) <= 1 {
		lastBit = 1
	}
	for _, o = range nums {
		if (o & 1) == int(lastBit) {
			break
		}
	}
	return
}

func main() {
	fmt.Println(FindOutlier([]int{2, 6, 8, -10, 3}))
	fmt.Println(FindOutlier([]int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781}))
	fmt.Println(FindOutlier([]int{math.MaxInt32, 0, 1}))
}
