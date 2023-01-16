package main

import "math"

func MaxProfitWithKTransactions(prices []int, k int) int {
	if len(prices) == 0 {
		return 0
	}
	evenProfits := make([]int, len(prices))
	oddProfits := make([]int, len(prices))
	var currProfits, prevProfits []int
	for t := 1; t < k+1; t++ {
		maxThusFar := math.MinInt32
		switch t % 2 {
		case 1: // odd
			currProfits, prevProfits = oddProfits, evenProfits
		case 0: // even
			currProfits, prevProfits = evenProfits, oddProfits
		}
		for d := 1; d < len(prices); d++ {
			maxThusFar = max(maxThusFar, prevProfits[d-1]-prices[d-1])
			currProfits[d] = max(currProfits[d-1], maxThusFar+prices[d])
		}
	}
	if k%2 == 0 {
		return evenProfits[len(prices)-1]
	}
	return oddProfits[len(prices)-1]
}

func max(n int, rest ...int) int {
	curr := n
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}

// Test Case 1
//
// {
//   "prices": [5, 11, 3, 50, 60, 90],
//   "k": 2
// }
//
// Test Case 2
//
// {
//   "prices": [],
//   "k": 1
// }
//
// Test Case 3
//
// {
//   "prices": [1],
//   "k": 1
// }
//
// Test Case 4
//
// {
//   "prices": [1, 10],
//   "k": 1
// }
//
// Test Case 5
//
// {
//   "prices": [1, 10],
//   "k": 3
// }
//
// Test Case 6
//
// {
//   "prices": [3, 2, 5, 7, 1, 3, 7],
//   "k": 1
// }
//
// Test Case 7
//
// {
//   "prices": [5, 11, 3, 50, 60, 90],
//   "k": 3
// }
//
// Test Case 8
//
// {
//   "prices": [5, 11, 3, 50, 40, 90],
//   "k": 2
// }
//
// Test Case 9
//
// {
//   "prices": [5, 11, 3, 50, 40, 90],
//   "k": 3
// }
//
// Test Case 10
//
// {
//   "prices": [50, 25, 12, 4, 3, 10, 1, 100],
//   "k": 2
// }
//
// Test Case 11
//
// {
//   "prices": [100, 99, 98, 97, 1],
//   "k": 5
// }
//
// Test Case 12
//
// {
//   "prices": [1, 100, 2, 200, 3, 300, 4, 400, 5, 500],
//   "k": 5
// }
//
// Test Case 13
//
// {
//   "prices": [1, 100, 101, 200, 201, 300, 301, 400, 401, 500],
//   "k": 5
// }
//
// Test Case 14
//
// {
//   "prices": [1, 25, 24, 23, 12, 36, 14, 40, 31, 41, 5],
//   "k": 4
// }
//
// Test Case 15
//
// {
//   "prices": [1, 25, 24, 23, 12, 36, 14, 40, 31, 41, 5],
//   "k": 2
// }
