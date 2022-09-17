package main

import "math"

func MaxSumIncreasingSubsequence(nums []int) (max int, sequence []int) {
	seqs := make([]int, len(nums))
	sums := make([]int, len(nums))
	for i := range seqs {
		seqs[i] = math.MinInt32
		sums[i] = nums[i]
	}

	maxSumIdx := 0
	for i, n := range nums {
		for j := 0; j < i; j++ {
			if nums[j] < n && sums[j]+n >= sums[i] {
				sums[i] = sums[j] + n
				seqs[i] = j
			}
		}
		if sums[i] > sums[maxSumIdx] {
			maxSumIdx = i
		}
	}
	return sums[maxSumIdx], buildSequence(nums, seqs, maxSumIdx)
}

func buildSequence(nums []int, seqs []int, i int) (sequence []int) {
	for i != math.MinInt32 {
		sequence = append(sequence, nums[i])
		i = seqs[i]
	}
	return reverse(sequence)
}

func reverse(a []int) []int {
	c := make([]int, len(a))
	copy(c, a)
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}
	return c
}

// Test Case 1
// {
//   "array": [10, 70, 20, 30, 50, 11, 30]
// }
// Test Case 2
// {
//   "array": [1]
// }
// Test Case 3
// {
//   "array": [-1]
// }
// Test Case 4
// {
//   "array": [-1, 1]
// }
// Test Case 5
// {
//   "array": [5, 4, 3, 2, 1]
// }
// Test Case 6
// {
//   "array": [1, 2, 3, 4, 5]
// }
// Test Case 7
// {
//   "array": [-5, -4, -3, -2, -1]
// }
// Test Case 8
// {
//   "array": [8, 12, 2, 3, 15, 5, 7]
// }
// Test Case 9
// {
//   "array": [10, 15, 4, 5, 11, 14, 31, 25, 31, 23, 25, 31, 50]
// }
// Test Case 10
// {
//   "array": [10, 1, 2, 3, 4, 5, 6, 7, 8, 9]
// }
