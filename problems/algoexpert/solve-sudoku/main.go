package main

const subgridLen = 3

func SolveSudoku(board [][]int) [][]int {
	solvePartialSudoku(0, 0, board)
	return board
}

func solvePartialSudoku(row int, col int, board [][]int) bool {
	r := row
	c := col
	if c == len(board[r]) {
		r++
		c = 0
		if r == len(board) {
			return true // board is completed
		}
	}

	if board[r][c] == 0 {
		return tryDigitsAtPosition(r, c, board)
	}
	return solvePartialSudoku(r, c+1, board)
}

func tryDigitsAtPosition(row int, col int, board [][]int) bool {
	for digit := 1; digit < 10; digit++ {
		if isValidAtPosition(digit, row, col, board) {
			board[row][col] = digit
			if solvePartialSudoku(row, col+1, board) {
				return true
			}
		}
	}
	board[row][col] = 0
	return false
}

func isValidAtPosition(target int, row int, col int, board [][]int) bool {
	if rowContains(board[row], target) || columnContains(board, col, target) {
		return false
	}
	subgridRowStart := (row / subgridLen) * subgridLen
	subgridColStart := (col / subgridLen) * subgridLen
	for rowIdx := 0; rowIdx < subgridLen; rowIdx++ {
		for colIdx := 0; colIdx < subgridLen; colIdx++ {
			if board[subgridRowStart+rowIdx][subgridColStart+colIdx] == target {
				return false
			}
		}
	}
	return true
}

func rowContains(row []int, target int) bool {
	for _, n := range row {
		if target == n {
			return true
		}
	}
	return false
}

func columnContains(board [][]int, col int, target int) bool {
	for _, row := range board {
		if row[col] == target {
			return true
		}
	}
	return false
}

// Test Case 1
// {
//   "board": [
//     [7, 8, 0, 4, 0, 0, 1, 2, 0],
//     [6, 0, 0, 0, 7, 5, 0, 0, 9],
//     [0, 0, 0, 6, 0, 1, 0, 7, 8],
//     [0, 0, 7, 0, 4, 0, 2, 6, 0],
//     [0, 0, 1, 0, 5, 0, 9, 3, 0],
//     [9, 0, 4, 0, 6, 0, 0, 0, 5],
//     [0, 7, 0, 3, 0, 0, 0, 1, 2],
//     [1, 2, 0, 0, 0, 7, 4, 0, 0],
//     [0, 4, 9, 2, 0, 6, 0, 0, 7]
//   ]
// }
// Test Case 2
// {
//   "board": [
//     [0, 0, 0, 0, 3, 0, 0, 0, 9],
//     [0, 4, 0, 5, 0, 0, 0, 7, 8],
//     [2, 9, 0, 0, 0, 1, 0, 5, 0],
//     [0, 7, 8, 0, 0, 3, 0, 0, 6],
//     [0, 3, 0, 0, 6, 0, 0, 8, 0],
//     [6, 0, 0, 8, 0, 0, 9, 3, 0],
//     [0, 6, 0, 9, 0, 0, 0, 2, 7],
//     [7, 2, 0, 0, 0, 5, 0, 6, 0],
//     [8, 0, 0, 0, 7, 0, 0, 0, 0]
//   ]
// }
// Test Case 3
// {
//   "board": [
//     [5, 3, 8, 0, 1, 0, 0, 0, 0],
//     [0, 7, 9, 6, 0, 0, 0, 0, 0],
//     [0, 0, 4, 0, 0, 2, 0, 0, 0],
//     [0, 0, 7, 0, 2, 3, 4, 0, 0],
//     [0, 0, 5, 0, 8, 0, 0, 0, 9],
//     [4, 6, 0, 0, 9, 0, 0, 0, 1],
//     [0, 9, 0, 2, 3, 4, 1, 5, 0],
//     [0, 4, 1, 5, 0, 0, 2, 0, 0],
//     [0, 0, 0, 8, 6, 1, 0, 3, 0]
//   ]
// }
// Test Case 4
// {
//   "board": [
//     [0, 2, 0, 0, 9, 0, 1, 0, 0],
//     [0, 0, 7, 8, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 3, 6, 0],
//     [0, 0, 1, 9, 0, 4, 0, 0, 0],
//     [0, 0, 0, 6, 0, 5, 0, 0, 7],
//     [8, 0, 0, 0, 0, 0, 0, 0, 9],
//     [0, 0, 0, 0, 2, 0, 0, 0, 0],
//     [7, 0, 0, 0, 0, 0, 0, 8, 5],
//     [4, 9, 0, 0, 3, 0, 0, 0, 0]
//   ]
// }
