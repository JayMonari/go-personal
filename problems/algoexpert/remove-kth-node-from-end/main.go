package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	count, first, second := 1, head, head
	for count <= k {
		count++
		second = second.Next
	}
	if second == nil {
		head.Value = head.Next.Value
		head.Next = head.Next.Next
		return
	}
	for second.Next != nil {
		second = second.Next
		first = first.Next
	}
	first.Next = first.Next.Next
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
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": null, "value": 9}
//     ]
//   },
//   "k": 4
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
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": null, "value": 9}
//     ]
//   },
//   "k": 1
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
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": null, "value": 9}
//     ]
//   },
//   "k": 2
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
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": null, "value": 9}
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
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": null, "value": 9}
//     ]
//   },
//   "k": 5
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
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": null, "value": 9}
//     ]
//   },
//   "k": 6
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
//       {"id": "9", "next": null, "value": 9}
//     ]
//   },
//   "k": 7
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
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": null, "value": 9}
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
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": null, "value": 9}
//     ]
//   },
//   "k": 9
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
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": null, "value": 9}
//     ]
//   },
//   "k": 10
// }
