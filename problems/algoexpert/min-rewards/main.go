package main

func MinRewards(scores []int) int {
	rewards := make([]int, len(scores))
	for i := range rewards {
		rewards[i] = 1
	}
	for i := 1; i < len(scores); i++ {
		if scores[i] > scores[i-1] {
			rewards[i] = rewards[i-1] + 1
		}
	}
	for i := len(scores) - 2; i >= 0; i-- {
		if scores[i] > scores[i+1] {
			rewards[i] = max(rewards[i], rewards[i+1]+1)
		}
	}
	sum := 0
	for _, n := range rewards {
		sum += n
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test Case 1
// {
//   "scores": [8, 4, 2, 1, 3, 6, 7, 9, 5]
// }
// Test Case 2
// {
//   "scores": [1]
// }
// Test Case 3
// {
//   "scores": [5, 10]
// }
// Test Case 4
// {
//   "scores": [10, 5]
// }
// Test Case 5
// {
//   "scores": [4, 2, 1, 3]
// }
// Test Case 6
// {
//   "scores": [0, 4, 2, 1, 3]
// }
// Test Case 7
// {
//   "scores": [2, 20, 13, 12, 11, 8, 4, 3, 1, 5, 6, 7, 9, 0]
// }
// Test Case 8
// {
//   "scores": [2, 1, 4, 3, 6, 5, 8, 7, 10, 9]
// }
// Test Case 9
// {
//   "scores": [800, 400, 20, 10, 30, 61, 70, 90, 17, 21, 22, 13, 12, 11, 8, 4, 2, 1, 3, 6, 7, 9, 0, 68, 55, 67, 57, 60, 51, 661, 50, 65, 53]
// }
