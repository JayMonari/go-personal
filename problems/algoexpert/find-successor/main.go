package main

type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

func FindSuccessor(root, node *BinaryTree) *BinaryTree {
	vals := inOrderTraversal(root, node, []*BinaryTree{})
	for i, n := range vals {
		if node.Value != n.Value {
			continue
		}
		if i != len(vals)-1 {
			return vals[i+1]
		}
	}
	return nil
}

func inOrderTraversal(root, node *BinaryTree, vals []*BinaryTree) []*BinaryTree {
	if root == nil {
		return vals
	}
	vals = inOrderTraversal(root.Left, node, vals)
	vals = append(vals, root)
	vals = inOrderTraversal(root.Right, node, vals)
	return vals
}

func FindSuccessorOpt(root, node *BinaryTree) *BinaryTree {
	if node.Right != nil {
		return node.Right.leftmostChild()
	}
	return node.rightmostParent()
}

func (t *BinaryTree) leftmostChild() *BinaryTree {
	for t.Left != nil {
		t = t.Left
	}
	return t
}

func (t *BinaryTree) rightmostParent() *BinaryTree {
	for t.Parent != nil && t.Parent.Right == t {
		t = t.Parent
	}
	return t.Parent
}

// Test Case 1
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "parent": null, "right": "3", "value": 1},
//       {"id": "2", "left": "4", "parent": "1", "right": "5", "value": 2},
//       {"id": "3", "left": null, "parent": "1", "right": null, "value": 3},
//       {"id": "4", "left": "6", "parent": "2", "right": null, "value": 4},
//       {"id": "5", "left": null, "parent": "2", "right": null, "value": 5},
//       {"id": "6", "left": null, "parent": "4", "right": null, "value": 6}
//     ],
//     "root": "1"
//   },
//   "node": "5"
// }
// Test Case 2
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "parent": null, "right": "3", "value": 1},
//       {"id": "2", "left": "4", "parent": "1", "right": "5", "value": 2},
//       {"id": "3", "left": null, "parent": "1", "right": null, "value": 3},
//       {"id": "4", "left": null, "parent": "2", "right": null, "value": 4},
//       {"id": "5", "left": "6", "parent": "2", "right": "7", "value": 5},
//       {"id": "6", "left": null, "parent": "5", "right": null, "value": 6},
//       {"id": "7", "left": "8", "parent": "5", "right": null, "value": 7},
//       {"id": "8", "left": null, "parent": "7", "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "node": "5"
// }
// Test Case 3
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "parent": null, "right": "3", "value": 1},
//       {"id": "2", "left": "4", "parent": "1", "right": "5", "value": 2},
//       {"id": "3", "left": "6", "parent": "1", "right": "7", "value": 3},
//       {"id": "4", "left": null, "parent": "2", "right": null, "value": 4},
//       {"id": "5", "left": null, "parent": "2", "right": null, "value": 5},
//       {"id": "6", "left": null, "parent": "3", "right": null, "value": 6},
//       {"id": "7", "left": null, "parent": "3", "right": null, "value": 7}
//     ],
//     "root": "1"
//   },
//   "node": "6"
// }
// Test Case 4
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "parent": null, "right": "2", "value": 1},
//       {"id": "2", "left": null, "parent": "1", "right": null, "value": 2}
//     ],
//     "root": "1"
//   },
//   "node": "2"
// }
// Test Case 5
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "parent": null, "right": "3", "value": 1},
//       {"id": "2", "left": null, "parent": "1", "right": null, "value": 2},
//       {"id": "3", "left": "4", "parent": "1", "right": null, "value": 3},
//       {"id": "4", "left": "5", "parent": "3", "right": null, "value": 4},
//       {"id": "5", "left": "6", "parent": "4", "right": null, "value": 5},
//       {"id": "6", "left": "7", "parent": "5", "right": null, "value": 6},
//       {"id": "7", "left": null, "parent": "6", "right": null, "value": 7}
//     ],
//     "root": "1"
//   },
//   "node": "1"
// }
// Test Case 6
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "parent": null, "right": null, "value": 1},
//       {"id": "2", "left": "3", "parent": "1", "right": null, "value": 2},
//       {"id": "3", "left": "4", "parent": "2", "right": null, "value": 3},
//       {"id": "4", "left": "5", "parent": "3", "right": null, "value": 4},
//       {"id": "5", "left": null, "parent": "4", "right": null, "value": 5}
//     ],
//     "root": "1"
//   },
//   "node": "3"
// }
// Test Case 7
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "parent": null, "right": null, "value": 1},
//       {"id": "2", "left": "3", "parent": "1", "right": "6", "value": 2},
//       {"id": "3", "left": "4", "parent": "2", "right": null, "value": 3},
//       {"id": "4", "left": "5", "parent": "3", "right": null, "value": 4},
//       {"id": "5", "left": null, "parent": "4", "right": null, "value": 5},
//       {"id": "6", "left": "7", "parent": "2", "right": "8", "value": 6},
//       {"id": "7", "left": null, "parent": "6", "right": null, "value": 7},
//       {"id": "8", "left": null, "parent": "6", "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "node": "2"
// }
// Test Case 8
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "parent": null, "right": null, "value": 1}
//     ],
//     "root": "1"
//   },
//   "node": "1"
// }
// Test Case 9
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "parent": null, "right": "2", "value": 1},
//       {"id": "2", "left": null, "parent": "1", "right": null, "value": 2}
//     ],
//     "root": "1"
//   },
//   "node": "1"
// }
// Test Case 10
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "parent": null, "right": "5", "value": 1},
//       {"id": "2", "left": null, "parent": "1", "right": "3", "value": 2},
//       {"id": "3", "left": null, "parent": "2", "right": "4", "value": 3},
//       {"id": "4", "left": null, "parent": "3", "right": null, "value": 4},
//       {"id": "5", "left": null, "parent": "1", "right": "6", "value": 5},
//       {"id": "6", "left": "7", "parent": "5", "right": "8", "value": 6},
//       {"id": "7", "left": null, "parent": "6", "right": null, "value": 7},
//       {"id": "8", "left": null, "parent": "6", "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "node": "1"
// }
// Test Case 11
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "parent": null, "right": "5", "value": 1},
//       {"id": "2", "left": null, "parent": "1", "right": "3", "value": 2},
//       {"id": "3", "left": null, "parent": "2", "right": "4", "value": 3},
//       {"id": "4", "left": null, "parent": "3", "right": null, "value": 4},
//       {"id": "5", "left": "9", "parent": "1", "right": "6", "value": 5},
//       {"id": "6", "left": "7", "parent": "5", "right": "8", "value": 6},
//       {"id": "7", "left": null, "parent": "6", "right": null, "value": 7},
//       {"id": "8", "left": null, "parent": "6", "right": null, "value": 8},
//       {"id": "9", "left": "10", "parent": "5", "right": null, "value": 9},
//       {"id": "10", "left": "11", "parent": "9", "right": null, "value": 10},
//       {"id": "11", "left": null, "parent": "10", "right": null, "value": 11}
//     ],
//     "root": "1"
//   },
//   "node": "1"
// }
// Test Case 12
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "parent": null, "right": "3", "value": 1},
//       {"id": "2", "left": "4", "parent": "1", "right": "5", "value": 2},
//       {"id": "3", "left": null, "parent": "1", "right": "7", "value": 3},
//       {"id": "4", "left": "6", "parent": "2", "right": null, "value": 4},
//       {"id": "5", "left": null, "parent": "2", "right": null, "value": 5},
//       {"id": "6", "left": null, "parent": "4", "right": null, "value": 6},
//       {"id": "7", "left": null, "parent": "3", "right": "8", "value": 7},
//       {"id": "8", "left": null, "parent": "7", "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "node": "1"
// }
