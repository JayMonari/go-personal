package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func FindNodesDistanceK(t *BinaryTree, target, k int) []int {
	parents := map[int]*BinaryTree{}
	fetchParents(t, parents, nil)

	type btInfo struct {
		node     *BinaryTree
		distance int
	}
	targetNode := getNode(target, t, parents)
	q := []btInfo{{node: targetNode}}
	seen := map[int]struct{}{targetNode.Value: {}}
	var node *BinaryTree
	var dist int
	resNodes := []int{}
	for len(q) > 0 {
		node, dist, q = q[0].node, q[0].distance, q[1:]

		if dist == k {
			for _, i := range q {
				resNodes = append(resNodes, i.node.Value)
			}
			resNodes = append(resNodes, node.Value)
			break
		}

		for _, n := range []*BinaryTree{node.Left, node.Right, parents[node.Value]} {
			if n == nil {
				continue
			}
			if _, ok := seen[n.Value]; ok {
				continue
			}
			seen[n.Value] = struct{}{}
			q = append(q, btInfo{node: n, distance: dist + 1})
		}
	}
	return resNodes
}

func getNode(from int, root *BinaryTree, parents map[int]*BinaryTree) *BinaryTree {
	if root.Value == from {
		return root
	}
	p := parents[from]
	if p.Left != nil && p.Left.Value == from {
		return p.Left
	}
	return p.Right
}

func fetchParents(node *BinaryTree, parents map[int]*BinaryTree, parent *BinaryTree) {
	if node == nil {
		return
	}
	parents[node.Value] = parent
	fetchParents(node.Left, parents, node)
	fetchParents(node.Right, parents, node)
}

// Test Case 1
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": null, "right": "6", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": "7", "right": "8", "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "target": 3,
//   "k": 2
// }
// Test Case 2
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": null, "value": 1},
//       {"id": "2", "left": "3", "right": null, "value": 2},
//       {"id": "3", "left": "4", "right": null, "value": 3},
//       {"id": "4", "left": "5", "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5}
//     ],
//     "root": "1"
//   },
//   "target": 2,
//   "k": 3
// }
// Test Case 3
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": null, "value": 2},
//       {"id": "3", "left": "5", "right": "6", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": "7", "value": 6},
//       {"id": "7", "left": null, "right": "8", "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "target": 8,
//   "k": 6
// }
// Test Case 4
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": null, "value": 2},
//       {"id": "3", "left": "5", "right": "6", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": null, "right": "7", "value": 6},
//       {"id": "7", "left": null, "right": "8", "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "target": 3,
//   "k": 1
// }
// Test Case 5
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
//   },
//   "target": 1,
//   "k": 2
// }
// Test Case 6
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
//   },
//   "target": 8,
//   "k": 2
// }
// Test Case 7
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": null, "value": 2},
//       {"id": "3", "left": null, "right": "5", "value": 3},
//       {"id": "4", "left": "6", "right": null, "value": 4},
//       {"id": "5", "left": "7", "right": "8", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "target": 6,
//   "k": 6
// }
// Test Case 8
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "1"
//   },
//   "target": 1,
//   "k": 1
// }
// Test Case 9
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": null, "value": 2},
//       {"id": "3", "left": null, "right": "5", "value": 3},
//       {"id": "4", "left": "6", "right": null, "value": 4},
//       {"id": "5", "left": "7", "right": "8", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "1"
//   },
//   "target": 6,
//   "k": 17
// }
// Test Case 10
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": "2", "right": "3", "value": 1},
//       {"id": "2", "left": "4", "right": "5", "value": 2},
//       {"id": "3", "left": "6", "right": "7", "value": 3},
//       {"id": "4", "left": "8", "right": "9", "value": 4},
//       {"id": "5", "left": "10", "right": "11", "value": 5},
//       {"id": "6", "left": "12", "right": "13", "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9},
//       {"id": "10", "left": null, "right": null, "value": 10},
//       {"id": "11", "left": null, "right": null, "value": 11},
//       {"id": "12", "left": null, "right": null, "value": 12},
//       {"id": "13", "left": null, "right": null, "value": 13}
//     ],
//     "root": "1"
//   },
//   "target": 2,
//   "k": 2
// }
