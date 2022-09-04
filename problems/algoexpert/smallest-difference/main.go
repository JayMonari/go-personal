package main

import "sort"

func SmallestDifference(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	smallest := make([]int, 2)
	smallest[0], smallest[1] = nums1[0], nums2[0]
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		n1, n2 := nums1[i], nums2[j]
		switch {
		case n1 < n2:
			i++
		case n1 > n2:
			j++
		case n1 == n2:
			return []int{n1, n2}
		}
		if absDiff(n1, n2) < absDiff(smallest[0], smallest[1]) {
			smallest[0], smallest[1] = n1, n2
		}
	}
	return smallest
}

func absDiff(a, b int) int {
	if a < b {
		return absDiff(b, a)
	}
	return a - b
}

// Test Case 1
// {
//   "arrayOne": [-1, 5, 10, 20, 28, 3],
//   "arrayTwo": [26, 134, 135, 15, 17]
// }
// Test Case 2
// {
//   "arrayOne": [-1, 5, 10, 20, 3],
//   "arrayTwo": [26, 134, 135, 15, 17]
// }
// Test Case 3
// {
//   "arrayOne": [10, 0, 20, 25],
//   "arrayTwo": [1005, 1006, 1014, 1032, 1031]
// }
// Test Case 4
// {
//   "arrayOne": [10, 0, 20, 25, 2200],
//   "arrayTwo": [1005, 1006, 1014, 1032, 1031]
// }
// Test Case 5
// {
//   "arrayOne": [10, 0, 20, 25, 2000],
//   "arrayTwo": [1005, 1006, 1014, 1032, 1031]
// }
// Test Case 6
// {
//   "arrayOne": [240, 124, 86, 111, 2, 84, 954, 27, 89],
//   "arrayTwo": [1, 3, 954, 19, 8]
// }
// Test Case 7
// {
//   "arrayOne": [0, 20],
//   "arrayTwo": [21, -2]
// }
// Test Case 8
// {
//   "arrayOne": [10, 1000],
//   "arrayTwo": [-1441, -124, -25, 1014, 1500, 660, 410, 245, 530]
// }
// Test Case 9
// {
//   "arrayOne": [10, 1000, 9124, 2142, 59, 24, 596, 591, 124, -123],
//   "arrayTwo": [-1441, -124, -25, 1014, 1500, 660, 410, 245, 530]
// }
// Test Case 10
// {
//   "arrayOne": [10, 1000, 9124, 2142, 59, 24, 596, 591, 124, -123, 530],
//   "arrayTwo": [-1441, -124, -25, 1014, 1500, 660, 410, 245, 530]
// }
