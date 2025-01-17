package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	sums := []int{}
	return branchSums(root, 0, sums)
}

func branchSums(root *BinaryTree, sum int, sums []int) []int {
	sum += root.Value
	if root.Left == nil && root.Right == nil {
		sums = append(sums, sum)
	}
	if root.Left != nil {
		sums = branchSums(root.Left, sum, sums)
	}
	if root.Right != nil {
		sums = branchSums(root.Right, sum, sums)
	}
	return sums
}

// Test Case 1
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "4", "left": "8", "right": "9", "value": 4},
//       {"id": "5", "left": "10", "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9},
//       {"id": "10", "left": null, "right": null, "value": 10}
//     ],
//     "root": "1"
//   }
// }
// Test Case 2
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "1"
//   }
// }
// Test Case 3
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": null, "value": 1},
//       {"id": "2", "left": null, "right": null, "value": 2}
//     ],
//     "root": "1"
//   }
// }
// Test Case 4
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": null, "right": null, "value": 2},
//       {"id": "3", "left": null, "right": null, "value": 3}
//     ],
//     "root": "1"
//   }
// }
// Test Case 5
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5}
//     ],
//     "root": "1"
//   }
// }
// Test Case 6
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "4", "left": "8", "right": "9", "value": 4},
//       {"id": "5", "left": "10", "right": "1-2", "value": 5},
//       {"id": "6", "left": "1-3", "right": "1-4", "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9},
//       {"id": "10", "left": null, "right": null, "value": 10},
//       {"id": "1-2", "left": null, "right": null, "value": 1},
//       {"id": "1-3", "left": null, "right": null, "value": 1},
//       {"id": "1-4", "left": null, "right": null, "value": 1}
//     ],
//     "root": "1"
//   }
// }
// Test Case 7
// {
//   "tree": {
//     "nodes": [
//       {"id": "0", "left": "1", "right": null, "value": 0},
//       {"id": "1", "left": "10", "right": null, "value": 1},
//       {"id": "10", "left": "100", "right": null, "value": 10},
//       {"id": "100", "left": null, "right": null, "value": 100}
//     ],
//     "root": "0"
//   }
// }
// Test Case 8
// {
//   "tree": {
//     "nodes": [
//       {"id": "0", "left": null, "right": "1", "value": 0},
//       {"id": "1", "left": null, "right": "10", "value": 1},
//       {"id": "10", "left": null, "right": "100", "value": 10},
//       {"id": "100", "left": null, "right": null, "value": 100}
//     ],
//     "root": "0"
//   }
// }
// Test Case 9
// {
//   "tree": {
//     "nodes": [
//       {"id": "0", "left": "9", "right": "1", "value": 0},
//       {"id": "9", "left": null, "right": null, "value": 9},
//       {"id": "1", "left": "15", "right": "10", "value": 1},
//       {"id": "15", "left": null, "right": null, "value": 15},
//       {"id": "10", "left": "100", "right": "200", "value": 10},
//       {"id": "100", "left": null, "right": null, "value": 100},
//       {"id": "200", "left": null, "right": null, "value": 200}
//     ],
//     "root": "0"
//   }
// }
