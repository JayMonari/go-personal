package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(lis *LinkedList) *LinkedList {
	for node := lis; node != nil; {
		nonDupNode := node.Next
		for nonDupNode != nil && nonDupNode.Value == node.Value {
			nonDupNode = nonDupNode.Next
		}
		node.Next = nonDupNode
		node = nonDupNode
	}
	return lis
}

// Test Case 1
// {
//   "linkedList": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": "1-2", "value": 1},
//       {"id": "1-2", "next": "1-3", "value": 1},
//       {"id": "1-3", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 3},
//       {"id": "3", "next": "3-2", "value": 4},
//       {"id": "3-2", "next": "3-3", "value": 4},
//       {"id": "3-3", "next": "4", "value": 4},
//       {"id": "4", "next": "5", "value": 5},
//       {"id": "5", "next": "5-2", "value": 6},
//       {"id": "5-2", "next": null, "value": 6}
//     ]
//   }
// }
// Test Case 2
// {
//   "linkedList": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": "1-2", "value": 1},
//       {"id": "1-2", "next": "1-3", "value": 1},
//       {"id": "1-3", "next": "1-4", "value": 1},
//       {"id": "1-4", "next": "1-5", "value": 1},
//       {"id": "1-5", "next": "4", "value": 1},
//       {"id": "4", "next": "4-2", "value": 4},
//       {"id": "4-2", "next": "5", "value": 4},
//       {"id": "5", "next": "6", "value": 5},
//       {"id": "6", "next": "6-2", "value": 6},
//       {"id": "6-2", "next": null, "value": 6}
//     ]
//   }
// }
// Test Case 3
// {
//   "linkedList": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": "1-2", "value": 1},
//       {"id": "1-2", "next": "1-3", "value": 1},
//       {"id": "1-3", "next": "1-4", "value": 1},
//       {"id": "1-4", "next": "1-5", "value": 1},
//       {"id": "1-5", "next": "1-6", "value": 1},
//       {"id": "1-6", "next": "1-7", "value": 1},
//       {"id": "1-7", "next": null, "value": 1}
//     ]
//   }
// }
// Test Case 4
// {
//   "linkedList": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": "9", "value": 1},
//       {"id": "9", "next": "11", "value": 9},
//       {"id": "11", "next": "15", "value": 11},
//       {"id": "15", "next": "15-2", "value": 15},
//       {"id": "15-2", "next": "16", "value": 15},
//       {"id": "16", "next": "17", "value": 16},
//       {"id": "17", "next": null, "value": 17}
//     ]
//   }
// }
// Test Case 5
// {
//   "linkedList": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": null, "value": 1}
//     ]
//   }
// }
// Test Case 6
// {
//   "linkedList": {
//     "head": "-5",
//     "nodes": [
//       {"id": "-5", "next": "-1", "value": -5},
//       {"id": "-1", "next": "-1-2", "value": -1},
//       {"id": "-1-2", "next": "-1-3", "value": -1},
//       {"id": "-1-3", "next": "5", "value": -1},
//       {"id": "5", "next": "5-2", "value": 5},
//       {"id": "5-2", "next": "5-3", "value": 5},
//       {"id": "5-3", "next": "8", "value": 5},
//       {"id": "8", "next": "8-2", "value": 8},
//       {"id": "8-2", "next": "9", "value": 8},
//       {"id": "9", "next": "10", "value": 9},
//       {"id": "10", "next": "11", "value": 10},
//       {"id": "11", "next": "11-2", "value": 11},
//       {"id": "11-2", "next": null, "value": 11}
//     ]
//   }
// }
// Test Case 7
// {
//   "linkedList": {
//     "head": "1",
//     "nodes": [
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
//       {"id": "12", "next": "12-2", "value": 12},
//       {"id": "12-2", "next": null, "value": 12}
//     ]
//   }
// }
