package main

type pair []int

func FourNumberSum(nums []int, target int) (quadruplets [][]int) {
	pairSums := make(map[int][]pair)
	for i := 1; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			diff := target - sum
			if pairs, found := pairSums[diff]; found {
				for _, p := range pairs {
					quadruplets = append(quadruplets, append(p, nums[i], nums[j]))
				}
			}
		}
		for k := 0; k < i; k++ {
			sum := nums[i] + nums[k]
			pairSums[sum] = append(pairSums[sum], pair{nums[k], nums[i]})
		}
	}
	return quadruplets
}

// Test Case 1
// {
//   "array": [7, 6, 4, -1, 1, 2],
//   "targetSum": 16
// }
// Test Case 2
// {
//   "array": [1, 2, 3, 4, 5, 6, 7],
//   "targetSum": 10
// }
// Test Case 3
// {
//   "array": [5, -5, -2, 2, 3, -3],
//   "targetSum": 0
// }
// Test Case 4
// {
//   "array": [-2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9],
//   "targetSum": 4
// }
// Test Case 5
// {
//   "array": [-1, 22, 18, 4, 7, 11, 2, -5, -3],
//   "targetSum": 30
// }
// Test Case 6
// {
//   "array": [-10, -3, -5, 2, 15, -7, 28, -6, 12, 8, 11, 5],
//   "targetSum": 20
// }
// Test Case 7
// {
//   "array": [1, 2, 3, 4, 5],
//   "targetSum": 100
// }
// Test Case 8
// {
//   "array": [1, 2, 3, 4, 5, -5, 6, -6],
//   "targetSum": 5
// }
