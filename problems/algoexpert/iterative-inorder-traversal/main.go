package main

type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

func (t *BinaryTree) IterativeInOrderTraversal(callback func(int)) {
	var prev, next *BinaryTree
	for curr := t; curr != nil; prev, curr = curr, next {
		switch {
		case prev == nil || prev == curr.Parent:
			if curr.Left != nil {
				next = curr.Left
				continue
			}

			callback(curr.Value)
			if curr.Right != nil {
				next = curr.Right
				continue
			}
			next = curr.Parent
		case prev == curr.Left:
			callback(curr.Value)
			if curr.Right != nil {
				next = curr.Right
				continue
			}
			next = curr.Parent
		default:
			next = curr.Parent
		}
	}
}

// Test Case 1
//
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": null, "value": 2},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "4", "left": null, "right": "9", "value": 4},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "9", "left": null, "right": null, "value": 9}
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
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "2", "left": "4", "right": null, "value": 2},
//       {"id": "4", "left": null, "right": null, "value": 4}
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
//
// Test Case 5
//
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "5", "left": "10", "right": null, "value": 5},
//       {"id": "10", "left": null, "right": null, "value": 10},
//       {"id": "4", "left": "8", "right": "9", "value": 4},
//       {"id": "9", "left": null, "right": null, "value": 9},
//       {"id": "8", "left": null, "right": null, "value": 8}
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
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "7", "left": null, "right": null, "value": 7},
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
//
// Test Case 7
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
//       {"id": "9", "left": "18", "right": null, "value": 9},
//       {"id": "18", "left": null, "right": null, "value": 18},
//       {"id": "8", "left": "16", "right": "17", "value": 8},
//       {"id": "17", "left": null, "right": null, "value": 17},
//       {"id": "16", "left": null, "right": null, "value": 16}
//     ],
//     "root": "1"
//   }
// }
