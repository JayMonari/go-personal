package main

type state uint8

const (
	unvisited state = iota
	inStack
	finished
)

func CycleInGraph(edges [][]int) bool {
	states := make([]state, len(edges))
	for node := 0; node < len(edges); node++ {
		if states[node] != unvisited {
			continue
		}
		if hasCycle(node, edges, states) {
			return true
		}
	}
	return false
}

func hasCycle(node int, edges [][]int, states []state) bool {
	states[node] = inStack
	for _, currNode := range edges[node] {
		switch states[currNode] {
		case inStack:
			return true
		case finished:
			continue
		}
		if hasCycle(currNode, edges, states) {
			return true
		}
	}
	states[node] = finished
	return false
}

// Test Case 1
// {
//   "edges": [
//     [1, 3],
//     [2, 3, 4],
//     [0],
//     [],
//     [2, 5],
//     []
//   ]
// }
// Test Case 2
// {
//   "edges": [
//     [1, 2],
//     [2],
//     []
//   ]
// }
// Test Case 3
// {
//   "edges": [
//     [1, 2],
//     [2],
//     [1]
//   ]
// }
// Test Case 4
// {
//   "edges": [
//     [1, 2],
//     [2],
//     [1, 3],
//     [3]
//   ]
// }
// Test Case 5
// {
//   "edges": [
//     [],
//     [0, 2],
//     [0, 3],
//     [0, 4],
//     [0, 5],
//     [0]
//   ]
// }
// Test Case 6
// {
//   "edges": [
//     [0]
//   ]
// }
// Test Case 7
// {
//   "edges": [
//     [8],
//     [0, 2],
//     [0, 3],
//     [0, 4],
//     [0, 5],
//     [0],
//     [7],
//     [8],
//     [6]
//   ]
// }
// Test Case 8
// {
//   "edges": [
//     [1],
//     [2, 3, 4, 5, 6, 7],
//     [],
//     [2, 7],
//     [5],
//     [],
//     [4],
//     []
//   ]
// }
// Test Case 9
// {
//   "edges": [
//     [1],
//     [2, 3, 4, 5, 6, 7],
//     [],
//     [2, 7],
//     [5],
//     [],
//     [4],
//     [0]
//   ]
// }
// Test Case 10
// {
//   "edges": [
//     [0],
//     [1]
//   ]
// }
// Test Case 11
// {
//   "edges": [
//     [1, 2],
//     [2],
//     []
//   ]
// }
// Test Case 12
// {
//   "edges": [
//     [],
//     [0, 3],
//     [0],
//     [1, 2]
//   ]
// }
