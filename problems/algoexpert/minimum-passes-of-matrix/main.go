package main

type coordinate [2]int

func MinimumPassesOfMatrix(matrix [][]int) int {
	nextQ := []coordinate{}
	for i := range matrix {
		for j := range matrix[i] {
			if num := matrix[i][j]; num > 0 {
				nextQ = append(nextQ, coordinate{i, j})
			}
		}
	}

	passes := 0
	for len(nextQ) > 0 {
		currQ := nextQ
		nextQ = []coordinate{}
		for len(currQ) > 0 {
			coord := currQ[0]
			currQ = currQ[1:]
			for _, c := range adjacentCoords(coord, matrix) {
				if matrix[c[0]][c[1]] < 0 {
					matrix[c[0]][c[1]] *= -1
					nextQ = append(nextQ, c)
				}
			}
		}
		passes++
	}
	// Check for all negatives
	for _, row := range matrix {
		for _, v := range row {
			if v < 0 {
				return -1
			}
		}
	}
	return passes - 1
}

func adjacentCoords(c coordinate, matrix [][]int) []coordinate {
	adj := []coordinate{}
	i, j := c[0], c[1]
	if i > 0 {
		adj = append(adj, coordinate{i - 1, j})
	}
	if j > 0 {
		adj = append(adj, coordinate{i, j - 1})
	}
	if i < len(matrix)-1 {
		adj = append(adj, coordinate{i + 1, j})
	}
	if j < len(matrix[i])-1 {
		adj = append(adj, coordinate{i, j + 1})
	}
	return adj
}

// Test Case 1
// {
//   "matrix": [
//     [0, -1, -3, 2, 0],
//     [1, -2, -5, -1, -3],
//     [3, 0, 0, -4, -1]
//   ]
// }
// Test Case 2
// {
//   "matrix": [
//     [1]
//   ]
// }
// Test Case 3
// {
//   "matrix": [
//     [1, 0, 0, -2, -3],
//     [-4, -5, -6, -2, -1],
//     [0, 0, 0, 0, -1],
//     [1, 2, 3, 0, -2]
//   ]
// }
// Test Case 4
// {
//   "matrix": [
//     [1, 0, 0, -2, -3],
//     [-4, -5, -6, -2, -1],
//     [0, 0, 0, 0, -1],
//     [1, 2, 3, 0, 3]
//   ]
// }
// Test Case 5
// {
//   "matrix": [
//     [1, 0, 0, -2, -3],
//     [-4, -5, -6, -2, -1],
//     [0, 0, 0, 0, -1],
//     [-1, 0, 3, 0, 3]
//   ]
// }
// Test Case 6
// {
//   "matrix": [
//     [-1]
//   ]
// }
// Test Case 7
// {
//   "matrix": [
//     [1, 2, 3],
//     [4, 5, 6]
//   ]
// }
// Test Case 8
// {
//   "matrix": [
//     [-1, -9, 0, -1, 0],
//     [-9, -4, -5, 4, -8],
//     [2, 0, 0, -3, 0],
//     [0, -17, -4, 2, -5]
//   ]
// }
// Test Case 9
// {
//   "matrix": [
//     [-2, -3, -4, -1, -9],
//     [-4, -3, -4, -1, -2],
//     [-6, -7, -2, -1, -1],
//     [0, 0, 0, 0, -3],
//     [1, -2, -3, -6, -1]
//   ]
// }
// Test Case 10
// {
//   "matrix": [
//     [-1, 2, 3],
//     [4, 5, 6]
//   ]
// }
// Test Case 11
// {
//   "matrix": [
//     [-1, 2, 3],
//     [4, -5, -6]
//   ]
// }
// Test Case 12
// {
//   "matrix": [
//     [-1, 0, 3],
//     [0, -5, -6]
//   ]
// }
// Test Case 13
// {
//   "matrix": [
//     [-1, 0, 3],
//     [0, -5, -6]
//   ]
// }
// Test Case 14
// {
//   "matrix": [
//     [0, 0, -1, -2],
//     [-2, -3, 4, -1],
//     [-2, -3, 1, -3],
//     [-14, -15, 2, 0],
//     [0, 0, 0, 0],
//     [1, -1, -1, -1]
//   ]
// }
// Test Case 15
// {
//   "matrix": [
//     [0, 0, -1, -2],
//     [-2, -3, 4, -1],
//     [-2, -3, 1, -3],
//     [-14, -15, 2, 0],
//     [0, 0, 0, 0],
//     [1, -1, -1, 1]
//   ]
// }
// Test Case 16
// {
//   "matrix": [
//     [-2, 0, -2, 1],
//     [-2, -1, -1, -1]
//   ]
// }
