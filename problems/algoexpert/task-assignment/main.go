package main

import (
	"sort"
)

func TaskAssignment(k int, tasks []int) [][]int {
	durationIdxs := map[int][]int{}
	for i, t := range tasks {
		durationIdxs[t] = append(durationIdxs[t], i)
	}
	sort.Ints(tasks)

	pairedTasks := [][]int{}
	var shortTask, longTask int
	for i := 0; i < k; i++ {
		taskIdxs := durationIdxs[tasks[i]]
		shortTask, durationIdxs[tasks[i]] = taskIdxs[0], taskIdxs[1:]

		taskIdxs = durationIdxs[tasks[len(tasks)-1-i]]
		longTask, durationIdxs[tasks[len(tasks)-1-i]] = taskIdxs[0], taskIdxs[1:]

		pairedTasks = append(pairedTasks, []int{shortTask, longTask})
	}
	return pairedTasks
}

// Test Case 1
// {
//   "k": 3,
//   "tasks": [1, 3, 5, 3, 1, 4]
// }
// Test Case 2
// {
//   "k": 4,
//   "tasks": [1, 2, 3, 4, 5, 6, 7, 8]
// }
// Test Case 3
// {
//   "k": 5,
//   "tasks": [1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
// }
// Test Case 4
// {
//   "k": 1,
//   "tasks": [3, 5]
// }
// Test Case 5
// {
//   "k": 7,
//   "tasks": [2, 1, 3, 4, 5, 13, 12, 9, 11, 10, 6, 7, 14, 8]
// }
// Test Case 6
// {
//   "k": 5,
//   "tasks": [3, 7, 5, 4, 4, 3, 6, 8, 3, 3]
// }
// Test Case 7
// {
//   "k": 10,
//   "tasks": [5, 6, 2, 3, 15, 15, 16, 19, 2, 10, 10, 3, 3, 32, 12, 1, 23, 32, 9, 2]
// }
// Test Case 8
// {
//   "k": 4,
//   "tasks": [1, 2, 2, 1, 3, 4, 4, 4]
// }
// Test Case 9
// {
//   "k": 3,
//   "tasks": [87, 65, 43, 32, 31, 320]
// }
// Test Case 10
// {
//   "k": 2,
//   "tasks": [3, 4, 5, 3]
// }
// Test Case 11
// {
//   "k": 3,
//   "tasks": [5, 2, 1, 6, 4, 4]
// }
// Test Case 12
// {
//   "k": 2,
//   "tasks": [1, 8, 9, 10]
// }
