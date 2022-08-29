package main

func SearchInSortedMatrix(matrix [][]int, target int) []int {
	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		switch {
		case matrix[row][col] > target:
			col--
		case matrix[row][col] < target:
			row++
		default:
			return []int{row, col}
		}
	}
	return []int{-1, -1}
}

// Test Case 1
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 44
// }
// Test Case 2
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 1
// }
// Test Case 3
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 2
// }
// Test Case 4
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 4
// }
// Test Case 5
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 15
// }
// Test Case 6
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 12
// }
// Test Case 7
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 32
// }
// Test Case 8
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 99
// }
// Test Case 9
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 100
// }
// Test Case 10
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 40
// }
// Test Case 11
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 128
// }
// Test Case 12
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 106
// }
// Test Case 13
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 45
// }
// Test Case 14
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 24
// }
// Test Case 15
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 43
// }
// Test Case 16
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": -1
// }
// Test Case 17
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 1000
// }
// Test Case 18
// {
//   "matrix": [
//     [1, 4, 7, 12, 15, 1000],
//     [2, 5, 19, 31, 32, 1001],
//     [3, 8, 24, 33, 35, 1002],
//     [40, 41, 42, 44, 45, 1003],
//     [99, 100, 103, 106, 128, 1004]
//   ],
//   "target": 1004
// }