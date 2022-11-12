package main

import "math"

type Block map[string]bool

func ApartmentHunting(blocks []Block, reqs []string) (optimalBlockIdx int) {
	minDistances := [][]int{}
	for _, req := range reqs {
		minDistances = append(minDistances, getMinDistances(blocks, req))
	}
	smallestMaxDistance := math.MaxInt32
	for i, dist := range getMaxDistancesAtBlocks(blocks, minDistances) {
		if dist < smallestMaxDistance {
			smallestMaxDistance = dist
			optimalBlockIdx = i
		}
	}
	return optimalBlockIdx
}

func getMinDistances(blocks []Block, req string) []int {
	minDists := make([]int, len(blocks))
	closestReq := math.MaxInt32
	for i := range blocks {
		if v, ok := blocks[i][req]; ok && v {
			closestReq = i
		}
		minDists[i] = distanceBetween(i, closestReq)
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		if v, ok := blocks[i][req]; ok && v {
			closestReq = i
		}
		minDists[i] = min(minDists[i], distanceBetween(i, closestReq))
	}
	return minDists
}

func getMaxDistancesAtBlocks(blocks []Block, minDistancesFromBlocks [][]int) []int {
	atBlocksMax := make([]int, len(blocks))
	for i := range blocks {
		atBlockMin := []int{}
		for _, d := range minDistancesFromBlocks {
			atBlockMin = append(atBlockMin, d[i])
		}
		atBlocksMax[i] = max(atBlockMin)
	}
	return atBlocksMax
}

func distanceBetween(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// Test Case 1
//
// {
//   "blocks": [
//     {
//       "gym": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": true,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": true,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "school": true,
//       "store": true
//     }
//   ],
//   "reqs": ["gym", "school", "store"]
// }
//
// Test Case 2
//
// {
//   "blocks": [
//     {
//       "gym": false,
//       "office": true,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": true,
//       "office": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": true,
//       "office": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": false,
//       "school": true,
//       "store": true
//     }
//   ],
//   "reqs": ["gym", "office", "school", "store"]
// }
//
// Test Case 3
//
// {
//   "blocks": [
//     {
//       "gym": false,
//       "office": true,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": true,
//       "office": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": true,
//       "office": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": false,
//       "school": true,
//       "store": true
//     }
//   ],
//   "reqs": ["gym", "office", "school", "store"]
// }
//
// Test Case 4
//
// {
//   "blocks": [
//     {
//       "foo": true,
//       "gym": false,
//       "kappa": false,
//       "office": true,
//       "school": true,
//       "store": false
//     },
//     {
//       "foo": true,
//       "gym": true,
//       "kappa": false,
//       "office": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "foo": true,
//       "gym": true,
//       "kappa": false,
//       "office": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "foo": true,
//       "gym": false,
//       "kappa": false,
//       "office": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "foo": true,
//       "gym": true,
//       "kappa": false,
//       "office": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "foo": true,
//       "gym": false,
//       "kappa": false,
//       "office": false,
//       "school": true,
//       "store": true
//     }
//   ],
//   "reqs": ["gym", "school", "store"]
// }
//
// Test Case 5
//
// {
//   "blocks": [
//     {
//       "gym": true,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "school": false,
//       "store": true
//     },
//     {
//       "gym": true,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "school": true,
//       "store": false
//     }
//   ],
//   "reqs": ["gym", "school", "store"]
// }
//
// Test Case 6
//
// {
//   "blocks": [
//     {
//       "gym": true,
//       "pool": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "pool": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "pool": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "pool": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "pool": false,
//       "school": false,
//       "store": true
//     },
//     {
//       "gym": true,
//       "pool": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "pool": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "pool": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "pool": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "pool": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "pool": true,
//       "school": false,
//       "store": false
//     }
//   ],
//   "reqs": ["gym", "pool", "school", "store"]
// }
//
// Test Case 7
//
// {
//   "blocks": [
//     {
//       "gym": true,
//       "office": false,
//       "pool": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": false,
//       "pool": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": true,
//       "pool": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": true,
//       "pool": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": false,
//       "pool": false,
//       "school": false,
//       "store": true
//     },
//     {
//       "gym": true,
//       "office": true,
//       "pool": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": false,
//       "pool": true,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": false,
//       "pool": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": false,
//       "pool": false,
//       "school": false,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": false,
//       "pool": false,
//       "school": true,
//       "store": false
//     },
//     {
//       "gym": false,
//       "office": false,
//       "pool": true,
//       "school": false,
//       "store": false
//     }
//   ],
//   "reqs": ["gym", "pool", "school", "store"]
// }
