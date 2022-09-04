package main

func NumberOfWaysToMakeChange(n int, coins []int) int {
	ways := make([]int, n+1)
	ways[0] = 1
	for _, coin := range coins {
		for amt := 1; amt < n+1; amt++ {
			if coin <= amt {
				ways[amt] += ways[amt-coin]
			}
		}
	}
	return ways[n]
}

// Test Case 1
// {
//   "n": 6,
//   "denoms": [1, 5]
// }
// Test Case 2
// {
//   "n": 0,
//   "denoms": [2, 3, 4, 7]
// }
// Test Case 3
// {
//   "n": 9,
//   "denoms": [5]
// }
// Test Case 4
// {
//   "n": 7,
//   "denoms": [2, 4]
// }
// Test Case 5
// {
//   "n": 4,
//   "denoms": [1, 5, 10, 25]
// }
// Test Case 6
// {
//   "n": 5,
//   "denoms": [1, 5, 10, 25]
// }
// Test Case 7
// {
//   "n": 10,
//   "denoms": [1, 5, 10, 25]
// }
// Test Case 8
// {
//   "n": 25,
//   "denoms": [1, 5, 10, 25]
// }
// Test Case 9
// {
//   "n": 12,
//   "denoms": [2, 3, 7]
// }
// Test Case 10
// {
//   "n": 7,
//   "denoms": [2, 3, 4, 7]
// }
