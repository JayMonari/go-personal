package main

func SpiralTraverse(mat [][]int) []int {
	result := []int{}
	for rStart, cStart, rEnd, cEnd := 0, 0, len(mat)-1, len(mat[0])-1; rStart <= rEnd &&
		cStart <= cEnd; {

		// left->right
		for col := cStart; col <= cEnd; col++ {
			result = append(result, mat[rStart][col])
		}
		// up->down
		for row := rStart + 1; row <= rEnd; row++ {
			result = append(result, mat[row][cEnd])
		}
		// right->left
		for col := cEnd - 1; col >= cStart; col-- {
			if rStart == rEnd {
				break
			}
			result = append(result, mat[rEnd][col])
		}
		// down->up
		for row := rEnd - 1; row > rStart; row-- {
			if cStart == cEnd {
				break
			}
			result = append(result, mat[row][cStart])
		}
		rStart, cStart, rEnd, cEnd = rStart+1, cStart+1, rEnd-1, cEnd-1
	}
	return result
}

// Test Case 1
// {
//   "array": [
//     [1, 2, 3, 4],
//     [12, 13, 14, 5],
//     [11, 16, 15, 6],
//     [10, 9, 8, 7]
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
//     [1, 2],
//     [4, 3]
//   ]
// }
// Test Case 4
// {
//   "array": [
//     [1, 2, 3],
//     [8, 9, 4],
//     [7, 6, 5]
//   ]
// }
// Test Case 5
// {
//   "array": [
//     [19, 32, 33, 34, 25, 8],
//     [16, 15, 14, 13, 12, 11],
//     [18, 31, 36, 35, 26, 9],
//     [1, 2, 3, 4, 5, 6],
//     [20, 21, 22, 23, 24, 7],
//     [17, 30, 29, 28, 27, 10]
//   ]
// }
// Test Case 6
// {
//   "array": [
//     [4, 2, 3, 6, 7, 8, 1, 9, 5, 10],
//     [12, 19, 15, 16, 20, 18, 13, 17, 11, 14]
//   ]
// }
// Test Case 7
// {
//   "array": [
//     [27, 12, 35, 26],
//     [25, 21, 94, 11],
//     [19, 96, 43, 56],
//     [55, 36, 10, 18],
//     [96, 83, 31, 94],
//     [93, 11, 90, 16]
//   ]
// }
// Test Case 8
// {
//   "array": [
//     [1, 2, 3, 4],
//     [10, 11, 12, 5],
//     [9, 8, 7, 6]
//   ]
// }
// Test Case 9
// {
//   "array": [
//     [1, 2, 3],
//     [12, 13, 4],
//     [11, 14, 5],
//     [10, 15, 6],
//     [9, 8, 7]
//   ]
// }
// Test Case 10
// {
//   "array": [
//     [1, 11],
//     [2, 12],
//     [3, 13],
//     [4, 14],
//     [5, 15],
//     [6, 16],
//     [7, 17],
//     [8, 18],
//     [9, 19],
//     [10, 20]
//   ]
// }
// Test Case 11
// {
//   "array": [
//     [1, 3, 2, 5, 4, 7, 6]
//   ]
// }
// Test Case 12
// {
//   "array": [
//     [1],
//     [3],
//     [2],
//     [5],
//     [4],
//     [7],
//     [6]
//   ]
// }
