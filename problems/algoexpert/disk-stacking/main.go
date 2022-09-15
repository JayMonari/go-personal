package main

import "sort"

type (
	Disk     []int
	ByHeight []Disk
)

func (s ByHeight) Len() int           { return len(s) }
func (s ByHeight) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByHeight) Less(i, j int) bool { return s[i][2] < s[j][2] }

func DiskStacking(input [][]int) [][]int {
	disks := make([]Disk, len(input))
	for i, disk := range input {
		disks[i] = disk
	}
	sort.Sort(ByHeight(disks))

	heights := make([]int, len(disks))
	sequences := make([]int, len(disks))
	for i := range disks {
		heights[i] = disks[i][2]
		sequences[i] = -1
	}

	for i := 1; i < len(disks); i++ {
		c := disks[i]
		for j := 0; j < i; j++ {
			o := disks[j]
			// If other disk has bigger dimensions or doesn't increase height of the
			// overall stack we skip it.
			if (o[0] >= c[0] || o[1] >= c[1] || o[2] >= c[2]) ||
				(heights[i] > c[2]+heights[j]) {
				continue
			}
			heights[i] = c[2] + heights[j]
			sequences[i] = j
		}
	}

	maxIdx := 0
	for i, h := range heights {
		if h >= heights[maxIdx] {
			maxIdx = i
		}
	}

	return buildSequence(disks, sequences, maxIdx)
}

func buildSequence(disks []Disk, sequences []int, idx int) [][]int {
	out := [][]int{}
	for idx != -1 {
		out = append(out, disks[idx])
		idx = sequences[idx]
	}
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}

// Test Case 1
// {
//   "disks": [
//     [2, 1, 2],
//     [3, 2, 3],
//     [2, 2, 8],
//     [2, 3, 4],
//     [1, 3, 1],
//     [4, 4, 5]
//   ]
// }
// Test Case 2
// {
//   "disks": [
//     [2, 1, 2]
//   ]
// }
// Test Case 3
// {
//   "disks": [
//     [2, 1, 2],
//     [3, 2, 3]
//   ]
// }
// Test Case 4
// {
//   "disks": [
//     [2, 1, 2],
//     [3, 2, 3],
//     [2, 2, 8]
//   ]
// }
// Test Case 5
// {
//   "disks": [
//     [2, 1, 2],
//     [3, 2, 3],
//     [2, 3, 4]
//   ]
// }
// Test Case 6
// {
//   "disks": [
//     [2, 1, 2],
//     [3, 2, 3],
//     [2, 2, 8],
//     [2, 3, 4],
//     [2, 2, 1],
//     [4, 4, 5]
//   ]
// }
// Test Case 7
// {
//   "disks": [
//     [2, 1, 2],
//     [3, 2, 5],
//     [2, 2, 8],
//     [2, 3, 4],
//     [2, 2, 1],
//     [4, 4, 5]
//   ]
// }
// Test Case 8
// {
//   "disks": [
//     [2, 1, 2],
//     [3, 2, 3],
//     [2, 2, 8],
//     [2, 3, 4],
//     [1, 2, 1],
//     [4, 4, 5],
//     [1, 1, 4]
//   ]
// }
// Test Case 9
// {
//   "disks": [
//     [3, 3, 4],
//     [2, 1, 2],
//     [3, 2, 3],
//     [2, 2, 8],
//     [2, 3, 4],
//     [5, 5, 6],
//     [1, 2, 1],
//     [4, 4, 5],
//     [1, 1, 4],
//     [2, 2, 3]
//   ]
// }
