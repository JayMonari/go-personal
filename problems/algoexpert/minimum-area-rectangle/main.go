package main

import (
	"fmt"
	"math"
)

type pointSet map[string]struct{}

func MinimumAreaRectangle(points [][]int) int {
	pointSet := newPointSet(points)
	minimumAreaFound := math.MaxInt32

	for i := range points {
		p2x, p2y := points[i][0], points[i][1]
		for j := 0; j < i; j++ {
			p1x, p1y := points[j][0], points[j][1]
			pointsShareValue := p1x == p2x || p1y == p2y

			if pointsShareValue {
				continue
			}

			// If (p1x, p2y) and (p2x, p1y), exist we've found a rectangle.
			_, ok1 := pointSet[pointToString(p1x, p2y)]
			_, ok2 := pointSet[pointToString(p2x, p1y)]
			if ok1 && ok2 {
				currentArea := abs(p2x-p1x) * abs(p2y-p1y)
				minimumAreaFound = min(minimumAreaFound, currentArea)
			}
		}
	}

	if minimumAreaFound == math.MaxInt32 {
		return 0
	}
	return minimumAreaFound
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func newPointSet(points [][]int) pointSet {
	set := pointSet{}
	for _, p := range points {
		set[pointToString(p[0], p[1])] = struct{}{}
	}
	return set
}

func pointToString(x, y int) string { return fmt.Sprintf("%d:%d", x, y) }

// Test Case 1
//
// {
//   "points": [
//     [1, 5],
//     [5, 1],
//     [4, 2],
//     [2, 4],
//     [2, 2],
//     [1, 2],
//     [4, 5],
//     [2, 5],
//     [-1, -2]
//   ]
// }
//
// Test Case 2
//
// {
//   "points": [
//     [-4, 4],
//     [4, 4],
//     [4, -2],
//     [-4, -2],
//     [0, -2],
//     [4, 2],
//     [0, 2]
//   ]
// }
//
// Test Case 3
//
// {
//   "points": [
//     [-4, 4],
//     [4, 4],
//     [4, -2],
//     [-4, -2],
//     [0, -2],
//     [4, 2],
//     [0, 2],
//     [0, 4],
//     [2, 3],
//     [0, 3],
//     [2, 4]
//   ]
// }
//
// Test Case 4
//
// {
//   "points": [
//     [0, 0],
//     [4, 4],
//     [8, 8],
//     [0, 8]
//   ]
// }
//
// Test Case 5
//
// {
//   "points": [
//     [0, 0],
//     [4, 4],
//     [8, 8],
//     [0, 8],
//     [0, 4],
//     [6, 0],
//     [6, 4]
//   ]
// }
//
// Test Case 6
//
// {
//   "points": [
//     [0, 0],
//     [4, 4],
//     [8, 8],
//     [0, 8],
//     [0, 4],
//     [6, 0],
//     [6, 4],
//     [8, 0],
//     [8, 4],
//     [6, 2],
//     [2, 4],
//     [2, 0]
//   ]
// }
//
// Test Case 7
//
// {
//   "points": [
//     [0, 0],
//     [1, 1],
//     [2, 2],
//     [-1, -1],
//     [-2, -2],
//     [-1, 1],
//     [-2, 2],
//     [1, -1],
//     [2, -2]
//   ]
// }
//
// Test Case 8
//
// {
//   "points": [
//     [0, 1],
//     [0, 0],
//     [2, 1],
//     [2, 0],
//     [4, 0],
//     [4, 1],
//     [0, 2],
//     [2, 2],
//     [4, 2],
//     [6, 0],
//     [6, 1],
//     [6, 2],
//     [7, 1],
//     [7, 0]
//   ]
// }
//
// Test Case 9
//
// {
//   "points": [
//     [0, 1],
//     [0, 0],
//     [2, 1],
//     [2, 0],
//     [4, 0],
//     [4, 1],
//     [0, 2],
//     [2, 2],
//     [4, 2],
//     [6, 0],
//     [6, 1],
//     [6, 2],
//     [7, 1]
//   ]
// }
//
// Test Case 10
//
// {
//   "points": [
//     [100, 100],
//     [76, 67],
//     [-100, 100],
//     [65, 76],
//     [100, -100],
//     [3, 4],
//     [-100, -100],
//     [5, 6],
//     [78, 54],
//     [-87, 7],
//     [1, 4],
//     [4, 1],
//     [-1, 5]
//   ]
// }
//
// Test Case 11
//
// {
//   "points": []
// }
//
// Test Case 12
//
// {
//   "points": [
//     [1, 2],
//     [4, 2]
//   ]
// }
//
// Test Case 13
//
// {
//   "points": [
//     [2, 2],
//     [3, 2],
//     [4, 2]
//   ]
// }
