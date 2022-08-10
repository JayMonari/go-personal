package main

func GetPermutations(nums []int) [][]int {
	perms := [][]int{}
	getPermutations(0, nums, &perms)
	return perms
}

func getPermutations(i int, nums []int, perms *[][]int) {
	if i == len(nums)-1 {
		p := make([]int, len(nums))
		copy(p, nums)
		*perms = append(*perms, p)
		return
	}
	for j := i; j < len(nums); j++ {
		nums[i], nums[j] = nums[j], nums[i]
		getPermutations(i+1, nums, perms)
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// Test Case 1
// {
//   "array": [1, 2, 3]
// }
// Test Case 2
// {
//   "array": []
// }
// Test Case 3
// {
//   "array": [1]
// }
// Test Case 4
// {
//   "array": [1, 2]
// }
// Test Case 5
// {
//   "array": [1, 2, 3, 4]
// }
// Test Case 6
// {
//   "array": [1, 2, 3, 4, 5]
// }
