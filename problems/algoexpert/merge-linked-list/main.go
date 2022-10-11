package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MergeLinkedLists(l1, l2 *LinkedList) *LinkedList {
	p1 := l1
	var p1Prev *LinkedList
	p2 := l2
	for p1 != nil && p2 != nil {
		if p1.Value < p2.Value {
			p1Prev = p1
			p1 = p1.Next
			continue
		}

		if p1Prev != nil {
			p1Prev.Next = p2
		}
		p1Prev = p2
		p2 = p2.Next
		p1Prev.Next = p1
	}

	if p1 == nil {
		p1Prev.Next = p2
	}

	if l1.Value < l2.Value {
		return l1
	}
	return l2
}

// Test Case 1
// {
//   "linkedListOne": {
//     "head": "2",
//     "nodes": [
//       {"id": "2", "next": "6", "value": 2},
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": null, "value": 8}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": "3", "value": 1},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": "9", "value": 5},
//       {"id": "9", "next": "10", "value": 9},
//       {"id": "10", "next": null, "value": 10}
//     ]
//   }
// }
// Test Case 2
// {
//   "linkedListOne": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "6",
//     "nodes": [
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": "10", "value": 9},
//       {"id": "10", "next": null, "value": 10}
//     ]
//   }
// }
// Test Case 3
// {
//   "linkedListOne": {
//     "head": "6",
//     "nodes": [
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": "10", "value": 9},
//       {"id": "10", "next": null, "value": 10}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   }
// }
// Test Case 4
// {
//   "linkedListOne": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": "3", "value": 1},
//       {"id": "3", "next": "5", "value": 3},
//       {"id": "5", "next": "7", "value": 5},
//       {"id": "7", "next": "9", "value": 7},
//       {"id": "9", "next": null, "value": 9}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "2",
//     "nodes": [
//       {"id": "2", "next": "4", "value": 2},
//       {"id": "4", "next": "6", "value": 4},
//       {"id": "6", "next": "8", "value": 6},
//       {"id": "8", "next": "10", "value": 8},
//       {"id": "10", "next": null, "value": 10}
//     ]
//   }
// }
// Test Case 5
// {
//   "linkedListOne": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": "7", "value": 5},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": "10", "value": 9},
//       {"id": "10", "next": null, "value": 10}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "6",
//     "nodes": [
//       {"id": "6", "next": null, "value": 6}
//     ]
//   }
// }
// Test Case 6
// {
//   "linkedListOne": {
//     "head": "6",
//     "nodes": [
//       {"id": "6", "next": null, "value": 6}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": "7", "value": 5},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": "10", "value": 9},
//       {"id": "10", "next": null, "value": 10}
//     ]
//   }
// }
// Test Case 7
// {
//   "linkedListOne": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": null, "value": 1}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "2",
//     "nodes": [
//       {"id": "2", "next": null, "value": 2}
//     ]
//   }
// }
// Test Case 8
// {
//   "linkedListOne": {
//     "head": "2",
//     "nodes": [
//       {"id": "2", "next": null, "value": 2}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": null, "value": 1}
//     ]
//   }
// }
// Test Case 9
// {
//   "linkedListOne": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": "1-2", "value": 1},
//       {"id": "1-2", "next": "1-3", "value": 1},
//       {"id": "1-3", "next": "3", "value": 1},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": "5-2", "value": 5},
//       {"id": "5-2", "next": "5-3", "value": 5},
//       {"id": "5-3", "next": "5-4", "value": 5},
//       {"id": "5-4", "next": "10", "value": 5},
//       {"id": "10", "next": null, "value": 10}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": "1-2", "value": 1},
//       {"id": "1-2", "next": "2", "value": 1},
//       {"id": "2", "next": "2-2", "value": 2},
//       {"id": "2-2", "next": "5", "value": 2},
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "10", "value": 6},
//       {"id": "10", "next": "10-2", "value": 10},
//       {"id": "10-2", "next": null, "value": 10}
//     ]
//   }
// }
