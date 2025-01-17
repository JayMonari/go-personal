package main

import "math"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func MaxPathSum(t *BinaryTree) int {
	_, maxSum := findMaxSum(t)
	return maxSum
}

func findMaxSum(t *BinaryTree) (int, int) {
	if t == nil {
		return 0, math.MinInt32
	}
	leftSumBranch, leftPathMax := findMaxSum(t.Left)
	rightSumBranch, rightPathMax := findMaxSum(t.Right)

	v := t.Value
	maxSumBranch := max(max(leftSumBranch, rightSumBranch)+v, v)
	return maxSumBranch, max(
		leftPathMax,
		rightPathMax,
		max(leftSumBranch+v+rightSumBranch, maxSumBranch))
}

func max(first int, vals ...int) int {
	for _, v := range vals {
		if v > first {
			first = v
		}
	}
	return first
}

// Test Case 1
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "4", "left": null, "right": null, "value": 4}
//     ],
//     "root": "1"
//   }
// }
// Test Case 2
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "2", "left": null, "right": null, "value": 2}
//     ],
//     "root": "1"
//   }
// }
// Test Case 3
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "-1", "value": 1},
//       {"id": "-1", "left": null, "right": null, "value": -1},
//       {"id": "2", "left": null, "right": null, "value": 2}
//     ],
//     "root": "1"
//   }
// }
// Test Case 4
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "-5", "right": "3", "value": 1},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "-5", "left": "6", "right": null, "value": -5},
//       {"id": "6", "left": null, "right": null, "value": 6}
//     ],
//     "root": "1"
//   }
// }
// Test Case 5
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "-10", "right": "-5", "value": 1},
//       {"id": "-5", "left": "-20", "right": "-21", "value": -5},
//       {"id": "-21", "left": "100-2", "right": "1-3", "value": -21},
//       {"id": "1-3", "left": null, "right": null, "value": 1},
//       {"id": "100-2", "left": null, "right": null, "value": 100},
//       {"id": "-20", "left": "100", "right": "2", "value": -20},
//       {"id": "2", "left": null, "right": null, "value": 2},
//       {"id": "100", "left": null, "right": null, "value": 100},
//       {"id": "-10", "left": "30", "right": "45", "value": -10},
//       {"id": "45", "left": "3", "right": "-3", "value": 45},
//       {"id": "-3", "left": null, "right": null, "value": -3},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "30", "left": "5", "right": "1-2", "value": 30},
//       {"id": "1-2", "left": null, "right": null, "value": 1},
//       {"id": "5", "left": null, "right": null, "value": 5}
//     ],
//     "root": "1"
//   }
// }
// Test Case 6
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "-10", "right": "-5", "value": 1},
//       {"id": "-5", "left": "-20", "right": "-21", "value": -5},
//       {"id": "-21", "left": "100-3", "right": "1-3", "value": -21},
//       {"id": "1-3", "left": null, "right": null, "value": 1},
//       {"id": "100-3", "left": null, "right": null, "value": 100},
//       {"id": "-20", "left": "100-2", "right": "2", "value": -20},
//       {"id": "2", "left": null, "right": null, "value": 2},
//       {"id": "100-2", "left": null, "right": null, "value": 100},
//       {"id": "-10", "left": "30", "right": "45", "value": -10},
//       {"id": "45", "left": "3", "right": "-3", "value": 45},
//       {"id": "-3", "left": null, "right": null, "value": -3},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "30", "left": "5", "right": "1-2", "value": 30},
//       {"id": "1-2", "left": null, "right": null, "value": 1},
//       {"id": "5", "left": "100", "right": null, "value": 5},
//       {"id": "100", "left": null, "right": null, "value": 100}
//     ],
//     "root": "1"
//   }
// }
// Test Case 7
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "-10", "right": "-5", "value": 1},
//       {"id": "-5", "left": "-20", "right": "-21", "value": -5},
//       {"id": "-21", "left": "100-3", "right": "1-3", "value": -21},
//       {"id": "1-3", "left": null, "right": null, "value": 1},
//       {"id": "100-3", "left": null, "right": null, "value": 100},
//       {"id": "-20", "left": "100-2", "right": "2", "value": -20},
//       {"id": "2", "left": null, "right": null, "value": 2},
//       {"id": "100-2", "left": null, "right": null, "value": 100},
//       {"id": "-10", "left": "30", "right": "75", "value": -10},
//       {"id": "75", "left": "3", "right": "-3", "value": 75},
//       {"id": "-3", "left": null, "right": null, "value": -3},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "30", "left": "5", "right": "1-2", "value": 30},
//       {"id": "1-2", "left": null, "right": null, "value": 1},
//       {"id": "5", "left": "100", "right": null, "value": 5},
//       {"id": "100", "left": null, "right": null, "value": 100}
//     ],
//     "root": "1"
//   }
// }
// Test Case 8
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "-150", "right": "-5", "value": 1},
//       {"id": "-5", "left": "-20", "right": "-21", "value": -5},
//       {"id": "-21", "left": "100-4", "right": "1-3", "value": -21},
//       {"id": "1-3", "left": null, "right": null, "value": 1},
//       {"id": "100-4", "left": null, "right": null, "value": 100},
//       {"id": "-20", "left": "100-3", "right": "2", "value": -20},
//       {"id": "2", "left": null, "right": null, "value": 2},
//       {"id": "100-3", "left": null, "right": null, "value": 100},
//       {"id": "-150", "left": "30", "right": "75", "value": -150},
//       {"id": "75", "left": "3", "right": "-3", "value": 75},
//       {"id": "-3", "left": null, "right": null, "value": -3},
//       {"id": "3", "left": "150", "right": "-8", "value": 3},
//       {"id": "-8", "left": null, "right": null, "value": -8},
//       {"id": "150", "left": null, "right": null, "value": 150},
//       {"id": "30", "left": "5", "right": "1-2", "value": 30},
//       {"id": "1-2", "left": "5-2", "right": "10", "value": 1},
//       {"id": "10", "left": null, "right": null, "value": 10},
//       {"id": "5-2", "left": null, "right": null, "value": 5},
//       {"id": "5", "left": "100", "right": "100-2", "value": 5},
//       {"id": "100-2", "left": null, "right": null, "value": 100},
//       {"id": "100", "left": null, "right": null, "value": 100}
//     ],
//     "root": "1"
//   }
// }
// Test Case 9
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "-150", "right": "-5", "value": 1},
//       {"id": "-5", "left": "-20", "right": "-21", "value": -5},
//       {"id": "-21", "left": "100-4", "right": "1-3", "value": -21},
//       {"id": "1-3", "left": null, "right": null, "value": 1},
//       {"id": "100-4", "left": null, "right": null, "value": 100},
//       {"id": "-20", "left": "100-3", "right": "2", "value": -20},
//       {"id": "2", "left": null, "right": null, "value": 2},
//       {"id": "100-3", "left": null, "right": null, "value": 100},
//       {"id": "-150", "left": "30", "right": "75", "value": -150},
//       {"id": "75", "left": "3", "right": "-3", "value": 75},
//       {"id": "-3", "left": null, "right": null, "value": -3},
//       {"id": "3", "left": "150", "right": "151", "value": 3},
//       {"id": "151", "left": null, "right": null, "value": 151},
//       {"id": "150", "left": null, "right": null, "value": 150},
//       {"id": "30", "left": "5", "right": "1-2", "value": 30},
//       {"id": "1-2", "left": "5-2", "right": "10", "value": 1},
//       {"id": "10", "left": null, "right": null, "value": 10},
//       {"id": "5-2", "left": null, "right": null, "value": 5},
//       {"id": "5", "left": "100", "right": "100-2", "value": 5},
//       {"id": "100-2", "left": null, "right": null, "value": 100},
//       {"id": "100", "left": null, "right": null, "value": 100}
//     ],
//     "root": "1"
//   }
// }
// Test Case 10
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "-5", "right": "-3-2", "value": 1},
//       {"id": "-3-2", "left": "2-2", "right": "1-7", "value": -3},
//       {"id": "1-7", "left": "1-8", "right": "1-10", "value": 1},
//       {"id": "1-10", "left": "-5-2", "right": "0-5", "value": 1},
//       {"id": "0-5", "left": null, "right": null, "value": 0},
//       {"id": "-5-2", "left": null, "right": null, "value": -5},
//       {"id": "1-8", "left": "0-4", "right": "1-9", "value": 1},
//       {"id": "1-9", "left": null, "right": null, "value": 1},
//       {"id": "0-4", "left": null, "right": null, "value": 0},
//       {"id": "2-2", "left": "0-3", "right": "5", "value": 2},
//       {"id": "5", "left": "2-3", "right": "1-6", "value": 5},
//       {"id": "1-6", "left": null, "right": null, "value": 1},
//       {"id": "2-3", "left": null, "right": null, "value": 2},
//       {"id": "0-3", "left": "-9", "right": "-91", "value": 0},
//       {"id": "-91", "left": null, "right": null, "value": -91},
//       {"id": "-9", "left": null, "right": null, "value": -9},
//       {"id": "-5", "left": "0", "right": "2", "value": -5},
//       {"id": "2", "left": "1-4", "right": "1-5", "value": 2},
//       {"id": "1-5", "left": "-1-3", "right": "-100", "value": 1},
//       {"id": "-100", "left": null, "right": null, "value": -100},
//       {"id": "-1-3", "left": null, "right": null, "value": -1},
//       {"id": "1-4", "left": "-1-2", "right": "-6", "value": 1},
//       {"id": "-6", "left": null, "right": null, "value": -6},
//       {"id": "-1-2", "left": null, "right": null, "value": -1},
//       {"id": "0", "left": "-3", "right": "3", "value": 0},
//       {"id": "3", "left": "1-3", "right": "-1", "value": 3},
//       {"id": "-1", "left": null, "right": null, "value": -1},
//       {"id": "1-3", "left": null, "right": null, "value": 1},
//       {"id": "-3", "left": "0-2", "right": "1-2", "value": -3},
//       {"id": "1-2", "left": null, "right": null, "value": 1},
//       {"id": "0-2", "left": null, "right": null, "value": 0}
//     ],
//     "root": "1"
//   }
// }
// Test Case 11
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "-5", "right": "-3-2", "value": 1},
//       {"id": "-3-2", "left": "2-2", "right": "1-6", "value": -3},
//       {"id": "1-6", "left": "1-7", "right": "1-9", "value": 1},
//       {"id": "1-9", "left": "-5-2", "right": "0-5", "value": 1},
//       {"id": "0-5", "left": null, "right": null, "value": 0},
//       {"id": "-5-2", "left": null, "right": null, "value": -5},
//       {"id": "1-7", "left": "0-4", "right": "1-8", "value": 1},
//       {"id": "1-8", "left": null, "right": null, "value": 1},
//       {"id": "0-4", "left": null, "right": null, "value": 0},
//       {"id": "2-2", "left": "0-3", "right": "5", "value": 2},
//       {"id": "5", "left": "2-3", "right": "1-5", "value": 5},
//       {"id": "1-5", "left": null, "right": null, "value": 1},
//       {"id": "2-3", "left": null, "right": null, "value": 2},
//       {"id": "0-3", "left": "-9", "right": "-91", "value": 0},
//       {"id": "-91", "left": null, "right": null, "value": -91},
//       {"id": "-9", "left": null, "right": null, "value": -9},
//       {"id": "-5", "left": "0", "right": "2", "value": -5},
//       {"id": "2", "left": "1-3", "right": "1-4", "value": 2},
//       {"id": "1-4", "left": "-1-3", "right": "-100", "value": 1},
//       {"id": "-100", "left": null, "right": null, "value": -100},
//       {"id": "-1-3", "left": null, "right": null, "value": -1},
//       {"id": "1-3", "left": "-1-2", "right": "-6", "value": 1},
//       {"id": "-6", "left": null, "right": null, "value": -6},
//       {"id": "-1-2", "left": null, "right": null, "value": -1},
//       {"id": "0", "left": "-3", "right": "-4", "value": 0},
//       {"id": "-4", "left": "10", "right": "-1", "value": -4},
//       {"id": "-1", "left": null, "right": null, "value": -1},
//       {"id": "10", "left": null, "right": null, "value": 10},
//       {"id": "-3", "left": "0-2", "right": "1-2", "value": -3},
//       {"id": "1-2", "left": null, "right": null, "value": 1},
//       {"id": "0-2", "left": null, "right": null, "value": 0}
//     ],
//     "root": "1"
//   }
// }
// Test Case 12
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "-5", "right": "-3-2", "value": 1},
//       {"id": "-3-2", "left": "2-4", "right": "1-7", "value": -3},
//       {"id": "1-7", "left": "1-8", "right": "1-10", "value": 1},
//       {"id": "1-10", "left": "-5-3", "right": "0-5", "value": 1},
//       {"id": "0-5", "left": null, "right": null, "value": 0},
//       {"id": "-5-3", "left": null, "right": null, "value": -5},
//       {"id": "1-8", "left": "0-4", "right": "1-9", "value": 1},
//       {"id": "1-9", "left": null, "right": null, "value": 1},
//       {"id": "0-4", "left": null, "right": null, "value": 0},
//       {"id": "2-4", "left": "0-3", "right": "5", "value": 2},
//       {"id": "5", "left": "2-5", "right": "1-6", "value": 5},
//       {"id": "1-6", "left": null, "right": null, "value": 1},
//       {"id": "2-5", "left": null, "right": null, "value": 2},
//       {"id": "0-3", "left": "-9", "right": "-91", "value": 0},
//       {"id": "-91", "left": null, "right": null, "value": -91},
//       {"id": "-9", "left": null, "right": null, "value": -9},
//       {"id": "-5", "left": "0", "right": "2-3", "value": -5},
//       {"id": "2-3", "left": "1-4", "right": "1-5", "value": 2},
//       {"id": "1-5", "left": "-1-3", "right": "-100", "value": 1},
//       {"id": "-100", "left": null, "right": null, "value": -100},
//       {"id": "-1-3", "left": null, "right": null, "value": -1},
//       {"id": "1-4", "left": "-1-2", "right": "-6", "value": 1},
//       {"id": "-6", "left": null, "right": null, "value": -6},
//       {"id": "-1-2", "left": null, "right": null, "value": -1},
//       {"id": "0", "left": "-3", "right": "-4", "value": 0},
//       {"id": "-4", "left": "3-2", "right": "-1", "value": -4},
//       {"id": "-1", "left": null, "right": null, "value": -1},
//       {"id": "3-2", "left": "7", "right": "-5-2", "value": 3},
//       {"id": "-5-2", "left": null, "right": null, "value": -5},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "-3", "left": "0-2", "right": "1-3", "value": -3},
//       {"id": "1-3", "left": "2", "right": "2-2", "value": 1},
//       {"id": "2-2", "left": null, "right": null, "value": 2},
//       {"id": "2", "left": null, "right": null, "value": 2},
//       {"id": "0-2", "left": "3", "right": "1-2", "value": 0},
//       {"id": "1-2", "left": null, "right": null, "value": 1},
//       {"id": "3", "left": null, "right": null, "value": 3}
//     ],
//     "root": "1"
//   }
// }
// Test Case 13
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "-5", "right": "-3-2", "value": 1},
//       {"id": "-3-2", "left": "2-2", "right": "1-7", "value": -3},
//       {"id": "1-7", "left": "1-8", "right": "1-10", "value": 1},
//       {"id": "1-10", "left": "5-2", "right": "0-5", "value": 1},
//       {"id": "0-5", "left": null, "right": null, "value": 0},
//       {"id": "5-2", "left": null, "right": null, "value": 5},
//       {"id": "1-8", "left": "0-4", "right": "1-9", "value": 1},
//       {"id": "1-9", "left": null, "right": null, "value": 1},
//       {"id": "0-4", "left": null, "right": null, "value": 0},
//       {"id": "2-2", "left": "0-3", "right": "5", "value": 2},
//       {"id": "5", "left": "2-3", "right": "1-6", "value": 5},
//       {"id": "1-6", "left": null, "right": null, "value": 1},
//       {"id": "2-3", "left": null, "right": null, "value": 2},
//       {"id": "0-3", "left": "-9", "right": "-91", "value": 0},
//       {"id": "-91", "left": null, "right": null, "value": -91},
//       {"id": "-9", "left": null, "right": null, "value": -9},
//       {"id": "-5", "left": "0", "right": "2", "value": -5},
//       {"id": "2", "left": "1-4", "right": "1-5", "value": 2},
//       {"id": "1-5", "left": "-1-3", "right": "-100", "value": 1},
//       {"id": "-100", "left": null, "right": null, "value": -100},
//       {"id": "-1-3", "left": null, "right": null, "value": -1},
//       {"id": "1-4", "left": "-1-2", "right": "-6", "value": 1},
//       {"id": "-6", "left": null, "right": null, "value": -6},
//       {"id": "-1-2", "left": null, "right": null, "value": -1},
//       {"id": "0", "left": "-3", "right": "3", "value": 0},
//       {"id": "3", "left": "1-3", "right": "-1", "value": 3},
//       {"id": "-1", "left": null, "right": null, "value": -1},
//       {"id": "1-3", "left": null, "right": null, "value": 1},
//       {"id": "-3", "left": "0-2", "right": "1-2", "value": -3},
//       {"id": "1-2", "left": null, "right": null, "value": 1},
//       {"id": "0-2", "left": null, "right": null, "value": 0}
//     ],
//     "root": "1"
//   }
// }
// Test Case 14
// {
//   "tree": {
//     "nodes": [
//       {"id": "-2", "left": null, "right": null, "value": -2}
//     ],
//     "root": "-2"
//   }
// }
// Test Case 15
// {
//   "tree": {
//     "nodes": [
//       {"id": "-2", "left": "-1", "right": null, "value": -2},
//       {"id": "-1", "left": null, "right": null, "value": -1}
//     ],
//     "root": "-2"
//   }
// }
// Test Case 16
// {
//   "tree": {
//     "nodes": [
//       {"id": "-2", "left": "-1", "right": null, "value": -2},
//       {"id": "-1", "left": "2", "right": "3", "value": -1},
//       {"id": "2", "left": null, "right": null, "value": 2},
//       {"id": "3", "left": null, "right": null, "value": 3}
//     ],
//     "root": "-2"
//   }
// }
