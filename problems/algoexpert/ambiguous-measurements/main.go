package main

import (
	"strconv"
)

func AmbiguousMeasurements(measuringCups [][]int, low int, high int) bool {
	cache := map[string]bool{}
	return canMeasureInRange(measuringCups, low, high, cache)
}

func canMeasureInRange(measuringCups [][]int, low, high int, cache map[string]bool) bool {
	ckey := strconv.Itoa(low) + ":" + strconv.Itoa(high)
	if val, found := cache[ckey]; found {
		return val
	}

	if low <= 0 && high <= 0 {
		return false
	}

	canMeasure := false
	for _, cup := range measuringCups {
		cupLow, cupHigh := cup[0], cup[1]
		if low <= cupLow && cupHigh <= high {
			canMeasure = true
			break
		}
		canMeasure = canMeasureInRange(
			measuringCups,
			max(0, low-cupLow),
			max(0, high-cupHigh),
			cache,
		)
		if canMeasure {
			break
		}
	}
	cache[ckey] = canMeasure
	return canMeasure
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test Case 1
// {
//   "measuringCups": [
//     [200, 210],
//     [450, 465],
//     [800, 850]
//   ],
//   "low": 2100,
//   "high": 2300
// }
// Test Case 2
// {
//   "measuringCups": [
//     [200, 210]
//   ],
//   "low": 10,
//   "high": 20
// }
// Test Case 3
// {
//   "measuringCups": [
//     [230, 240],
//     [290, 310],
//     [500, 515]
//   ],
//   "low": 2100,
//   "high": 2300
// }
// Test Case 4
// {
//   "measuringCups": [
//     [1, 3],
//     [2, 4],
//     [5, 6]
//   ],
//   "low": 100,
//   "high": 101
// }
// Test Case 5
// {
//   "measuringCups": [
//     [1, 3],
//     [2, 4],
//     [5, 6]
//   ],
//   "low": 100,
//   "high": 120
// }
// Test Case 6
// {
//   "measuringCups": [
//     [1, 3],
//     [2, 4],
//     [5, 6],
//     [10, 20]
//   ],
//   "low": 10,
//   "high": 12
// }
// Test Case 7
// {
//   "measuringCups": [
//     [1, 3],
//     [2, 4],
//     [5, 7],
//     [10, 20]
//   ],
//   "low": 10,
//   "high": 12
// }
// Test Case 8
// {
//   "measuringCups": [
//     [50, 60],
//     [100, 120],
//     [20, 40],
//     [10, 15],
//     [400, 500]
//   ],
//   "low": 1000,
//   "high": 1050
// }
// Test Case 9
// {
//   "measuringCups": [
//     [50, 65],
//     [100, 120],
//     [20, 40],
//     [10, 15],
//     [400, 500]
//   ],
//   "low": 1000,
//   "high": 1200
// }
// Test Case 10
// {
//   "measuringCups": [
//     [50, 65],
//     [100, 120],
//     [20, 40],
//     [10, 15],
//     [400, 500],
//     [300, 350],
//     [10, 25]
//   ],
//   "low": 3000,
//   "high": 3300
// }
// Test Case 11
// {
//   "measuringCups": [
//     [50, 60],
//     [100, 120],
//     [20, 40],
//     [400, 500]
//   ],
//   "low": 1000,
//   "high": 1050
// }
// Test Case 12
// {
//   "measuringCups": [
//     [50, 65]
//   ],
//   "low": 200,
//   "high": 200
// }
// Test Case 13
// {
//   "measuringCups": [
//     [50, 50]
//   ],
//   "low": 200,
//   "high": 200
// }
// Test Case 14
// {
//   "measuringCups": [
//     [50, 50],
//     [50, 51]
//   ],
//   "low": 200,
//   "high": 200
// }
// Test Case 15
// {
//   "measuringCups": [
//     [100, 150],
//     [1000, 2000]
//   ],
//   "low": 0,
//   "high": 1000
// }
// Test Case 16
// {
//   "measuringCups": [
//     [10, 20]
//   ],
//   "low": 10,
//   "high": 20
// }
// Test Case 17
// {
//   "measuringCups": [
//     [15, 18]
//   ],
//   "low": 10,
//   "high": 20
// }
// Test Case 18
// {
//   "measuringCups": [
//     [15, 22]
//   ],
//   "low": 10,
//   "high": 20
// }
