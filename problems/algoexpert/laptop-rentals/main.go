package main

import (
	"sort"
)

func LaptopRentals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	startTimes, endTimes := make([]int, len(intervals)), make([]int, len(intervals))
	for i, in := range intervals {
		startTimes[i], endTimes[i] = in[0], in[1]
	}
	sort.Ints(startTimes)
	sort.Ints(endTimes)

	usedLaptops := 0
	sIdx, eIdx := 0, 0
	for sIdx < len(intervals) {
		if startTimes[sIdx] >= endTimes[eIdx] {
			usedLaptops--
			eIdx++
		}

		usedLaptops++
		sIdx++
	}
	return usedLaptops
}

// Test Case 1
// {
//   "times": [
//     [0, 2],
//     [1, 4],
//     [4, 6],
//     [0, 4],
//     [7, 8],
//     [9, 11],
//     [3, 10]
//   ]
// }
// Test Case 2
// {
//   "times": [
//     [0, 4],
//     [2, 3],
//     [2, 3],
//     [2, 3]
//   ]
// }
// Test Case 3
// {
//   "times": [
//     [1, 5],
//     [5, 6],
//     [6, 7],
//     [7, 9]
//   ]
// }
// Test Case 4
// {
//   "times": [
//     [0, 4]
//   ]
// }
// Test Case 5
// {
//   "times": []
// }
// Test Case 6
// {
//   "times": [
//     [0, 5],
//     [2, 4],
//     [4, 7],
//     [5, 7],
//     [9, 20],
//     [3, 15],
//     [6, 10]
//   ]
// }
// Test Case 7
// {
//   "times": [
//     [10, 20],
//     [0, 5],
//     [5, 10],
//     [10, 15]
//   ]
// }
// Test Case 8
// {
//   "times": [
//     [0, 5],
//     [3, 8],
//     [4, 10],
//     [7, 11],
//     [6, 10]
//   ]
// }
// Test Case 9
// {
//   "times": [
//     [0, 5],
//     [1, 4],
//     [2, 3],
//     [3, 8],
//     [7, 9],
//     [11, 20],
//     [0, 20],
//     [3, 10]
//   ]
// }
// Test Case 10
// {
//   "times": [
//     [10, 20],
//     [5, 15],
//     [0, 6],
//     [0, 20],
//     [21, 22],
//     [0, 1],
//     [2, 5]
//   ]
// }
// Test Case 11
// {
//   "times": [
//     [0, 10],
//     [1, 9],
//     [2, 8],
//     [3, 7],
//     [4, 6],
//     [5, 6]
//   ]
// }
// Test Case 12
// {
//   "times": [
//     [0, 20],
//     [0, 10],
//     [1, 9],
//     [2, 8],
//     [3, 7],
//     [4, 6],
//     [5, 6],
//     [10, 15],
//     [11, 12]
//   ]
// }
// Test Case 13
// {
//   "times": [
//     [5, 10],
//     [1, 2],
//     [1, 2],
//     [1, 2],
//     [3, 5],
//     [4, 5]
//   ]
// }
// Test Case 14
// {
//   "times": [
//     [1, 3],
//     [2, 5],
//     [4, 5],
//     [0, 20],
//     [1, 10],
//     [10, 20],
//     [11, 15],
//     [12, 13],
//     [0, 1],
//     [0, 2]
//   ]
// }
// Test Case 15
// {
//   "times": [
//     [5, 6],
//     [4, 5],
//     [3, 4],
//     [2, 3],
//     [1, 2]
//   ]
// }
