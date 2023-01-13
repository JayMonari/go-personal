package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

type TreePair struct{ First, Second *BinaryTree }

func CompareLeafTraversal(t1 *BinaryTree, t2 *BinaryTree) bool {
	node1 := connectLeafNodes(t1, nil, nil).First
	node2 := connectLeafNodes(t2, nil, nil).First
	for node1 != nil && node2 != nil {
		if node1.Value != node2.Value {
			return false
		}

		node1 = node1.Right
		node2 = node2.Right
	}
	return node1 == nil && node2 == nil
}

func connectLeafNodes(currNode *BinaryTree, head *BinaryTree, prevNode *BinaryTree) TreePair {
	if currNode == nil {
		return TreePair{head, prevNode}
	}

	newHead := head
	newPrevNode := prevNode

	if currNode.Left == nil && currNode.Right == nil {
		newPrevNode = currNode
		switch prevNode {
		case nil:
			newHead = currNode
		default:
			prevNode.Right = currNode
		}
	}

	leftPair := connectLeafNodes(currNode.Left, newHead, newPrevNode)
	leftHead, leftPreviousNode := leftPair.First, leftPair.Second
	return connectLeafNodes(currNode.Right, leftHead, leftPreviousNode)
}

// Test Case 1
//
// {
//   "tree1": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": null, "right": "6", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": "7", "right": "8", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "tree2": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "7", "value": 2},
//       {"id": "3", "left": null, "right": "5", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": "8", "right": "6", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 2
//
// {
//   "tree1": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "1"
//   },
//   "tree2": {
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
//   "tree1": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": null, "value": 1},
//       {"id": "2", "left": "3", "right": null, "value": 2},
//       {"id": "3", "left": null, "right": "4", "value": 3},
//       {"id": "4", "left": "5", "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5}
//     ],
//     "root": "1"
//   },
//   "tree2": {
//     "nodes": [
//       {"id": "1", "left": null, "right": "2", "value": 1},
//       {"id": "2", "left": null, "right": "3", "value": 2},
//       {"id": "3", "left": "4", "right": null, "value": 3},
//       {"id": "4", "left": null, "right": "5", "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 4
//
// {
//   "tree1": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": null, "right": "6", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": "7", "right": "8", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "tree2": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "7", "value": 2},
//       {"id": "3", "left": null, "right": "5", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": "8", "right": "6", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": "9", "right": null, "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 5
//
// {
//   "tree1": {
//     "nodes": [
//       {"id": "2", "left": null, "right": null, "value": 2}
//     ],
//     "root": "2"
//   },
//   "tree2": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 6
//
// {
//   "tree1": {
//     "nodes": [
//       {"id": "2", "left": null, "right": "1", "value": 2},
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "2"
//   },
//   "tree2": {
//     "nodes": [
//       {"id": "2", "left": null, "right": null, "value": 2},
//       {"id": "1", "left": "2", "right": null, "value": 1}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 7
//
// {
//   "tree1": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": null, "value": 2},
//       {"id": "3", "left": null, "right": "5", "value": 3},
//       {"id": "4", "left": "7", "right": null, "value": 4},
//       {"id": "5", "left": "8", "right": "9", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": "6", "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9}
//     ],
//     "root": "1"
//   },
//   "tree2": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "4", "left": "6", "right": null, "value": 4},
//       {"id": "5", "left": "8", "right": "9", "value": 5},
//       {"id": "6", "left": null, "right": "7", "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 8
//
// {
//   "tree1": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": "8", "right": null, "value": 6},
//       {"id": "7", "left": null, "right": "9", "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9}
//     ],
//     "root": "1"
//   },
//   "tree2": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": null, "value": 1},
//       {"id": "2", "left": "4", "right": "3", "value": 2},
//       {"id": "3", "left": "5", "right": "6", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": "7", "right": "9", "value": 6},
//       {"id": "7", "left": "8", "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 9
//
// {
//   "tree1": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "8", "value": 1},
//       {"id": "2", "left": "3", "right": "5", "value": 2},
//       {"id": "3", "left": "4", "right": "6", "value": 3},
//       {"id": "4", "left": null, "right": "7", "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "tree2": {
//     "nodes": [
//       {"id": "1", "left": "8", "right": "2", "value": 1},
//       {"id": "2", "left": "5", "right": "3", "value": 2},
//       {"id": "3", "left": "6", "right": "4", "value": 3},
//       {"id": "4", "left": "7", "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 10
//
// {
//   "tree1": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "8", "value": 1},
//       {"id": "2", "left": "3", "right": "5", "value": 2},
//       {"id": "3", "left": "4", "right": null, "value": 3},
//       {"id": "4", "left": "6", "right": "7", "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "tree2": {
//     "nodes": [
//       {"id": "1", "left": "8", "right": "2", "value": 1},
//       {"id": "2", "left": "5", "right": "3", "value": 2},
//       {"id": "3", "left": "6", "right": "4", "value": 3},
//       {"id": "4", "left": "7", "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 11
//
// {
//   "tree1": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "8", "value": 1},
//       {"id": "2", "left": "3", "right": "5", "value": 2},
//       {"id": "3", "left": "4", "right": "6", "value": 3},
//       {"id": "4", "left": null, "right": "7", "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "tree2": {
//     "nodes": [
//       {"id": "1", "left": "7", "right": "2", "value": 1},
//       {"id": "2", "left": "6", "right": "3", "value": 2},
//       {"id": "3", "left": "5", "right": "4", "value": 3},
//       {"id": "4", "left": "8", "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   }
// }
//
// Test Case 12
//
// {
//   "tree1": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "8", "value": 1},
//       {"id": "2", "left": "3", "right": "5", "value": 2},
//       {"id": "3", "left": "4", "right": "6", "value": 3},
//       {"id": "4", "left": null, "right": "7", "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": "1-2", "value": 8},
//       {"id": "1-2", "left": null, "right": null, "value": 1}
//     ],
//     "root": "1"
//   },
//   "tree2": {
//     "nodes": [
//       {"id": "1", "left": "7", "right": "2", "value": 1},
//       {"id": "2", "left": "6", "right": "3", "value": 2},
//       {"id": "3", "left": "5", "right": "4", "value": 3},
//       {"id": "4", "left": "8", "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": "1-2", "value": 8},
//       {"id": "1-2", "left": null, "right": null, "value": 1}
//     ],
//     "root": "1"
//   }
// }
