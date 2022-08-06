package main

func StaircaseTraversal(height int, maxSteps int) int {
	return traverseToTop(height, maxSteps, map[int]int{0: 1, 1: 1})
}

func traverseToTop(height, maxSteps int, cache map[int]int) int {
	if ways, found := cache[height]; found {
		return ways
	}

	nWays := 0
	for step := 1; step < min(maxSteps, height)+1; step++ {
		nWays += traverseToTop(height-step, maxSteps, cache)
	}
	cache[height] = nWays
	return nWays
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func StaircaseTraversalIter(height, maxSteps int) int {
	waysToTop := make([]int, height+1)
	waysToTop[0], waysToTop[1] = 1, 1
	for h := 2; h < height+1; h++ {
		for step := 1; step <= maxSteps && step <= h; step++ {
			waysToTop[h] = waysToTop[h] + waysToTop[h-step]
		}
	}
	return waysToTop[height]
}

func StaircaseTraversalSlidingWindow(height, maxSteps int) int {
	nWays := 0
	waysToTop := make([]int, height+1)
	waysToTop[0] = 1
	for h := 1; h < height+1; h++ {
		if prevWindowStart := h - maxSteps - 1; prevWindowStart >= 0 {
			nWays -= waysToTop[prevWindowStart]
		}
		nWays += waysToTop[h-1]
		waysToTop[h] = nWays
	}
	return waysToTop[height]
}

// Test Case 1
// {
//   "height": 4,
//   "maxSteps": 2
// }
// Test Case 2
// {
//   "height": 10,
//   "maxSteps": 1
// }
// Test Case 3
// {
//   "height": 10,
//   "maxSteps": 2
// }
// Test Case 4
// {
//   "height": 4,
//   "maxSteps": 3
// }
// Test Case 5
// {
//   "height": 1,
//   "maxSteps": 1
// }
// Test Case 6
// {
//   "height": 5,
//   "maxSteps": 2
// }
// Test Case 7
// {
//   "height": 4,
//   "maxSteps": 4
// }
// Test Case 8
// {
//   "height": 6,
//   "maxSteps": 2
// }
// Test Case 9
// {
//   "height": 100,
//   "maxSteps": 1
// }
// Test Case 10
// {
//   "height": 15,
//   "maxSteps": 5
// }
// Test Case 11
// {
//   "height": 7,
//   "maxSteps": 2
// }
// Test Case 12
// {
//   "height": 6,
//   "maxSteps": 3
// }
// Test Case 13
// {
//   "height": 3,
//   "maxSteps": 2
// }
