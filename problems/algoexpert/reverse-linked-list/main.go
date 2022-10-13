package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ReverseLinkedList(head *LinkedList) *LinkedList {
	prev, curr := (*LinkedList)(nil), head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
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
//   }
// }
// Test Case 2
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": null, "value": 0}
//     ]
//   }
// }
// Test Case 3
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": null, "value": 1}
//     ]
//   }
// }
// Test Case 4
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": null, "value": 2}
//     ]
//   }
// }
// Test Case 5
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 0},
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": null, "value": 3}
//     ]
//   }
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
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": null, "value": 6}
//     ]
//   }
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
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": "10", "value": 9},
//       {"id": "10", "next": "11", "value": 10},
//       {"id": "11", "next": "12", "value": 11},
//       {"id": "12", "next": null, "value": 12}
//     ]
//   }
// }
