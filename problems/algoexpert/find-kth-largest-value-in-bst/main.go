package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func FindKthLargestValueInBst(tree *BST, k int) int {
	nodeVals := inOrderTraverse(tree, []int{})
	return nodeVals[len(nodeVals)-k]
}

func inOrderTraverse(node *BST, vals []int) []int {
	if node == nil {
		return vals
	}
	vals = inOrderTraverse(node.Left, vals)
	vals = append(vals, node.Value)
	vals = inOrderTraverse(node.Right, vals)
	return vals
}

type treeInfo struct {
	nodesVisited  int
	lastNodeValue int
}

func FindKthLargestValueInBstOpt(tree *BST, k int) int {
	ti := treeInfo{nodesVisited: 0, lastNodeValue: -1}
	reverseInOrderTraverse(tree, k, &ti)
	return ti.lastNodeValue
}

func reverseInOrderTraverse(node *BST, k int, ti *treeInfo) {
	if node == nil || ti.nodesVisited >= k {
		return
	}
	reverseInOrderTraverse(node.Right, k, ti)
	if ti.nodesVisited >= k {
		return
	}
	ti.nodesVisited++
	ti.lastNodeValue = node.Value
	reverseInOrderTraverse(node.Left, k, ti)
}

// Test Case 1
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22]
// }
// Test Case 2
// {
//   "array": [1]
// }
// Test Case 3
// {
//   "array": [1, 2]
// }
// Test Case 4
// {
//   "array": [1, 2, 5]
// }
// Test Case 5
// {
//   "array": [1, 2, 5, 7]
// }
// Test Case 6
// {
//   "array": [1, 2, 5, 7, 10]
// }
// Test Case 7
// {
//   "array": [1, 2, 5, 7, 10, 13]
// }
// Test Case 8
// {
//   "array": [1, 2, 5, 7, 10, 13, 14]
// }
// Test Case 9
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15]
// }
// Test Case 10
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22]
// }
// Test Case 11
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28]
// }
// Test Case 12
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32]
// }
// Test Case 13
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36]
// }
// Test Case 14
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36, 89]
// }
// Test Case 15
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36, 89, 92]
// }
// Test Case 16
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36, 89, 92, 9000]
// }
// Test Case 17
// {
//   "array": [1, 2, 5, 7, 10, 13, 14, 15, 22, 28, 32, 36, 89, 92, 9000, 9001]
// }
