package main

func LargestRectangleUnderSkyline(buildings []int) int {
	idxs := []int{}
	area := 0
	extendedBuildings := append(buildings, 0)
	for i := range extendedBuildings {
		height := extendedBuildings[i]
		for len(idxs) != 0 && buildings[idxs[len(idxs)-1]] >= height {
			idx := idxs[len(idxs)-1]
			idxs = idxs[:len(idxs)-1]
			pillarHeight := buildings[idx]
			width := i
			if len(idxs) != 0 {
				width = i - idxs[len(idxs)-1] - 1
			}
			area = max(width*pillarHeight, area)
		}
		idxs = append(idxs, i)
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test Case 1
//
// {
//   "buildings": [1, 3, 3, 2, 4, 1, 5, 3, 2]
// }
//
// Test Case 2
//
// {
//   "buildings": [4, 4, 4, 2, 2, 1]
// }
//
// Test Case 3
//
// {
//   "buildings": [1, 3, 3, 2, 4, 1, 5, 3]
// }
//
// Test Case 4
//
// {
//   "buildings": [5, 5, 2, 2, 4, 1]
// }
//
// Test Case 5
//
// {
//   "buildings": [1, 2, 3, 4, 5, 11]
// }
//
// Test Case 6
//
// {
//   "buildings": [25, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
// }
//
// Test Case 7
//
// {
//   "buildings": [20, 2, 2, 2, 2, 2, 10, 5, 5, 5, 4, 4]
// }
//
// Test Case 8
//
// {
//   "buildings": [5, 10, 5, 15, 10, 25]
// }
//
// Test Case 9
//
// {
//   "buildings": [1, 1, 1, 1]
// }
//
// Test Case 10
//
// {
//   "buildings": [10, 21]
// }
//
// Test Case 11
//
// {
//   "buildings": [11, 21]
// }
//
// Test Case 12
//
// {
//   "buildings": [3, 3, 3, 4, 4, 4, 1, 3, 1, 2, 8, 9, 1]
// }
//
// Test Case 13
//
// {
//   "buildings": [5]
// }
//
// Test Case 14
//
// {
//   "buildings": [10, 1, 2, 3, 4, 5, 6, 7]
// }
//
// Test Case 15
//
// {
//   "buildings": [10, 1, 2, 3, 3, 5, 6, 7]
// }
//
// Test Case 16
//
// {
//   "buildings": []
// }
