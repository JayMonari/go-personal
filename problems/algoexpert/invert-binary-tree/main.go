package main

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func (t *BinaryTree) InvertBinaryTree() {
	if t == nil {
		return
	}
	t.Left, t.Right = t.Right, t.Left
	t.Left.InvertBinaryTree()
	t.Right.InvertBinaryTree()
}

func (t *BinaryTree) InvertBinaryTreeIter() {
	q := []*BinaryTree{t}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			continue
		}
		node.Left, node.Right = node.Right, node.Left
		q = append(q, node.Left, node.Right)
	}
}

// Test Case 1
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "4", "left": "8", "right": "9", "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9}
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
//       {"id": "2", "left": "4", "right": null, "value": 2},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4}
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
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5}
//     ],
//     "root": "1"
//   }
// }
// Test Case 7
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": "6", "right": null, "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6}
//     ],
//     "root": "1"
//   }
// }
// Test Case 8
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7}
//     ],
//     "root": "1"
//   }
// }
// Test Case 9
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "4", "left": "8", "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   }
// }
// Test Case 10
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
// Test Case 11
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "4", "left": "8", "right": "9", "value": 4},
//       {"id": "5", "left": "10", "right": null, "value": 5},
//       {"id": "6", "left": null, "right": "11", "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": "12", "value": 8},
//       {"id": "9", "left": "13", "right": "14", "value": 9},
//       {"id": "10", "left": null, "right": null, "value": 10},
//       {"id": "11", "left": "15", "right": "16", "value": 11},
//       {"id": "12", "left": null, "right": null, "value": 12},
//       {"id": "13", "left": null, "right": null, "value": 13},
//       {"id": "14", "left": null, "right": null, "value": 14},
//       {"id": "15", "left": null, "right": "17", "value": 15},
//       {"id": "16", "left": null, "right": null, "value": 16},
//       {"id": "17", "left": null, "right": "18", "value": 17},
//       {"id": "18", "left": null, "right": "19", "value": 18},
//       {"id": "19", "left": "20", "right": null, "value": 19},
//       {"id": "20", "left": "21", "right": null, "value": 20},
//       {"id": "21", "left": null, "right": null, "value": 21}
//     ],
//     "root": "1"
//   }
// }
