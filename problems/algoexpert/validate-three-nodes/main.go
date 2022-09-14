package main

type BST struct {
	Value int

	Left, Right *BST
}

// O(d) time | O(1) space - where d is the distance between t1 and t3
func ValidateThreeNodes(t1 *BST, t2 *BST, t3 *BST) bool {
	node1 := t1
	node3 := t3
	for node1 != t3 || node3 != t1 {
		if node1 == t2 || node3 == t2 || (node1 == nil && node3 == nil) {
			break
		}

		if node1 != nil {
			if node1.Value > t2.Value {
				node1 = node1.Left
				continue
			}
			node1 = node1.Right
		}

		if node3 != nil {
			if node3.Value > t2.Value {
				node3 = node3.Left
				continue
			}
			node3 = node3.Right
		}
	}

	if !(node1 == t2 || node3 == t2) || node1 == t3 || node3 == t1 {
		return false
	}

	if node1 == t2 {
		t1 = t3
	}
	node1 = t2
	for node1 != nil && node1 != t1 {
		if t1.Value < node1.Value {
			node1 = node1.Left
			continue
		}
		node1 = node1.Right
	}
	return node1 == t1
}

// Test Case 1
// {
//   "tree": {
//     "nodes": [
//       {"id": "0", "left": null, "right": null, "value": 0},
//       {"id": "1", "left": "0", "right": null, "value": 1},
//       {"id": "2", "left": "1", "right": "4", "value": 2},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "4", "left": "3", "right": null, "value": 4},
//       {"id": "5", "left": "2", "right": "7", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": "6", "right": "8", "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "5"
//   },
//   "nodeOne": "5",
//   "nodeTwo": "2",
//   "nodeThree": "3"
// }
// Test Case 2
// {
//   "tree": {
//     "nodes": [
//       {"id": "0", "left": null, "right": null, "value": 0},
//       {"id": "1", "left": "0", "right": null, "value": 1},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "3", "left": "2", "right": "4", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": "3", "right": "7", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": "6", "right": "8", "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "5"
//   },
//   "nodeOne": "5",
//   "nodeTwo": "8",
//   "nodeThree": "1"
// }
// Test Case 3
// {
//   "tree": {
//     "nodes": [
//       {"id": "0", "left": null, "right": null, "value": 0},
//       {"id": "1", "left": "0", "right": null, "value": 1},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "3", "left": "2", "right": "4", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": "3", "right": "7", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": "6", "right": "8", "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "5"
//   },
//   "nodeOne": "8",
//   "nodeTwo": "5",
//   "nodeThree": "2"
// }
// Test Case 4
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1},
//       {"id": "2", "left": "1", "right": "3", "value": 2},
//       {"id": "3", "left": null, "right": "4", "value": 3},
//       {"id": "4", "left": null, "right": "5", "value": 4},
//       {"id": "5", "left": null, "right": "6", "value": 5},
//       {"id": "6", "left": null, "right": "7", "value": 6},
//       {"id": "7", "left": null, "right": "8", "value": 7},
//       {"id": "8", "left": null, "right": "9", "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9}
//     ],
//     "root": "2"
//   },
//   "nodeOne": "2",
//   "nodeTwo": "5",
//   "nodeThree": "8"
// }
// Test Case 5
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": "2", "value": 1},
//       {"id": "2", "left": null, "right": null, "value": 2},
//       {"id": "3", "left": "1", "right": null, "value": 3},
//       {"id": "4", "left": "3", "right": null, "value": 4},
//       {"id": "5", "left": null, "right": "5", "value": 5},
//       {"id": "6", "left": "4", "right": "8", "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": "7", "right": "9", "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9}
//     ],
//     "root": "6"
//   },
//   "nodeOne": "4",
//   "nodeTwo": "1",
//   "nodeThree": "2"
// }
// Test Case 6
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1},
//       {"id": "2", "left": "1", "right": "3", "value": 2},
//       {"id": "3", "left": null, "right": "4", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4}
//     ],
//     "root": "2"
//   },
//   "nodeOne": "1",
//   "nodeTwo": "2",
//   "nodeThree": "3"
// }
// Test Case 7
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "3", "left": "2", "right": null, "value": 3},
//       {"id": "4", "left": "3", "right": "5", "value": 4},
//       {"id": "5", "left": null, "right": "7", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": "6", "right": null, "value": 7},
//       {"id": "8", "left": "4", "right": "10", "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9},
//       {"id": "10", "left": "9", "right": "14", "value": 10},
//       {"id": "11", "left": null, "right": null, "value": 11},
//       {"id": "12", "left": "11", "right": "13", "value": 12},
//       {"id": "13", "left": null, "right": null, "value": 13},
//       {"id": "14", "left": "12", "right": null, "value": 14}
//     ],
//     "root": "8"
//   },
//   "nodeOne": "2",
//   "nodeTwo": "4",
//   "nodeThree": "13"
// }
// Test Case 8
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "3", "left": "2", "right": null, "value": 3},
//       {"id": "4", "left": "3", "right": null, "value": 4},
//       {"id": "5", "left": "4", "right": null, "value": 5},
//       {"id": "6", "left": "5", "right": null, "value": 6},
//       {"id": "7", "left": "6", "right": null, "value": 7},
//       {"id": "8", "left": "7", "right": "9", "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9}
//     ],
//     "root": "8"
//   },
//   "nodeOne": "8",
//   "nodeTwo": "7",
//   "nodeThree": "1"
// }
// Test Case 9
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "3", "left": "2", "right": null, "value": 3}
//     ],
//     "root": "3"
//   },
//   "nodeOne": "2",
//   "nodeTwo": "1",
//   "nodeThree": "3"
// }
// Test Case 10
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "3", "left": "2", "right": null, "value": 3}
//     ],
//     "root": "3"
//   },
//   "nodeOne": "1",
//   "nodeTwo": "2",
//   "nodeThree": "3"
// }
// Test Case 11
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": null, "value": 1},
//       {"id": "2", "left": "1", "right": "3", "value": 2},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "4", "left": "2", "right": "5", "value": 4},
//       {"id": "5", "left": null, "right": null, "value": 5},
//       {"id": "6", "left": "4", "right": "8", "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": "7", "right": "9", "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9}
//     ],
//     "root": "6"
//   },
//   "nodeOne": "9",
//   "nodeTwo": "8",
//   "nodeThree": "6"
// }
// Test Case 12
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": "2", "value": 1},
//       {"id": "2", "left": null, "right": null, "value": 2},
//       {"id": "3", "left": "1", "right": "4", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": "3", "right": null, "value": 5},
//       {"id": "6", "left": "5", "right": "8", "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": "7", "right": "9", "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9},
//       {"id": "10", "left": "6", "right": "15", "value": 10},
//       {"id": "11", "left": null, "right": "12", "value": 11},
//       {"id": "12", "left": null, "right": null, "value": 12},
//       {"id": "13", "left": "11", "right": null, "value": 13},
//       {"id": "14", "left": "13", "right": null, "value": 14},
//       {"id": "15", "left": "14", "right": "18", "value": 15},
//       {"id": "16", "left": null, "right": null, "value": 16},
//       {"id": "17", "left": "16", "right": null, "value": 17},
//       {"id": "18", "left": "17", "right": null, "value": 18}
//     ],
//     "root": "10"
//   },
//   "nodeOne": "12",
//   "nodeTwo": "13",
//   "nodeThree": "15"
// }
// Test Case 13
// {
//   "tree": {
//     "nodes": [
//       {"id": "1", "left": null, "right": "2", "value": 1},
//       {"id": "2", "left": null, "right": null, "value": 2},
//       {"id": "3", "left": "1", "right": "4", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": "3", "right": null, "value": 5},
//       {"id": "6", "left": "5", "right": "8", "value": 6},
//       {"id": "7", "left": null, "right": null, "value": 7},
//       {"id": "8", "left": "7", "right": "9", "value": 8},
//       {"id": "9", "left": null, "right": null, "value": 9},
//       {"id": "10", "left": "6", "right": "15", "value": 10},
//       {"id": "11", "left": null, "right": "12", "value": 11},
//       {"id": "12", "left": null, "right": null, "value": 12},
//       {"id": "13", "left": "11", "right": null, "value": 13},
//       {"id": "14", "left": "13", "right": null, "value": 14},
//       {"id": "15", "left": "14", "right": "18", "value": 15},
//       {"id": "16", "left": null, "right": null, "value": 16},
//       {"id": "17", "left": "16", "right": null, "value": 17},
//       {"id": "18", "left": "17", "right": null, "value": 18}
//     ],
//     "root": "10"
//   },
//   "nodeOne": "5",
//   "nodeTwo": "10",
//   "nodeThree": "15"
// }
// Test Case 14
// {
//   "tree": {
//     "nodes": [
//       {"id": "0", "left": null, "right": null, "value": 0},
//       {"id": "1", "left": "0", "right": null, "value": 1},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "3", "left": "2", "right": "4", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": "3", "right": "7", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": "6", "right": "8", "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "5"
//   },
//   "nodeOne": "5",
//   "nodeTwo": "3",
//   "nodeThree": "4"
// }
// Test Case 15
// {
//   "tree": {
//     "nodes": [
//       {"id": "0", "left": null, "right": null, "value": 0},
//       {"id": "1", "left": "0", "right": null, "value": 1},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "3", "left": "2", "right": "4", "value": 3},
//       {"id": "4", "left": null, "right": null, "value": 4},
//       {"id": "5", "left": "3", "right": "7", "value": 5},
//       {"id": "6", "left": null, "right": null, "value": 6},
//       {"id": "7", "left": "6", "right": "8", "value": 7},
//       {"id": "8", "left": null, "right": null, "value": 8}
//     ],
//     "root": "5"
//   },
//   "nodeOne": "5",
//   "nodeTwo": "3",
//   "nodeThree": "1"
// }
