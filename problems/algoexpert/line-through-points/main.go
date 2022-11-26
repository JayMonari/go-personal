package main

import (
	"fmt"
)

func LineThroughPoints(points [][]int) int {
	maxNumberOfPointsOnLine := 1

	for i, p1 := range points {
		slopes := map[string]int{}
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			rise, run := calculateSlope(p1, p2)
			slopeKey := hash(rise, run)
			if slopes[slopeKey] == 0 {
				slopes[slopeKey] = 1
			}
			slopes[slopeKey]++
		}
		maxNumberOfPointsOnLine = max(maxNumberOfPointsOnLine, maxSlope(slopes))
	}
	return maxNumberOfPointsOnLine
}

func calculateSlope(p1, p2 []int) (int, int) {
	p1x, p1y := p1[0], p1[1]
	p2x, p2y := p2[0], p2[1]

	if p1x == p2x {
		return 1, 0
	}

	xDiff := p1x - p2x
	yDiff := p1y - p2y
	gcd := greatestCommonDivisor(abs(xDiff), abs(yDiff))
	xDiff = xDiff / gcd
	yDiff = yDiff / gcd
	if xDiff < 0 {
		xDiff = -xDiff
		yDiff = -yDiff
	}
	return yDiff, xDiff
}

func greatestCommonDivisor(num1, num2 int) int {
	a, b := num1, num2
	for {
		switch {
		case a == 0:
			return b
		case b == 0:
			return a
		}
		a, b = b, a%b
	}
}

func hash(numerator, denominator int) string {
	return fmt.Sprintf("%d:%d", numerator, denominator)
}

func maxSlope(slopes map[string]int) int {
	cmax := 0
	for _, s := range slopes {
		cmax = max(s, cmax)
	}
	return cmax
}

func max(a, b int) int {
	if a > b {
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

// Test Case 1
//
// {
//   "points": [
//     [1, 1],
//     [2, 2],
//     [3, 3],
//     [0, 4],
//     [-2, 6],
//     [4, 0],
//     [2, 1]
//   ]
// }
//
// Test Case 2
//
// {
//   "points": [
//     [3, 3],
//     [0, 4],
//     [-2, 6],
//     [4, 0],
//     [2, 1],
//     [3, 4],
//     [5, 6],
//     [0, 0]
//   ]
// }
//
// Test Case 3
//
// {
//   "points": [
//     [1, 4],
//     [3, 5],
//     [7, 1],
//     [5, 4],
//     [4, 5],
//     [9, 2],
//     [1, 3],
//     [2, 8]
//   ]
// }
//
// Test Case 4
//
// {
//   "points": [
//     [1, 4],
//     [4, 1],
//     [3, 3]
//   ]
// }
//
// Test Case 5
//
// {
//   "points": [
//     [0, 0]
//   ]
// }
//
// Test Case 6
//
// {
//   "points": [
//     [1, 4],
//     [4, 1],
//     [1, 1],
//     [4, 4],
//     [2, 3],
//     [3, 2],
//     [3, 3],
//     [2, 2],
//     [0, 3]
//   ]
// }
//
// Test Case 7
//
// {
//   "points": [
//     [1, 4],
//     [4, 1],
//     [1, 1],
//     [4, 4],
//     [2, 3],
//     [3, 2],
//     [3, 3],
//     [2, 2],
//     [0, 3],
//     [5, 3],
//     [3, -1],
//     [2, -3],
//     [1, -5]
//   ]
// }
//
// Test Case 8
//
// {
//   "points": [
//     [-1, -1],
//     [-3, -1],
//     [-4, -1],
//     [1, 1],
//     [4, 1]
//   ]
// }
//
// Test Case 9
//
// {
//   "points": [
//     [1, 1],
//     [1, 2],
//     [1, 3],
//     [1, 4],
//     [1, 5],
//     [2, 1],
//     [2, 2],
//     [2, 3],
//     [2, 4],
//     [2, 5],
//     [3, 1],
//     [3, 2],
//     [3, 4],
//     [3, 5],
//     [4, 1],
//     [4, 2],
//     [4, 3],
//     [4, 4],
//     [4, 5],
//     [5, 1],
//     [5, 2],
//     [5, 3],
//     [5, 4],
//     [5, 5],
//     [6, 6],
//     [2, 6]
//   ]
// }
//
// Test Case 10
//
// {
//   "points": [
//     [1, 1],
//     [1, 2],
//     [1, 3],
//     [1, 4],
//     [1, 5],
//     [2, 1],
//     [2, 2],
//     [2, 4],
//     [2, 5],
//     [4, 1],
//     [4, 2],
//     [4, 4],
//     [4, 5],
//     [5, 1],
//     [5, 2],
//     [5, 4],
//     [5, 5],
//     [6, 6],
//     [2, 6],
//     [-1, -1],
//     [0, 0],
//     [-2, -2]
//   ]
// }
//
// Test Case 11
//
// {
//   "points": [
//     [-78, -9],
//     [67, 87],
//     [46, 87],
//     [4, 5],
//     [9, 83],
//     [34, 47]
//   ]
// }
//
// Test Case 12
//
// {
//   "points": [
//     [1000000001, 1],
//     [1, 1],
//     [0, 0]
//   ]
// }
