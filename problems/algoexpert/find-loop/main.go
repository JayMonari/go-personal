package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// FindLoop is Floyd's Cycle Detection Algorithm
func FindLoop(head *LinkedList) *LinkedList {
	slow := head.Next
	fast := slow.Next
	for slow != fast {
		slow, fast = slow.Next, fast.Next.Next
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
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
//       {"id": "9", "next": "4", "value": 9}
//     ]
//   }
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
//       {"id": "9", "next": "0", "value": 9}
//     ]
//   }
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
//       {"id": "9", "next": "1", "value": 9}
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
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": "2", "value": 9}
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
//       {"id": "3", "next": "4", "value": 3},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": "3", "value": 9}
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
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": "5", "value": 9}
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
//       {"id": "9", "next": "6", "value": 9}
//     ]
//   }
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
//       {"id": "9", "next": "7", "value": 9}
//     ]
//   }
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
//       {"id": "9", "next": "8", "value": 9}
//     ]
//   }
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
//       {"id": "9", "next": "9", "value": 9}
//     ]
//   }
// }
// Test Case 11
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "1", "value": 5},
//       {"id": "1", "next": "2", "value": 4},
//       {"id": "2", "next": "3", "value": 3},
//       {"id": "3", "next": "2", "value": 2}
//     ]
//   }
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
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "2-2", "value": 6},
//       {"id": "2-2", "next": "7", "value": 2},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": "4", "value": 9}
//     ]
//   }
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
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "2-2", "value": 6},
//       {"id": "2-2", "next": "7", "value": 2},
//       {"id": "7", "next": "8", "value": 7},
//       {"id": "8", "next": "9", "value": 8},
//       {"id": "9", "next": "2-2", "value": 9}
//     ]
//   }
// }
// Test Case 14
// {
//   "linkedList": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "0-2", "value": 0},
//       {"id": "0-2", "next": "0-3", "value": 0},
//       {"id": "0-3", "next": "0-4", "value": 0},
//       {"id": "0-4", "next": "0-5", "value": 0},
//       {"id": "0-5", "next": "0-6", "value": 0},
//       {"id": "0-6", "next": "0-7", "value": 0},
//       {"id": "0-7", "next": "0-8", "value": 0},
//       {"id": "0-8", "next": "0-9", "value": 0},
//       {"id": "0-9", "next": "0-7", "value": 0}
//     ]
//   }
// }
