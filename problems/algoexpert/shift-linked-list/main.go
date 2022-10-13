package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ShiftLinkedList(head *LinkedList, by int) *LinkedList {
	listLen := 1
	listTail := head
	for listTail.Next != nil {
		listTail = listTail.Next
		listLen++
	}

	offset := abs(by) % listLen
	if offset == 0 {
		return head
	}

	newTailPosition := listLen - offset
	if by <= 0 {
		newTailPosition = offset
	}

	newTail := head
	for i := 1; i < newTailPosition; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil
	listTail.Next = head
	return newHead
}

func abs(k int) int {
	if k > 0 {
		return k
	}
	return -k
}

// Test Case 1
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": 2
// }
// Test Case 2
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": 0
// }
// Test Case 3
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": 1
// }
// Test Case 4
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": 3
// }
// Test Case 5
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": 4
// }
// Test Case 6
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": 5
// }
// Test Case 7
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": 6
// }
// Test Case 8
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": 8
// }
// Test Case 9
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": 14
// }
// Test Case 10
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": 18
// }
// Test Case 11
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": -1
// }
// Test Case 12
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": -2
// }
// Test Case 13
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": -3
// }
// Test Case 14
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": -4
// }
// Test Case 15
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": -5
// }
// Test Case 16
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": -6
// }
// Test Case 17
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": -8
// }
// Test Case 18
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": -14
// }
// Test Case 19
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "k": -18
// }
// Test Case 20
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "4", "value": 1},
//       {"id": "2", "next": null, "value": 2},
//       {"id": "3", "next": "5", "value": 3},
//       {"id": "4", "next": "3", "value": 4},
//       {"id": "5", "next": "2", "value": 5}
//     ]
//   },
//   "k": 2
// }
