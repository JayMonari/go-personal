package main

type BST struct {
	Value       int
	Left, Right *BST
}

func (t *BST) Insert(value int) *BST {
	node := t
	for {
		if value < node.Value {
			if node.Left != nil {
				node = node.Left
				continue
			}
			node.Left = &BST{Value: value}
			return t
		}

		if node.Right != nil {
			node = node.Right
			continue
		}
		node.Right = &BST{Value: value}
		return t
	}
}

func (t *BST) Contains(value int) bool {
	node := t
	for node != nil {
		switch {
		case value < node.Value:
			node = node.Left
		case value > node.Value:
			node = node.Right
		default:
			return true
		}
	}
	return false
}

func (t *BST) Remove(value int) *BST {
	t.remove(value, nil)
	return t
}

func (t *BST) remove(value int, parent *BST) {
	node := t
	for node != nil && value != node.Value {
		if value < node.Value {
			parent = node
			node = node.Left
			continue
		}
		parent = node
		node = node.Right
	}
	if node == nil {
		return
	}

	switch {
	case node.Left != nil && node.Right != nil:
		node.Value = node.Right.minVal()
		node.Right.remove(node.Value, node)
	case parent == nil:
		if node.Left != nil {
			node.Value = node.Left.Value
			node.Right = node.Left.Right
			node.Left = node.Left.Left
		}
		if node.Right != nil {
			node.Value = node.Right.Value
			node.Left = node.Right.Left
			node.Right = node.Right.Right
		}
	case parent.Left == node:
		if node.Left != nil {
			parent.Left = node.Left
			break
		}
		parent.Left = node.Right
	case parent.Right == node:
		if node.Left != nil {
			parent.Right = node.Left
			break
		}
		parent.Right = node.Right
	}
}

// func (tree *BST) remove(value int, parent *BST) {
// 	current := tree
// 	for current != nil {
// 		if value < current.Value {
// 			parent = current
// 			current = current.Left
// 		} else if value > current.Value {
// 			parent = current
// 			current = current.Right
// 		} else {
// 			if current.Left != nil && current.Right != nil {
// 				current.Value = current.Right.minVal()
// 				current.Right.remove(current.Value, current)
// 			} else if parent == nil {
// 				if current.Left != nil {
// 					current.Value = current.Left.Value
// 					current.Right = current.Left.Right
// 					current.Left = current.Left.Left
// 				} else if current.Right != nil {
// 					current.Value = current.Right.Value
// 					current.Left = current.Right.Left
// 					current.Right = current.Right.Right
// 				} else {
// 					// This is a single-node tree; do nothing.
// 				}
// 			} else if parent.Left == current {
// 				if current.Left != nil {
// 					parent.Left = current.Left
// 				} else {
// 					parent.Left = current.Right
// 				}
// 			} else if parent.Right == current {
// 				if current.Left != nil {
// 					parent.Right = current.Left
// 				} else {
// 					parent.Right = current.Right
// 				}
// 			}
// 			break
// 		}
// 	}
// }

func (t *BST) minVal() int {
	if t.Left == nil {
		return t.Value
	}
	return t.Left.minVal()
}

