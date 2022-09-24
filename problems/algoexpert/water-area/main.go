package main

func WaterArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	leftIdx := 0
	rightIdx := len(heights) - 1
	leftMax := heights[leftIdx]
	rightMax := heights[rightIdx]
	area := 0
	for leftIdx < rightIdx {
		if heights[leftIdx] < heights[rightIdx] {
			leftIdx++
			leftMax = max(leftMax, heights[leftIdx])
			area += leftMax - heights[leftIdx]
			continue
		}

		rightIdx--
		rightMax = max(rightMax, heights[rightIdx])
		area += rightMax - heights[rightIdx]
	}
	return area
}

func WaterArea2(heights []int) int {
	maxes := make([]int, len(heights))
	leftmax := 0
	for i, h := range heights {
		maxes[i], leftmax = leftmax, max(leftmax, h)
	}

	rightmax := 0
	for i := range heights {
		j := len(heights) - i - 1
		h := heights[j]
		minheight := min(rightmax, maxes[j])
		maxes[j] = 0
		if h < minheight {
			maxes[j] = minheight - h
		}
		rightmax = max(rightmax, h)
	}

	sum := 0
	for _, n := range maxes {
		sum += n
	}
	return sum
}

func min(first int, rest ...int) int {
	m := first
	for _, n := range rest {
		if n < m {
			m = n
		}
	}
	return m
}

func max(first int, rest ...int) int {
	m := first
	for _, n := range rest {
		if n > m {
			m = n
		}
	}
	return m
}

// Test Case 1
// {
//   "heights": [0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3]
// }
// Test Case 2
// {
//   "heights": []
// }
// Test Case 3
// {
//   "heights": [0, 0, 0, 0, 0]
// }
// Test Case 4
// {
//   "heights": [0, 1, 0, 0, 0]
// }
// Test Case 5
// {
//   "heights": [0, 1, 1, 0, 0]
// }
// Test Case 6
// {
//   "heights": [0, 1, 2, 1, 1]
// }
// Test Case 7
// {
//   "heights": [0, 1, 0, 1, 0]
// }
// Test Case 8
// {
//   "heights": [0, 1, 0, 1, 0, 2, 0, 3]
// }
// Test Case 9
// {
//   "heights": [0, 8, 0, 0, 10, 0, 0, 10, 0, 0, 1, 1, 0, 3]
// }
// Test Case 10
// {
//   "heights": [0, 100, 0, 0, 10, 1, 1, 10, 1, 0, 1, 1, 0, 100]
// }
// Test Case 11
// {
//   "heights": [0, 100, 0, 0, 10, 1, 1, 10, 1, 0, 1, 1, 0, 0]
// }
// Test Case 12
// {
//   "heights": [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8]
// }
// Test Case 13
// {
//   "heights": [8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1]
// }
// Test Case 14
// {
//   "heights": [1, 8, 6, 2, 5, 4, 8, 3, 7]
// }
