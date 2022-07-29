package main

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

type treeInfo struct {
	diameter, height int
}

func newTreeInfo(t *BinaryTree) treeInfo {
	if t == nil {
		return treeInfo{}
	}
	leftInfo := newTreeInfo(t.Left)
	rightInfo := newTreeInfo(t.Right)

	longestPathWithRoot := leftInfo.height + rightInfo.height
	maxSubDiameter := max(leftInfo.diameter, rightInfo.diameter)
	return treeInfo{
		diameter: max(longestPathWithRoot, maxSubDiameter),
		height:   1 + max(leftInfo.height, rightInfo.height),
	}
}

func BinaryTreeDiameter(t *BinaryTree) int {
	return newTreeInfo(t).diameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test Case 1
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "3", "right": "2", "value": 1},
//       {"id": "3", "left": "7", "right": "4", "value": 3},
//       {"id": "7", "left": "8", "right": null, "value": 7},
//       {"id": "8", "left": "9", "right": null, "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9},
//       {"id": "4", "left": null, "right": "5", "value": 4},
//       {"id": "5", "left": null, "right": "6", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "2", "left": null, "right": null, "value": 2}
//     ],
//     "root": "1"
//   }
// }
// Test Case 2
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
// Test Case 3
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
// Test Case 4
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
// Test Case 5
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
// Test Case 6
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "3", "right": "9", "value": 1},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "9", "left": "14", "right": "10", "value": 9},
//       {"id": "10", "left": null, "right": "11", "value": 10},
//       {"id": "11", "left": null, "right": "12", "value": 11},
//       {"id": "12", "left": null, "right": "17", "value": 12},
//       {"id": "17", "left": null, "right": null, "value": 17},
//       {"id": "14", "left": null, "right": "19", "value": 14},
//       {"id": "19", "left": "25", "right": null, "value": 19},
//       {"id": "25", "left": null, "right": null, "value": 25}
//     ],
//     "root": "1"
//   }
// }
// Test Case 7
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "3", "right": "5", "value": 1},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "5", "left": null, "right": null, "value": 5}
//     ],
//     "root": "1"
//   }
// }
// Test Case 8
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "5", "right": null, "value": 1},
//       {"id": "5", "left": "7", "right": "9", "value": 5},
//       {"id": "9", "left": null, "right": "12", "value": 9},
//       {"id": "12", "left": null, "right": null, "value": 12},
//       {"id": "7", "left": "8", "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   }
// }
// Test Case 9
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "1"
//   }
// }
// Test Case 10
// {
//   "tree": {
//     "nodes": [
//       {"id": "4", "left": "2", "right": null, "value": 4},
//       {"id": "2", "left": null, "right": null, "value": 2}
//     ],
//     "root": "4"
//   }
// }
// Test Case 11
// {
//   "tree": {
//     "nodes": [
//       {"id": "4", "left": "2", "right": null, "value": 4},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "4"
//   }
// }
// Test Case 12
// {
//   "tree": {
//     "nodes": [
//       {"id": "4", "left": "2", "right": null, "value": 4},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "1", "left": null, "right": "3", "value": 1},
//       {"id": "3", "left": "19", "right": null, "value": 3},
//       {"id": "19", "left": null, "right": null, "value": 19}
//     ],
//     "root": "4"
//   }
// }
// Test Case 13
// {
//   "tree": {
//     "nodes": [
//       {"id": "6", "left": null, "right": "1", "value": 6},
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "6"
//   }
// }
// Test Case 14
// {
//   "tree": {
//     "nodes": [
//       {"id": "3", "left": null, "right": "10", "value": 3},
//       {"id": "10", "left": "1", "right": null, "value": 10},
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "3"
//   }
// }
// Test Case 15
// {
//   "tree": {
//     "nodes": [
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "1", "left": "3", "right": null, "value": 1},
//       {"id": "3", "left": null, "right": "5", "value": 3},
//       {"id": "5", "left": null, "right": "10", "value": 5},
//       {"id": "10", "left": null, "right": null, "value": 10}
//     ],
//     "root": "2"
//   }
// }
