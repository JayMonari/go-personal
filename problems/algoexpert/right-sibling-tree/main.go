package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func RightSiblingTree(root *BinaryTree) *BinaryTree {
	mutate(root, nil, false)
	return root
}

func mutate(node, parent *BinaryTree, isLeftChild bool) {
	if node == nil {
		return
	}

	left, right := node.Left, node.Right
	mutate(left, node, true)
	switch {
	case parent == nil:
		node.Right = nil
	case isLeftChild:
		node.Right = parent.Right
	default:
		if parent.Right == nil {
			node.Right = nil
			break
		}
		node.Right = parent.Right.Left
	}
	mutate(right, node, false)
}

// Test Case 1
//
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "7", "left": "12", "right": "13", "value": 7},
//       {"id": "13", "left": null, "right": null, "value": 13},
//       {"id": "12", "left": null, "right": null, "value": 12},
//       {"id": "6", "left": "11", "right": null, "value": 6},
//       {"id": "11", "left": "14", "right": null, "value": 11},
//       {"id": "14", "left": null, "right": null, "value": 14},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "5", "left": null, "right": "10", "value": 5},
//       {"id": "10", "left": null, "right": null, "value": 10},
//       {"id": "4", "left": "8", "right": "9", "value": 4},
//       {"id": "9", "left": null, "right": null, "value": 9},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 2
//
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 3
//
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": null, "value": 1},
//       {"id": "2", "left": null, "right": null, "value": 2}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 4
//
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
//
// Test Case 5
//
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "2", "left": "4", "right": null, "value": 2},
//       {"id": "4", "left": null, "right": null, "value": 4}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 6
//
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "4", "left": null, "right": null, "value": 4}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 7
//
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "3", "left": "6", "right": null, "value": 3},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "4", "left": null, "right": null, "value": 4}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 8
//
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "3", "left": "6", "right": null, "value": 3},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "5", "left": "7", "right": "8", "value": 5},
//       {"id": "8", "left": null, "right": null, "value": 8},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "4", "left": null, "right": null, "value": 4}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 9
//
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "7", "left": "14", "right": "15", "value": 7},
//       {"id": "15", "left": null, "right": null, "value": 15},
//       {"id": "14", "left": null, "right": null, "value": 14},
//       {"id": "6", "left": "12", "right": "13", "value": 6},
//       {"id": "13", "left": null, "right": null, "value": 13},
//       {"id": "12", "left": null, "right": null, "value": 12},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "5", "left": "10", "right": "11", "value": 5},
//       {"id": "11", "left": null, "right": null, "value": 11},
//       {"id": "10", "left": null, "right": null, "value": 10},
//       {"id": "4", "left": "8", "right": "9", "value": 4},
//       {"id": "9", "left": null, "right": null, "value": 9},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   }
// }
