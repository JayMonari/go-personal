package main

func Greedy(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	jumps := 0
	maxReach := nums[0]
	steps := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
		steps--
		if steps == 0 {
			jumps++
			steps = maxReach - i
		}
	}
	return jumps + 1
}

func MinNumberOfJumps(nums []int) int {
	jumps := make([]int, len(nums))
	for i := range jumps {
		jumps[i] = -1
	}
	jumps[0] = 0

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] >= i-j && (jumps[i] == -1 || jumps[j]+1 < jumps[i]) {
				jumps[i] = jumps[j] + 1
			}
		}
	}
	return jumps[len(jumps)-1]
}

// Test Case 1
// {
//   "array": [3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3]
// }
// Test Case 2
// {
//   "array": [1]
// }
// Test Case 3
// {
//   "array": [1, 1]
// }
// Test Case 4
// {
//   "array": [3, 1]
// }
// Test Case 5
// {
//   "array": [1, 1, 1]
// }
// Test Case 6
// {
//   "array": [2, 1, 1]
// }
// Test Case 7
// {
//   "array": [2, 1, 2, 3, 1]
// }
// Test Case 8
// {
//   "array": [2, 1, 2, 3, 1, 1, 1]
// }
// Test Case 9
// {
//   "array": [2, 1, 2, 2, 1, 1, 1]
// }
// Test Case 10
// {
//   "array": [3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3, 2, 6, 2, 1, 1, 1, 1]
// }
// Test Case 11
// {
//   "array": [3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3, 2, 3, 2, 1, 1, 1, 1]
// }
// Test Case 12
// {
//   "array": [3, 10, 2, 1, 2, 3, 7, 1, 1, 1, 3, 2, 3, 2, 1, 1, 1, 1]
// }
// Test Case 13
// {
//   "array": [3, 12, 2, 1, 2, 3, 7, 1, 1, 1, 3, 2, 3, 2, 1, 1, 1, 1]
// }
// Test Case 14
// {
//   "array": [3, 12, 2, 1, 2, 3, 15, 1, 1, 1, 3, 2, 3, 2, 1, 1, 1, 1]
// }
