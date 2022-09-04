package main

import (
	"strconv"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func SumOfLinkedLists(one *LinkedList, two *LinkedList) *LinkedList {
	sum := &LinkedList{}
	node := sum
	for n := makeNumber(one) + makeNumber(two); ; {
		node.Value = n % 10
		if n /= 10; n == 0 {
			break
		}
		node.Next = &LinkedList{}
		node = node.Next
	}
	return sum
}

func makeNumber(head *LinkedList) int {
	num := []byte{}
	for head != nil {
		num = append(num, byte(head.Value)+'0')
		head = head.Next
	}
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
	n, err := strconv.Atoi(string(num))
	if err != nil {
		return 0
	}
	return n
}

// Test Case 1
// {
//   "linkedListOne": {
//     "head": "2",
//     "nodes": [
//       {"id": "2", "next": "4", "value": 2},
//       {"id": "4", "next": "7", "value": 4},
//       {"id": "7", "next": "1", "value": 7},
//       {"id": "1", "next": null, "value": 1}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "9",
//     "nodes": [
//       {"id": "9", "next": "4", "value": 9},
//       {"id": "4", "next": "5", "value": 4},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   }
// }
// Test Case 2
// {
//   "linkedListOne": {
//     "head": "2",
//     "nodes": [
//       {"id": "2", "next": null, "value": 2}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "9",
//     "nodes": [
//       {"id": "9", "next": null, "value": 9}
//     ]
//   }
// }
// Test Case 3
// {
//   "linkedListOne": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "0-2", "value": 0},
//       {"id": "0-2", "next": "0-3", "value": 0},
//       {"id": "0-3", "next": "5", "value": 0},
//       {"id": "5", "next": null, "value": 5}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "9",
//     "nodes": [
//       {"id": "9", "next": null, "value": 9}
//     ]
//   }
// }
// Test Case 4
// {
//   "linkedListOne": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": "1-2", "value": 1},
//       {"id": "1-2", "next": "1-3", "value": 1},
//       {"id": "1-3", "next": null, "value": 1}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "9",
//     "nodes": [
//       {"id": "9", "next": "9-2", "value": 9},
//       {"id": "9-2", "next": "9-3", "value": 9},
//       {"id": "9-3", "next": null, "value": 9}
//     ]
//   }
// }
// Test Case 5
// {
//   "linkedListOne": {
//     "head": "1",
//     "nodes": [
//       {"id": "1", "next": "2", "value": 1},
//       {"id": "2", "next": "3", "value": 2},
//       {"id": "3", "next": null, "value": 3}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "6",
//     "nodes": [
//       {"id": "6", "next": "7", "value": 6},
//       {"id": "7", "next": "9", "value": 7},
//       {"id": "9", "next": "1", "value": 9},
//       {"id": "1", "next": "8", "value": 1},
//       {"id": "8", "next": null, "value": 8}
//     ]
//   }
// }
// Test Case 6
// {
//   "linkedListOne": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": null, "value": 0}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": null, "value": 0}
//     ]
//   }
// }
// Test Case 7
// {
//   "linkedListOne": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": null, "value": 0}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "0-2", "value": 0},
//       {"id": "0-2", "next": "0-3", "value": 0},
//       {"id": "0-3", "next": "0-4", "value": 0},
//       {"id": "0-4", "next": "0-5", "value": 0},
//       {"id": "0-5", "next": "8", "value": 0},
//       {"id": "8", "next": null, "value": 8}
//     ]
//   }
// }
// Test Case 8
// {
//   "linkedListOne": {
//     "head": "4",
//     "nodes": [
//       {"id": "4", "next": "6", "value": 4},
//       {"id": "6", "next": "9", "value": 6},
//       {"id": "9", "next": "3", "value": 9},
//       {"id": "3", "next": "1", "value": 3},
//       {"id": "1", "next": null, "value": 1}
//     ]
//   },
//   "linkedListTwo": {
//     "head": "0",
//     "nodes": [
//       {"id": "0", "next": "0-2", "value": 0},
//       {"id": "0-2", "next": "0-3", "value": 0},
//       {"id": "0-3", "next": "0-4", "value": 0},
//       {"id": "0-4", "next": "2", "value": 0},
//       {"id": "2", "next": "7", "value": 2},
//       {"id": "7", "next": null, "value": 7}
//     ]
//   }
// }
