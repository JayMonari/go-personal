package main

func ZigzagTraverse(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	width, height := len(matrix[0])-1, len(matrix)-1
	out := make([]int, 0, width+height+2)
	row, col := 0, 0
	goingDown := true
	for !(row < 0 || row > height || col < 0 || col > width) {
		out = append(out, matrix[row][col])
		if goingDown {
			if col == 0 || row == height {
				goingDown = false
				switch row {
				case height:
					col++
				default:
					row++
				}
				continue
			}
			row++
			col--
			continue
		}

		if row == 0 || col == width {
			goingDown = true
			switch col {
			case width:
				row++
			default:
				col++
			}
			continue
		}
		row--
		col++
	}
	return out
}

// Test Case 1
// {
//   "array": [
//     [1, 3, 4, 10],
//     [2, 5, 9, 11],
//     [6, 8, 12, 15],
//     [7, 13, 14, 16]
//   ]
// }
// Test Case 2
// {
//   "array": [
//     [1]
//   ]
// }
// Test Case 3
// {
//   "array": [
//     [1, 2, 3, 4, 5]
//   ]
// }
// Test Case 4
// {
//   "array": [
//     [1],
//     [2],
//     [3],
//     [4],
//     [5]
//   ]
// }
// Test Case 5
// {
//   "array": [
//     [1, 3],
//     [2, 4],
//     [5, 7],
//     [6, 8],
//     [9, 10]
//   ]
// }
// Test Case 6
// {
//   "array": [
//     [1, 3, 4, 7, 8],
//     [2, 5, 6, 9, 10]
//   ]
// }
// Test Case 7
// {
//   "array": [
//     [1, 3, 4, 10, 11],
//     [2, 5, 9, 12, 19],
//     [6, 8, 13, 18, 20],
//     [7, 14, 17, 21, 24],
//     [15, 16, 22, 23, 25]
//   ]
// }
// Test Case 8
// {
//   "array": [
//     [1, 3, 4, 10, 11, 20],
//     [2, 5, 9, 12, 19, 21],
//     [6, 8, 13, 18, 22, 27],
//     [7, 14, 17, 23, 26, 28],
//     [15, 16, 24, 25, 29, 30]
//   ]
// }
// Test Case 9
// {
//   "array": [
//     [1, 3, 4, 10, 11],
//     [2, 5, 9, 12, 20],
//     [6, 8, 13, 19, 21],
//     [7, 14, 18, 22, 27],
//     [15, 17, 23, 26, 28],
//     [16, 24, 25, 29, 30]
//   ]
// }
// Test Case 10
// {
//   "array": [
//     [1, 21, -3, 4, 15, 6, -7, 88, 9],
//     [10, 11, 112, 12, 20, -1, -2, -3, -4],
//     [6, 8, 113, 19, 21, 0, 0, 0, 0],
//     [7, 2, 18, 22, -27, 12, 32, -111, 66],
//     [15, 17, 23, 226, 28, -28, -226, -23, -17],
//     [16, 24, 27, 299, 30, 29, 32, 31, 88]
//   ]
// }
