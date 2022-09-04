package main

// O(2^(n+m))
func NumberOfWaysToTraverseGraph(width int, height int) int {
	if width == 1 || height == 1 {
		return 1
	}
	return NumberOfWaysToTraverseGraph(width-1, height) + NumberOfWaysToTraverseGraph(width, height-1)
}

// O(n * m)
func NumberOfWaysToTraverseGraphDP(width int, height int) int {
	ways := make([][]int, height+1)
	for i := range ways {
		ways[i] = make([]int, width+1)
	}
	for wIdx := 1; wIdx < width+1; wIdx++ {
		for hIdx := 1; hIdx < height+1; hIdx++ {
			if wIdx == 1 || hIdx == 1 {
				ways[hIdx][wIdx] = 1
				continue
			}
			ways[hIdx][wIdx] = ways[hIdx][wIdx-1] + ways[hIdx-1][wIdx]
		}
	}
	return ways[height][width]
}

// Test Case 1
// {
//   "width": 4,
//   "height": 3
// }
// Test Case 2
// {
//   "width": 3,
//   "height": 2
// }
// Test Case 3
// {
//   "width": 2,
//   "height": 3
// }
// Test Case 4
// {
//   "width": 5,
//   "height": 5
// }
// Test Case 5
// {
//   "width": 5,
//   "height": 6
// }
// Test Case 6
// {
//   "width": 7,
//   "height": 5
// }
// Test Case 7
// {
//   "width": 10,
//   "height": 2
// }
// Test Case 8
// {
//   "width": 11,
//   "height": 2
// }
// Test Case 9
// {
//   "width": 5,
//   "height": 9
// }
// Test Case 10
// {
//   "width": 6,
//   "height": 7
// }
// Test Case 11
// {
//   "width": 8,
//   "height": 5
// }
// Test Case 12
// {
//   "width": 2,
//   "height": 2
// }
// Test Case 13
// {
//   "width": 2,
//   "height": 1
// }
// Test Case 14
// {
//   "width": 1,
//   "height": 2
// }
// Test Case 15
// {
//   "width": 3,
//   "height": 3
// }
