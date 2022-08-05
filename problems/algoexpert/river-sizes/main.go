package main

const (
	ground = 0
	water  = 1
)

type coordinate [2]int

func RiverSizes(matrix [][]int) []int {
	sizes := []int{}
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}
	for i := range matrix {
		for j := range matrix[i] {
			if visited[i][j] {
				continue
			}
			sizes = traverseNode(i, j, matrix, visited, sizes)
		}
	}
	return sizes
}

func traverseNode(i, j int, matrix [][]int, visited [][]bool, sizes []int) []int {
	currSize := 0
	nodesToExplore := []coordinate{{i, j}}
	for len(nodesToExplore) > 0 {
		i, j := nodesToExplore[0][0], nodesToExplore[0][1]
		nodesToExplore = nodesToExplore[1:]
		if visited[i][j] {
			continue
		}
		visited[i][j] = true
		if matrix[i][j] == ground {
			continue
		}
		currSize++
		nodesToExplore = append(nodesToExplore, unvisitedNeighbors(i, j, matrix, visited)...)
	}
	if currSize > 0 {
		sizes = append(sizes, currSize)
	}
	return sizes
}

func unvisitedNeighbors(i, j int, matrix [][]int, visited [][]bool) []coordinate {
	unvisited := []coordinate{}
	if i > 0 && !visited[i-1][j] {
		unvisited = append(unvisited, coordinate{i - 1, j})
	}
	if j > 0 && !visited[i][j-1] {
		unvisited = append(unvisited, coordinate{i, j - 1})
	}
	if i < len(matrix)-1 && !visited[i+1][j] {
		unvisited = append(unvisited, coordinate{i + 1, j})
	}
	if j < len(matrix[0])-1 && !visited[i][j+1] {
		unvisited = append(unvisited, coordinate{i, j + 1})
	}
	return unvisited
}

// Test Case 1
// {
//   "matrix": [
//     [1, 0, 0, 1, 0],
//     [1, 0, 1, 0, 0],
//     [0, 0, 1, 0, 1],
//     [1, 0, 1, 0, 1],
//     [1, 0, 1, 1, 0]
//   ]
// }
// Test Case 2
// {
//   "matrix": [
//     [0]
//   ]
// }
// Test Case 3
// {
//   "matrix": [
//     [1]
//   ]
// }
// Test Case 4
// {
//   "matrix": [
//     [1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0]
//   ]
// }
// Test Case 5
// {
//   "matrix": [
//     [1, 0, 0, 1],
//     [1, 0, 1, 0],
//     [0, 0, 1, 0],
//     [1, 0, 1, 0]
//   ]
// }
// Test Case 6
// {
//   "matrix": [
//     [1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0],
//     [1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0],
//     [0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 1],
//     [1, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0],
//     [1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 1]
//   ]
// }
// Test Case 7
// {
//   "matrix": [
//     [1, 0, 0, 0, 0, 0, 1],
//     [0, 1, 0, 0, 0, 1, 0],
//     [0, 0, 1, 0, 1, 0, 0],
//     [0, 0, 0, 1, 0, 0, 0],
//     [0, 0, 1, 0, 1, 0, 0],
//     [0, 1, 0, 0, 0, 1, 0],
//     [1, 0, 0, 0, 0, 0, 1]
//   ]
// }
// Test Case 8
// {
//   "matrix": [
//     [1, 0, 0, 0, 0, 0, 1],
//     [0, 1, 0, 0, 0, 1, 0],
//     [0, 0, 1, 0, 1, 0, 0],
//     [0, 0, 1, 1, 1, 0, 0],
//     [0, 0, 1, 0, 1, 0, 0],
//     [0, 1, 0, 0, 0, 1, 0],
//     [1, 0, 0, 0, 0, 0, 1]
//   ]
// }
// Test Case 9
// {
//   "matrix": [
//     [0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0],
//     [0, 0, 0, 0, 0, 0, 0]
//   ]
// }
// Test Case 10
// {
//   "matrix": [
//     [1, 1, 1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, 1],
//     [1, 1, 1, 1, 1, 1, 1]
//   ]
// }
// Test Case 11
// {
//   "matrix": [
//     [1, 1, 0, 0, 0, 0, 1, 1],
//     [1, 0, 1, 1, 1, 1, 0, 1],
//     [0, 1, 1, 0, 0, 0, 1, 1]
//   ]
// }
// Test Case 12
// {
//   "matrix": [
//     [1, 1, 0],
//     [1, 0, 1],
//     [1, 1, 1],
//     [1, 1, 0],
//     [1, 0, 1],
//     [0, 1, 0],
//     [1, 0, 0],
//     [1, 0, 0],
//     [0, 0, 0],
//     [1, 0, 0],
//     [1, 0, 1],
//     [1, 1, 1]
//   ]
// }
