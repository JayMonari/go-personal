package main

const (
	water  = 0
	island = 1
)

type coordinate [2]int

func RemoveIslands(matrix [][]int) [][]int {
	borderIslands := map[coordinate]struct{}{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			switch {
			case !(i == 0 || j == 0 || i == len(matrix)-1 || j == len(matrix[i])-1):
				continue
			case matrix[i][j] != 1:
				continue
			}

			q := []coordinate{{i, j}}
			for len(q) > 0 {
				coord := q[0]
				q = q[1:]
				if _, visited := borderIslands[coord]; visited {
					continue
				}
				borderIslands[coord] = struct{}{}
				for _, n := range getNeighbors(matrix, coord) {
					if matrix[n[0]][n[1]] != island {
						continue
					}
					q = append(q, n)
				}
			}
		}
	}

	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			if _, is := borderIslands[coordinate{i, j}]; is {
				continue
			}
			matrix[i][j] = water
		}
	}
	return matrix
}

func getNeighbors(matrix [][]int, c coordinate) []coordinate {
	neighbors := []coordinate{}
	i, j := c[0], c[1]
	if i > 0 {
		neighbors = append(neighbors, coordinate{i - 1, j})
	}
	if j > 0 {
		neighbors = append(neighbors, coordinate{i, j - 1})
	}
	if i < len(matrix)-1 {
		neighbors = append(neighbors, coordinate{i + 1, j})
	}
	if j < len(matrix[i])-1 {
		neighbors = append(neighbors, coordinate{i, j + 1})
	}
	return neighbors
}

// Test Case 1
// {
//   "matrix": [
//     [1, 0, 0, 0, 0, 0],
//     [0, 1, 0, 1, 1, 1],
//     [0, 0, 1, 0, 1, 0],
//     [1, 1, 0, 0, 1, 0],
//     [1, 0, 1, 1, 0, 0],
//     [1, 0, 0, 0, 0, 1]
//   ]
// }
// Test Case 2
// {
//   "matrix": [
//     [1, 0, 0, 0, 1],
//     [0, 1, 0, 1, 0],
//     [0, 0, 1, 0, 0],
//     [0, 1, 0, 1, 0],
//     [1, 0, 0, 0, 1]
//   ]
// }
// Test Case 3
// {
//   "matrix": [
//     [1, 0, 0, 1, 0],
//     [0, 1, 0, 1, 0],
//     [0, 0, 1, 1, 0]
//   ]
// }
// Test Case 4
// {
//   "matrix": [
//     [1, 1, 1, 1, 1],
//     [1, 0, 0, 0, 1],
//     [1, 0, 1, 0, 1],
//     [1, 0, 0, 0, 1],
//     [1, 0, 1, 0, 1],
//     [1, 0, 1, 0, 1],
//     [1, 0, 1, 1, 1],
//     [1, 0, 1, 0, 1]
//   ]
// }
// Test Case 5
// {
//   "matrix": [
//     [0, 0, 0, 0, 0],
//     [0, 1, 1, 1, 0],
//     [0, 1, 1, 1, 0],
//     [0, 1, 1, 1, 0],
//     [0, 0, 0, 0, 0]
//   ]
// }
// Test Case 6
// {
//   "matrix": [
//     [1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1]
//   ]
// }
// Test Case 7
// {
//   "matrix": [
//     [1]
//   ]
// }
// Test Case 8
// {
//   "matrix": [
//     [1, 0, 0, 0, 1, 0, 0, 0],
//     [1, 0, 1, 0, 1, 0, 1, 0],
//     [1, 1, 0, 1, 0, 0, 1, 0],
//     [1, 1, 0, 1, 1, 0, 1, 0],
//     [1, 0, 0, 0, 1, 0, 0, 0]
//   ]
// }
// Test Case 9
// {
//   "matrix": [
//     [1, 1, 1, 1, 1],
//     [1, 0, 0, 0, 1],
//     [1, 0, 1, 0, 1],
//     [1, 0, 0, 0, 1],
//     [1, 1, 1, 1, 1]
//   ]
// }
// Test Case 10
// {
//   "matrix": [
//     [1, 0, 1, 0, 1],
//     [0, 0, 1, 0, 0],
//     [1, 1, 0, 1, 1],
//     [0, 0, 1, 0, 0],
//     [1, 0, 1, 0, 1]
//   ]
// }
// Test Case 11
// {
//   "matrix": [
//     [0, 0, 0, 0, 0],
//     [0, 0, 1, 0, 0],
//     [0, 1, 1, 1, 0],
//     [0, 0, 1, 0, 0],
//     [0, 0, 0, 0, 0]
//   ]
// }
// Test Case 12
// {
//   "matrix": [
//     [1, 0, 1, 0, 1, 0],
//     [0, 1, 0, 1, 0, 1],
//     [1, 0, 1, 0, 1, 0],
//     [0, 1, 0, 1, 0, 1],
//     [1, 0, 1, 0, 1, 0],
//     [0, 1, 0, 1, 0, 1]
//   ]
// }
// Test Case 13
// {
//   "matrix": [
//     [1, 0, 1, 1, 1, 0],
//     [1, 1, 0, 1, 0, 1],
//     [1, 0, 1, 0, 1, 0],
//     [0, 1, 0, 1, 0, 1],
//     [1, 0, 1, 0, 1, 0],
//     [0, 1, 1, 1, 0, 1]
//   ]
// }
// Test Case 14
// {
//   "matrix": [
//     [0, 1, 0],
//     [0, 1, 0],
//     [1, 0, 0]
//   ]
// }
// Test Case 15
// {
//   "matrix": [
//     [1, 1],
//     [1, 1]
//   ]
// }
// Test Case 16
// {
//   "matrix": [
//     [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1],
//     [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//     [0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
//   ]
// }
