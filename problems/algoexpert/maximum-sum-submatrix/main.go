package main

import (
	"math"
)

func MaximumSumSubmatrix(matrix [][]int, size int) int {
	sums := newSumMatrix(matrix)
	maxSubMatrixSum := math.MinInt32
	for row := size - 1; row < len(matrix); row++ {
		for col := size - 1; col < len(matrix[row]); col++ {
			total := sums[row][col]

			touchesTopBorder := row-size < 0
			if !touchesTopBorder {
				total -= sums[row-size][col]
			}

			touchesLeftBorder := col-size < 0
			if !touchesLeftBorder {
				total -= sums[row][col-size]
			}

			if !(touchesTopBorder || touchesLeftBorder) {
				total += sums[row-size][col-size]
			}

			if maxSubMatrixSum < total {
				maxSubMatrixSum = total
			}
		}
	}

	return maxSubMatrixSum
}

func newSumMatrix(matrix [][]int) [][]int {
	sums := make([][]int, len(matrix))
	for i := range sums {
		sums[i] = make([]int, len(matrix[0]))
	}
	sums[0][0] = matrix[0][0]
	for i := 1; i < len(matrix[0]); i++ {
		sums[0][i] = sums[0][i-1] + matrix[0][i]
	}
	for i := 1; i < len(matrix); i++ {
		sums[i][0] = sums[i-1][0] + matrix[i][0]
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			sums[i][j] = sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1] + matrix[i][j]
		}
	}
	return sums
}

// Test Case 1
// {
//   "matrix": [
//     [5, 3, -1, 5],
//     [-7, 3, 7, 4],
//     [12, 8, 0, 0],
//     [1, -8, -8, 2]
//   ],
//   "size": 2
// }
// Test Case 2
// {
//   "matrix": [
//     [3, -4, 6, -5, 1],
//     [1, -2, 8, -4, -2],
//     [3, -8, 9, 3, 1],
//     [-7, 3, 4, 2, 7],
//     [-3, 7, -5, 7, -6]
//   ],
//   "size": 3
// }
// Test Case 3
// {
//   "matrix": [
//     [2, 4],
//     [5, 6],
//     [-3, 2]
//   ],
//   "size": 2
// }
// Test Case 4
// {
//   "matrix": [
//     [3, -4, 6, -5, 1],
//     [1, -2, 8, -4, -2],
//     [3, -8, 9, 3, 1],
//     [-7, 3, 4, 2, 7],
//     [-3, 7, -5, 7, -6],
//     [2, 4, 5, 2, 3]
//   ],
//   "size": 4
// }
// Test Case 5
// {
//   "matrix": [
//     [1]
//   ],
//   "size": 1
// }
// Test Case 6
// {
//   "matrix": [
//     [1, 1],
//     [1, 1]
//   ],
//   "size": 2
// }
// Test Case 7
// {
//   "matrix": [
//     [1, 1, 2, -1],
//     [1, 1, 2, -1]
//   ],
//   "size": 2
// }
// Test Case 8
// {
//   "matrix": [
//     [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
//   ],
//   "size": 5
// }
// Test Case 9
// {
//   "matrix": [
//     [2, 1, 1, 1, 1, 4, -1, 1, 1, 5],
//     [1, -1, 1, 1, 1, 1, -1, 1, 4, 1],
//     [-50, 12, -1, 1, 5, 1, -1, 1, 1, 1],
//     [-52, 99, 1, -1, 1, 1, -1, 1, 1, 1],
//     [1, -10, -287, 9, -1, 1, -1, 1, 1, 1],
//     [1, 2, 1, 8, 1, -1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, -1, 1, 1, 1]
//   ],
//   "size": 6
// }
// Test Case 10
// {
//   "matrix": [
//     [-1, -2, -3, -4, -5],
//     [-5, -4, -3, -2, -1],
//     [-1, -2, -3, -4, -5]
//   ],
//   "size": 2
// }
// Test Case 11
// {
//   "matrix": [
//     [-1, -2, -3, -4, -5],
//     [-5, -4, -3, -2, -1],
//     [-1, -2, -3, -4, -5],
//     [-5, -4, -3, -2, -1],
//     [-5, -4, -3, -2, -1]
//   ],
//   "size": 3
// }
// Test Case 12
// {
//   "matrix": [
//     [-1, -2, -3, -4, -5],
//     [1, 1, 1, 1, 1],
//     [-1, -10, -10, -4, -5],
//     [5, 5, 5, 5, 5],
//     [-5, -4, -3, -10, -1]
//   ],
//   "size": 1
// }
// Test Case 13
// {
//   "matrix": [
//     [-1, -2, -3, -4, -5],
//     [1, 1, 1, 1, 1],
//     [-1, -10, -10, -4, -5],
//     [5, 5, 5, 5, 5],
//     [-5, -4, -3, -10, -1]
//   ],
//   "size": 2
// }
// Test Case 14
// {
//   "matrix": [
//     [3, -4, 6, -5, 1],
//     [1, -2, 8, -4, -2],
//     [3, -8, 9, 3, 1],
//     [-7, 3, 4, 2, 7],
//     [-3, 7, -5, 7, -6]
//   ],
//   "size": 4
// }
// Test Case 15
// {
//   "matrix": [
//     [3, -4, 6, -5, 1],
//     [1, -2, 8, -4, -2],
//     [3, -8, 9, 3, 1],
//     [-7, 3, 4, 2, 7],
//     [-3, 7, -5, 7, -6]
//   ],
//   "size": 5
// }
// Test Case 16
// {
//   "matrix": [
//     [22, 24, -9, 23],
//     [12, 10, -19, 35],
//     [45, -20, -20, 99],
//     [0, 0, 0, 0],
//     [0, 0, 0, 0],
//     [-100, 200, -50, 5],
//     [5, 6, 7, 8]
//   ],
//   "size": 3
// }
