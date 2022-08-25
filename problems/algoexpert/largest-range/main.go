package main

type Range struct{ left, right int }

func LargestRange(nums []int) []int {
	var r Range
	maxLen := 0
	visit := make(map[int]struct{})
	for _, n := range nums {
		visit[n] = struct{}{}
	}
	for _, n := range nums {
		if _, ok := visit[n]; !ok {
			continue
		}
		delete(visit, n)

		currlen, left, right := 1, n-1, n+1
		for _, ok := visit[left]; ok; _, ok = visit[left] {
			delete(visit, left)
			currlen++
			left--
		}
		for _, ok := visit[right]; ok; _, ok = visit[right] {
			delete(visit, right)
			currlen++
			right++
		}
		if currlen > maxLen {
			maxLen = currlen
			r = Range{left: left + 1, right: right - 1}
		}
	}
	return []int{r.left, r.right}
}

// Test Case 1
// {
//   "array": [1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6]
// }
// Test Case 2
// {
//   "array": [1]
// }
// Test Case 3
// {
//   "array": [1, 2]
// }
// Test Case 4
// {
//   "array": [4, 2, 1, 3]
// }
// Test Case 5
// {
//   "array": [4, 2, 1, 3, 6]
// }
// Test Case 6
// {
//   "array": [8, 4, 2, 10, 3, 6, 7, 9, 1]
// }
// Test Case 7
// {
//   "array": [19, -1, 18, 17, 2, 10, 3, 12, 5, 16, 4, 11, 8, 7, 6, 15, 12, 12, 2, 1, 6, 13, 14]
// }
// Test Case 8
// {
//   "array": [0, 9, 19, -1, 18, 17, 2, 10, 3, 12, 5, 16, 4, 11, 8, 7, 6, 15, 12, 12, 2, 1, 6, 13, 14]
// }
// Test Case 9
// {
//   "array": [0, -5, 9, 19, -1, 18, 17, 2, -4, -3, 10, 3, 12, 5, 16, 4, 11, 7, -6, -7, 6, 15, 12, 12, 2, 1, 6, 13, 14, -2]
// }
// Test Case 10
// {
//   "array": [-7, -7, -7, -7, 8, -8, 0, 9, 19, -1, -3, 18, 17, 2, 10, 3, 12, 5, 16, 4, 11, -6, 8, 7, 6, 15, 12, 12, -5, 2, 1, 6, 13, 14, -4, -2]
// }
// Test Case 11
// {
//   "array": [1, 1, 1, 3, 4]
// }
// Test Case 12
// {
//   "array": [-1, 0, 1]
// }
// Test Case 13
// {
//   "array": [10, 0, 1]
// }
