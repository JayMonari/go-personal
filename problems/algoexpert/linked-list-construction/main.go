package main

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.InsertBefore(ll.Head, node)
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	if ll.Tail == nil {
		ll.SetHead(node)
		return
	}
	ll.InsertAfter(ll.Tail, node)
}

func (ll *DoublyLinkedList) InsertBefore(node, insert *Node) {
	if insert == ll.Head && insert == ll.Tail {
		return
	}
	ll.Remove(insert)
	insert.Prev = node.Prev
	insert.Next = node
	if node.Prev == nil {
		ll.Head = insert
	} else {
		node.Prev.Next = insert
	}
	node.Prev = insert
}

func (ll *DoublyLinkedList) InsertAfter(node, insert *Node) {
	if insert == ll.Head && insert == ll.Tail {
		return
	}
	ll.Remove(insert)
	insert.Prev = node
	insert.Next = node.Next
	if node.Next == nil {
		ll.Tail = insert
	} else {
		node.Next.Prev = insert
	}
	node.Next = insert
}

func (ll *DoublyLinkedList) InsertAtPosition(pos int, n *Node) {
	if pos == 1 {
		ll.SetHead(n)
		return
	}
	node := ll.Head
	currPos := 1
	for node != nil && currPos != pos {
		node = node.Next
		currPos++
	}
	if node == nil {
		ll.SetTail(n)
		return
	}
	ll.InsertBefore(node, n)
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	node := ll.Head
	for node != nil {
		removeNode := node
		node = node.Next
		if removeNode.Value != value {
			continue
		}
		ll.Remove(removeNode)
	}
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	if node == ll.Head {
		ll.Head = ll.Head.Next
	}
	if node == ll.Tail {
		ll.Tail = ll.Tail.Prev
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	node := ll.Head
	for node != nil && node.Value != value {
		node = node.Next
	}
	return node != nil
}

// Test Case 1
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2},
//     {"id": "3", "next": null, "prev": null, "value": 3},
//     {"id": "3-2", "next": null, "prev": null, "value": 3},
//     {"id": "3-3", "next": null, "prev": null, "value": 3},
//     {"id": "4", "next": null, "prev": null, "value": 4},
//     {"id": "5", "next": null, "prev": null, "value": 5},
//     {"id": "6", "next": null, "prev": null, "value": 6}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["5"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["4"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["3"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["2"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["4"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["6"],
//       "method": "setTail"
//     },
//     {
//       "arguments": ["6", "3"],
//       "method": "insertBefore"
//     },
//     {
//       "arguments": ["6", "3-2"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": [1, "3-3"],
//       "method": "insertAtPosition"
//     },
//     {
//       "arguments": [3],
//       "method": "removeNodesWithValue"
//     },
//     {
//       "arguments": ["2"],
//       "method": "remove"
//     },
//     {
//       "arguments": [5],
//       "method": "containsNodeWithValue"
//     }
//   ]
// }
// Test Case 2
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     }
//   ]
// }
// Test Case 3
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setTail"
//     }
//   ]
// }
// Test Case 4
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": [1, "1"],
//       "method": "insertAtPosition"
//     }
//   ]
// }
// Test Case 5
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["2"],
//       "method": "setTail"
//     }
//   ]
// }
// Test Case 6
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["2"],
//       "method": "setHead"
//     }
//   ]
// }
// Test Case 7
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertAfter"
//     }
//   ]
// }
// Test Case 8
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertBefore"
//     }
//   ]
// }
// Test Case 9
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2},
//     {"id": "3", "next": null, "prev": null, "value": 3},
//     {"id": "4", "next": null, "prev": null, "value": 4}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["2", "3"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["3", "4"],
//       "method": "insertAfter"
//     }
//   ]
// }
// Test Case 10
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2},
//     {"id": "3", "next": null, "prev": null, "value": 3},
//     {"id": "4", "next": null, "prev": null, "value": 4}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setTail"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertBefore"
//     },
//     {
//       "arguments": ["2", "3"],
//       "method": "insertBefore"
//     },
//     {
//       "arguments": ["3", "4"],
//       "method": "insertBefore"
//     }
//   ]
// }
// Test Case 11
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2},
//     {"id": "3", "next": null, "prev": null, "value": 3},
//     {"id": "4", "next": null, "prev": null, "value": 4}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["2", "3"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["3", "4"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["1"],
//       "method": "setTail"
//     }
//   ]
// }
// Test Case 12
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2},
//     {"id": "3", "next": null, "prev": null, "value": 3},
//     {"id": "4", "next": null, "prev": null, "value": 4}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setTail"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertBefore"
//     },
//     {
//       "arguments": ["2", "3"],
//       "method": "insertBefore"
//     },
//     {
//       "arguments": ["3", "4"],
//       "method": "insertBefore"
//     },
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     }
//   ]
// }
// Test Case 13
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2},
//     {"id": "3", "next": null, "prev": null, "value": 3},
//     {"id": "4", "next": null, "prev": null, "value": 4}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["2", "3"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["3", "4"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["2", "1"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["3", "4"],
//       "method": "insertBefore"
//     }
//   ]
// }
// Test Case 14
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2},
//     {"id": "3", "next": null, "prev": null, "value": 3},
//     {"id": "4", "next": null, "prev": null, "value": 4},
//     {"id": "5", "next": null, "prev": null, "value": 5},
//     {"id": "6", "next": null, "prev": null, "value": 6},
//     {"id": "7", "next": null, "prev": null, "value": 7}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["2", "3"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["3", "4"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["4", "5"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["5", "6"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["6", "7"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": [7, "1"],
//       "method": "insertAtPosition"
//     },
//     {
//       "arguments": [1, "1"],
//       "method": "insertAtPosition"
//     },
//     {
//       "arguments": [2, "1"],
//       "method": "insertAtPosition"
//     },
//     {
//       "arguments": [3, "1"],
//       "method": "insertAtPosition"
//     },
//     {
//       "arguments": [4, "1"],
//       "method": "insertAtPosition"
//     },
//     {
//       "arguments": [5, "1"],
//       "method": "insertAtPosition"
//     },
//     {
//       "arguments": [6, "1"],
//       "method": "insertAtPosition"
//     }
//   ]
// }
// Test Case 15
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1"],
//       "method": "remove"
//     }
//   ]
// }
// Test Case 16
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": [1],
//       "method": "removeNodesWithValue"
//     }
//   ]
// }
// Test Case 17
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2},
//     {"id": "3", "next": null, "prev": null, "value": 3},
//     {"id": "4", "next": null, "prev": null, "value": 4}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["2", "3"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["3", "4"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["1"],
//       "method": "remove"
//     }
//   ]
// }
// Test Case 18
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2},
//     {"id": "3", "next": null, "prev": null, "value": 3},
//     {"id": "4", "next": null, "prev": null, "value": 4}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["2", "3"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["3", "4"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["4"],
//       "method": "remove"
//     }
//   ]
// }
// Test Case 19
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2},
//     {"id": "3", "next": null, "prev": null, "value": 3},
//     {"id": "4", "next": null, "prev": null, "value": 4}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["2", "3"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["3", "4"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["2"],
//       "method": "remove"
//     }
//   ]
// }
// Test Case 20
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "1-2", "next": null, "prev": null, "value": 1},
//     {"id": "1-3", "next": null, "prev": null, "value": 1},
//     {"id": "1-4", "next": null, "prev": null, "value": 1}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1", "1-2"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["1-2", "1-3"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["1-3", "1-4"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": [1],
//       "method": "removeNodesWithValue"
//     }
//   ]
// }
// Test Case 21
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "1-2", "next": null, "prev": null, "value": 1},
//     {"id": "1-3", "next": null, "prev": null, "value": 1},
//     {"id": "1-4", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2},
//     {"id": "3", "next": null, "prev": null, "value": 3},
//     {"id": "4", "next": null, "prev": null, "value": 4}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["2", "1-2"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["1-2", "3"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["3", "1-3"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["1-3", "4"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": [1],
//       "method": "removeNodesWithValue"
//     }
//   ]
// }
// Test Case 22
// {
//   "nodes": [
//     {"id": "1", "next": null, "prev": null, "value": 1},
//     {"id": "2", "next": null, "prev": null, "value": 2},
//     {"id": "3", "next": null, "prev": null, "value": 3},
//     {"id": "4", "next": null, "prev": null, "value": 4}
//   ],
//   "classMethodsToCall": [
//     {
//       "arguments": ["1"],
//       "method": "setHead"
//     },
//     {
//       "arguments": ["1", "2"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["2", "3"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": ["3", "4"],
//       "method": "insertAfter"
//     },
//     {
//       "arguments": [1],
//       "method": "containsNodeWithValue"
//     },
//     {
//       "arguments": [2],
//       "method": "containsNodeWithValue"
//     },
//     {
//       "arguments": [3],
//       "method": "containsNodeWithValue"
//     },
//     {
//       "arguments": [4],
//       "method": "containsNodeWithValue"
//     },
//     {
//       "arguments": [5],
//       "method": "containsNodeWithValue"
//     }
//   ]
// }
