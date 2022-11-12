package main

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func ReconstructBst(preOrderValues []int) *BST {
	if len(preOrderValues) == 0 {
		return nil
	}
	val := preOrderValues[0]
	rightSubtreeIdx := len(preOrderValues)
	for i := 1; i < len(preOrderValues); i++ {
		if preOrderValues[i] >= val {
			rightSubtreeIdx = i
			break
		}
	}
	return &BST{
		Value: val,
		Left:  ReconstructBst(preOrderValues[1:rightSubtreeIdx]),
		Right: ReconstructBst(preOrderValues[rightSubtreeIdx:]),
	}
}

type treeInfo struct {
	rootIdx int
}

// O(n) time | O(n) space - where n is the length of the input array
func ReconstructBstOpt(preOrderTraversalValues []int) *BST {
	treeInfo := &treeInfo{rootIdx: 0}
	return reconstructBstFromRange(
		math.MinInt32,
		math.MaxInt32,
		preOrderTraversalValues,
		treeInfo,
	)
}

func reconstructBstFromRange(lowerBound, upperBound int, preOrderTraversalValues []int, currentSubtreeInfo *treeInfo) *BST {
	if currentSubtreeInfo.rootIdx == len(preOrderTraversalValues) {
		return nil
	}

	rootValue := preOrderTraversalValues[currentSubtreeInfo.rootIdx]
	if rootValue < lowerBound || rootValue >= upperBound {
		return nil
	}

	currentSubtreeInfo.rootIdx++
	leftSubtree := reconstructBstFromRange(
		lowerBound,
		rootValue,
		preOrderTraversalValues,
		currentSubtreeInfo,
	)
	rightSubtree := reconstructBstFromRange(
		rootValue,
		upperBound,
		preOrderTraversalValues,
		currentSubtreeInfo,
	)
	return &BST{Value: rootValue, Left: leftSubtree, Right: rightSubtree}
}

// Test Case 1
// {
//   "preOrderTraversalValues": [10, 4, 2, 1, 5, 17, 19, 18]
// }
// Test Case 2
// {
//   "preOrderTraversalValues": [100]
// }
// Test Case 3
// {
//   "preOrderTraversalValues": [10, 9, 8, 7, 6, 5]
// }
// Test Case 4
// {
//   "preOrderTraversalValues": [5, 6, 7, 8]
// }
// Test Case 5
// {
//   "preOrderTraversalValues": [5, -10, -5, 6, 9, 7]
// }
// Test Case 6
// {
//   "preOrderTraversalValues": [10, 4, 2, 1, 3, 5, 6, 9, 7, 17, 19, 18]
// }
// Test Case 7
// {
//   "preOrderTraversalValues": [1, 0, 2]
// }
// Test Case 8
// {
//   "preOrderTraversalValues": [2, 0, 1]
// }
// Test Case 9
// {
//   "preOrderTraversalValues": [2, 0, 1, 4, 3]
// }
// Test Case 10
// {
//   "preOrderTraversalValues": [2, 0, 1, 4, 3, 3]
// }
// Test Case 11
// {
//   "preOrderTraversalValues": [2, 0, 1, 3, 4, 3]
// }
// Test Case 12
// {
//   "preOrderTraversalValues": [10, 4, 2, 1, 3, 5, 5, 6, 5, 5, 9, 7, 17, 19, 18, 18, 19]
// }