// Test Case 1
// {
//   "rootValue": 10,
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [15],
//       "method": "insert"
//     },
//     {
//       "arguments": [2],
//       "method": "insert"
//     },
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [13],
//       "method": "insert"
//     },
//     {
//       "arguments": [22],
//       "method": "insert"
//     },
//     {
//       "arguments": [1],
//       "method": "insert"
//     },
//     {
//       "arguments": [14],
//       "method": "insert"
//     },
//     {
//       "arguments": [12],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "remove"
//     },
//     {
//       "arguments": [15],
//       "method": "contains"
//     }
//   ]
// }
// Test Case 2
// {
//   "rootValue": 10,
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [15],
//       "method": "insert"
//     }
//   ]
// }
// Test Case 3
// {
//   "rootValue": 10,
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [15],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "contains"
//     },
//     {
//       "arguments": [5],
//       "method": "contains"
//     },
//     {
//       "arguments": [15],
//       "method": "contains"
//     },
//     {
//       "arguments": [1],
//       "method": "contains"
//     },
//     {
//       "arguments": [6],
//       "method": "contains"
//     },
//     {
//       "arguments": [11],
//       "method": "contains"
//     },
//     {
//       "arguments": [16],
//       "method": "contains"
//     }
//   ]
// }
// Test Case 4
// {
//   "rootValue": 10,
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [15],
//       "method": "insert"
//     },
//     {
//       "arguments": [5],
//       "method": "remove"
//     },
//     {
//       "arguments": [15],
//       "method": "remove"
//     },
//     {
//       "arguments": [10],
//       "method": "remove"
//     }
//   ]
// }
// Test Case 5
// {
//   "rootValue": 10,
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [15],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "contains"
//     },
//     {
//       "arguments": [5],
//       "method": "contains"
//     },
//     {
//       "arguments": [15],
//       "method": "contains"
//     },
//     {
//       "arguments": [10],
//       "method": "remove"
//     },
//     {
//       "arguments": [5],
//       "method": "remove"
//     },
//     {
//       "arguments": [15],
//       "method": "remove"
//     },
//     {
//       "arguments": [10],
//       "method": "contains"
//     },
//     {
//       "arguments": [5],
//       "method": "contains"
//     },
//     {
//       "arguments": [15],
//       "method": "contains"
//     }
//   ]
// }
// Test Case 6
// {
//   "rootValue": 1,
//   "classMethodsToCall": [
//     {
//       "arguments": [2],
//       "method": "insert"
//     },
//     {
//       "arguments": [3],
//       "method": "insert"
//     },
//     {
//       "arguments": [4],
//       "method": "insert"
//     },
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [6],
//       "method": "insert"
//     },
//     {
//       "arguments": [7],
//       "method": "insert"
//     },
//     {
//       "arguments": [8],
//       "method": "insert"
//     },
//     {
//       "arguments": [9],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "insert"
//     },
//     {
//       "arguments": [11],
//       "method": "insert"
//     },
//     {
//       "arguments": [12],
//       "method": "insert"
//     },
//     {
//       "arguments": [13],
//       "method": "insert"
//     },
//     {
//       "arguments": [14],
//       "method": "insert"
//     },
//     {
//       "arguments": [15],
//       "method": "insert"
//     },
//     {
//       "arguments": [16],
//       "method": "insert"
//     },
//     {
//       "arguments": [17],
//       "method": "insert"
//     },
//     {
//       "arguments": [18],
//       "method": "insert"
//     },
//     {
//       "arguments": [19],
//       "method": "insert"
//     },
//     {
//       "arguments": [20],
//       "method": "insert"
//     },
//     {
//       "arguments": [2],
//       "method": "remove"
//     },
//     {
//       "arguments": [4],
//       "method": "remove"
//     },
//     {
//       "arguments": [6],
//       "method": "remove"
//     },
//     {
//       "arguments": [8],
//       "method": "remove"
//     },
//     {
//       "arguments": [11],
//       "method": "remove"
//     },
//     {
//       "arguments": [13],
//       "method": "remove"
//     },
//     {
//       "arguments": [15],
//       "method": "remove"
//     },
//     {
//       "arguments": [17],
//       "method": "remove"
//     },
//     {
//       "arguments": [19],
//       "method": "remove"
//     },
//     {
//       "arguments": [1],
//       "method": "insert"
//     },
//     {
//       "arguments": [2],
//       "method": "insert"
//     },
//     {
//       "arguments": [3],
//       "method": "insert"
//     },
//     {
//       "arguments": [4],
//       "method": "insert"
//     },
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [6],
//       "method": "insert"
//     },
//     {
//       "arguments": [7],
//       "method": "insert"
//     },
//     {
//       "arguments": [8],
//       "method": "insert"
//     },
//     {
//       "arguments": [9],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "insert"
//     },
//     {
//       "arguments": [9000],
//       "method": "contains"
//     }
//   ]
// }
// Test Case 7
// {
//   "rootValue": 1,
//   "classMethodsToCall": [
//     {
//       "arguments": [2],
//       "method": "insert"
//     },
//     {
//       "arguments": [3],
//       "method": "insert"
//     },
//     {
//       "arguments": [4],
//       "method": "insert"
//     },
//     {
//       "arguments": [1],
//       "method": "remove"
//     }
//   ]
// }
// Test Case 8
// {
//   "rootValue": 1,
//   "classMethodsToCall": [
//     {
//       "arguments": [-2],
//       "method": "insert"
//     },
//     {
//       "arguments": [-3],
//       "method": "insert"
//     },
//     {
//       "arguments": [-4],
//       "method": "insert"
//     },
//     {
//       "arguments": [1],
//       "method": "remove"
//     }
//   ]
// }
// Test Case 9
// {
//   "rootValue": 10,
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "remove"
//     },
//     {
//       "arguments": [15],
//       "method": "contains"
//     }
//   ]
// }
// Test Case 10
// {
//   "rootValue": 10,
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [15],
//       "method": "insert"
//     },
//     {
//       "arguments": [2],
//       "method": "insert"
//     },
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [13],
//       "method": "insert"
//     },
//     {
//       "arguments": [22],
//       "method": "insert"
//     },
//     {
//       "arguments": [1],
//       "method": "insert"
//     },
//     {
//       "arguments": [14],
//       "method": "insert"
//     },
//     {
//       "arguments": [12],
//       "method": "insert"
//     },
//     {
//       "arguments": [5],
//       "method": "remove"
//     },
//     {
//       "arguments": [5],
//       "method": "remove"
//     },
//     {
//       "arguments": [12],
//       "method": "remove"
//     },
//     {
//       "arguments": [13],
//       "method": "remove"
//     },
//     {
//       "arguments": [14],
//       "method": "remove"
//     },
//     {
//       "arguments": [22],
//       "method": "remove"
//     },
//     {
//       "arguments": [2],
//       "method": "remove"
//     },
//     {
//       "arguments": [1],
//       "method": "remove"
//     },
//     {
//       "arguments": [15],
//       "method": "contains"
//     }
//   ]
// }
