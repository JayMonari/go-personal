package main

import (
	"sort"
)

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)
	redsCan, bluesCan := true, true
	for i := range redShirtHeights {
		if redShirtHeights[i] >= blueShirtHeights[i] {
			redsCan = false
		}
		if blueShirtHeights[i] >= redShirtHeights[i] {
			bluesCan = false
		}
	}
	return redsCan || bluesCan
}

// Test Case 1
// {
//   "redShirtHeights": [5, 8, 1, 3, 4],
//   "blueShirtHeights": [6, 9, 2, 4, 5]
// }
// Test Case 2
// {
//   "redShirtHeights": [6, 9, 2, 4, 5],
//   "blueShirtHeights": [5, 8, 1, 3, 4]
// }
// Test Case 3
// {
//   "redShirtHeights": [6, 9, 2, 4, 5, 1],
//   "blueShirtHeights": [5, 8, 1, 3, 4, 9]
// }
// Test Case 4
// {
//   "redShirtHeights": [6],
//   "blueShirtHeights": [6]
// }
// Test Case 5
// {
//   "redShirtHeights": [126],
//   "blueShirtHeights": [125]
// }
// Test Case 6
// {
//   "redShirtHeights": [1, 2, 3, 4, 5],
//   "blueShirtHeights": [2, 3, 4, 5, 6]
// }
// Test Case 7
// {
//   "redShirtHeights": [1, 1, 1, 1, 1, 1, 1, 1],
//   "blueShirtHeights": [5, 6, 7, 2, 3, 1, 2, 3]
// }
// Test Case 8
// {
//   "redShirtHeights": [1, 1, 1, 1, 1, 1, 1, 1],
//   "blueShirtHeights": [2, 2, 2, 2, 2, 2, 2, 2]
// }
// Test Case 9
// {
//   "redShirtHeights": [125],
//   "blueShirtHeights": [126]
// }
// Test Case 10
// {
//   "redShirtHeights": [19, 2, 4, 6, 2, 3, 1, 1, 4],
//   "blueShirtHeights": [21, 5, 4, 4, 4, 4, 4, 4, 4]
// }
// Test Case 11
// {
//   "redShirtHeights": [19, 19, 21, 1, 1, 1, 1, 1],
//   "blueShirtHeights": [20, 5, 4, 4, 4, 4, 4, 4]
// }
// Test Case 12
// {
//   "redShirtHeights": [3, 5, 6, 8, 2],
//   "blueShirtHeights": [2, 4, 7, 5, 1]
// }
// Test Case 13
// {
//   "redShirtHeights": [3, 5, 6, 8, 2, 1],
//   "blueShirtHeights": [2, 4, 7, 5, 1, 6]
// }
// Test Case 14
// {
//   "redShirtHeights": [4, 5],
//   "blueShirtHeights": [5, 4]
// }
// Test Case 15
// {
//   "redShirtHeights": [5, 6],
//   "blueShirtHeights": [5, 4]
// }
