package main

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (t BST) ValidateBst() bool {
	return t.valid(math.MinInt32, math.MaxInt32)
}

func (t BST) valid(min, max int) bool {
	if t.Value < min || t.Value >= max {
		return false
	}
	if t.Left != nil && !t.Left.valid(min, t.Value) {
		return false
	}
	if t.Right != nil && !t.Right.valid(t.Value, max) {
		return false
	}
	return true
}

// Test Case 1
// {
//   "tree": {
//     "nodes": [
//       {"id": "10", "left": "5", "right": "15", "value": 10},
//       {"id": "15", "left": "13", "right": "22", "value": 15},
//       {"id": "22", "left": null, "right": null, "value": 22},
//       {"id": "13", "left": null, "right": "14", "value": 13},
//       {"id": "14", "left": null, "right": null, "value": 14},
//       {"id": "5", "left": "2", "right": "5-2", "value": 5},
//       {"id": "5-2", "left": null, "right": null, "value": 5},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "10"
//   }
// }
// Test Case 2
// {
//   "tree": {
//     "nodes": [
//       {"id": "10", "left": "5", "right": "15", "value": 10},
//       {"id": "15", "left": null, "right": "22", "value": 15},
//       {"id": "22", "left": null, "right": null, "value": 22},
//       {"id": "5", "left": "2", "right": "5-2", "value": 5},
//       {"id": "5-2", "left": null, "right": null, "value": 5},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "1", "left": "-5", "right": null, "value": 1},
//       {"id": "-5", "left": "-15", "right": "-5-2", "value": -5},
//       {"id": "-5-2", "left": null, "right": "-2", "value": -5},
//       {"id": "-2", "left": null, "right": "-1", "value": -2},
//       {"id": "-1", "left": null, "right": null, "value": -1},
//       {"id": "-15", "left": "-22", "right": null, "value": -15},
//       {"id": "-22", "left": null, "right": null, "value": -22}
//     ],
//     "root": "10"
//   }
// }
// Test Case 3
// {
//   "tree": {
//     "nodes": [
//       {"id": "10", "left": null, "right": null, "value": 10}
//     ],
//     "root": "10"
//   }
// }
// Test Case 4
// {
//   "tree": {
//     "nodes": [
//       {"id": "10", "left": "5", "right": "15", "value": 10},
//       {"id": "15", "left": null, "right": "22", "value": 15},
//       {"id": "22", "left": null, "right": "500", "value": 22},
//       {"id": "500", "left": "50", "right": "1500", "value": 500},
//       {"id": "1500", "left": null, "right": "10000", "value": 1500},
//       {"id": "10000", "left": "2200", "right": null, "value": 10000},
//       {"id": "2200", "left": null, "right": null, "value": 2200},
//       {"id": "50", "left": null, "right": "200", "value": 50},
//       {"id": "200", "left": null, "right": null, "value": 200},
//       {"id": "5", "left": "2", "right": "5-2", "value": 5},
//       {"id": "5-2", "left": null, "right": null, "value": 5},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "10"
//   }
// }
// Test Case 5
// {
//   "tree": {
//     "nodes": [
//       {"id": "5000", "left": "5", "right": "55000", "value": 5000},
//       {"id": "55000", "left": null, "right": null, "value": 55000},
//       {"id": "5", "left": "2", "right": "15", "value": 5},
//       {"id": "15", "left": "5-2", "right": "22", "value": 15},
//       {"id": "22", "left": null, "right": "502", "value": 22},
//       {"id": "502", "left": "204", "right": null, "value": 502},
//       {"id": "204", "left": "203", "right": "205", "value": 204},
//       {"id": "205", "left": null, "right": "207", "value": 205},
//       {"id": "207", "left": "206", "right": "208", "value": 207},
//       {"id": "208", "left": null, "right": null, "value": 208},
//       {"id": "206", "left": null, "right": null, "value": 206},
//       {"id": "203", "left": null, "right": null, "value": 203},
//       {"id": "5-2", "left": null, "right": null, "value": 5},
//       {"id": "2", "left": "1", "right": "3", "value": 2},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "1", "left": null, "right": "1-2", "value": 1},
//       {"id": "1-2", "left": null, "right": "1-3", "value": 1},
//       {"id": "1-3", "left": null, "right": "1-4", "value": 1},
//       {"id": "1-4", "left": null, "right": "1-5", "value": 1},
//       {"id": "1-5", "left": null, "right": null, "value": 1}
//     ],
//     "root": "5000"
//   }
// }
// Test Case 6
// {
//   "tree": {
//     "nodes": [
//       {"id": "10", "left": "5", "right": "15", "value": 10},
//       {"id": "15", "left": null, "right": "22", "value": 15},
//       {"id": "22", "left": null, "right": null, "value": 22},
//       {"id": "5", "left": "2", "right": "5-2", "value": 5},
//       {"id": "5-2", "left": null, "right": "11", "value": 5},
//       {"id": "11", "left": null, "right": null, "value": 11},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "10"
//   }
// }
// Test Case 7
// {
//   "tree": {
//     "nodes": [
//       {"id": "10", "left": "5", "right": "15", "value": 10},
//       {"id": "15", "left": null, "right": "22", "value": 15},
//       {"id": "22", "left": null, "right": null, "value": 22},
//       {"id": "5", "left": "2", "right": "5-2", "value": 5},
//       {"id": "5-2", "left": null, "right": null, "value": 5},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "1", "left": "-5", "right": null, "value": 1},
//       {"id": "-5", "left": "-15", "right": "-5-2", "value": -5},
//       {"id": "-5-2", "left": null, "right": "-2", "value": -5},
//       {"id": "-2", "left": null, "right": "-1", "value": -2},
//       {"id": "-1", "left": null, "right": null, "value": -1},
//       {"id": "-15", "left": "-22", "right": null, "value": -15},
//       {"id": "-22", "left": "11", "right": null, "value": -22},
//       {"id": "11", "left": null, "right": null, "value": 11}
//     ],
//     "root": "10"
//   }
// }
// Test Case 8
// {
//   "tree": {
//     "nodes": [
//       {"id": "10", "left": "11", "right": "12", "value": 10},
//       {"id": "12", "left": null, "right": null, "value": 12},
//       {"id": "11", "left": null, "right": null, "value": 11}
//     ],
//     "root": "10"
//   }
// }
// Test Case 9
// {
//   "tree": {
//     "nodes": [
//       {"id": "10", "left": "5", "right": "15", "value": 10},
//       {"id": "15", "left": null, "right": "22", "value": 15},
//       {"id": "22", "left": null, "right": "500", "value": 22},
//       {"id": "500", "left": "50", "right": "1500", "value": 500},
//       {"id": "1500", "left": null, "right": "10000", "value": 1500},
//       {"id": "10000", "left": "2200", "right": "9999", "value": 10000},
//       {"id": "9999", "left": null, "right": null, "value": 9999},
//       {"id": "2200", "left": null, "right": null, "value": 2200},
//       {"id": "50", "left": null, "right": "200", "value": 50},
//       {"id": "200", "left": null, "right": null, "value": 200},
//       {"id": "5", "left": "2", "right": "5-2", "value": 5},
//       {"id": "5-2", "left": null, "right": null, "value": 5},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "10"
//   }
// }
// Test Case 10
// {
//   "tree": {
//     "nodes": [
//       {"id": "100", "left": "5", "right": "502", "value": 100},
//       {"id": "502", "left": "204", "right": "55000", "value": 502},
//       {"id": "55000", "left": null, "right": null, "value": 55000},
//       {"id": "204", "left": "203", "right": "205", "value": 204},
//       {"id": "205", "left": "300", "right": "207", "value": 205},
//       {"id": "207", "left": "206", "right": "208", "value": 207},
//       {"id": "208", "left": null, "right": null, "value": 208},
//       {"id": "206", "left": null, "right": null, "value": 206},
//       {"id": "300", "left": null, "right": null, "value": 300},
//       {"id": "203", "left": null, "right": null, "value": 203},
//       {"id": "5", "left": "2", "right": "15", "value": 5},
//       {"id": "15", "left": "5-2", "right": "22", "value": 15},
//       {"id": "22", "left": null, "right": null, "value": 22},
//       {"id": "5-2", "left": null, "right": null, "value": 5},
//       {"id": "2", "left": "1", "right": "3", "value": 2},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "1", "left": null, "right": "1-2", "value": 1},
//       {"id": "1-2", "left": null, "right": "1-3", "value": 1},
//       {"id": "1-3", "left": null, "right": "1-4", "value": 1},
//       {"id": "1-4", "left": null, "right": "1-5", "value": 1},
//       {"id": "1-5", "left": null, "right": null, "value": 1}
//     ],
//     "root": "100"
//   }
// }
// Test Case 11
// {
//   "tree": {
//     "nodes": [
//       {"id": "10", "left": "5", "right": "15", "value": 10},
//       {"id": "15", "left": null, "right": null, "value": 15},
//       {"id": "5", "left": null, "right": "10-2", "value": 5},
//       {"id": "10-2", "left": null, "right": null, "value": 10}
//     ],
//     "root": "10"
//   }
// }
// Test Case 12
// {
//   "tree": {
//     "nodes": [
//       {"id": "10", "left": "5", "right": "15", "value": 10},
//       {"id": "15", "left": "13", "right": "22", "value": 15},
//       {"id": "22", "left": null, "right": null, "value": 22},
//       {"id": "13", "left": null, "right": "16", "value": 13},
//       {"id": "16", "left": null, "right": null, "value": 16},
//       {"id": "5", "left": "2", "right": "5-2", "value": 5},
//       {"id": "5-2", "left": null, "right": null, "value": 5},
//       {"id": "2", "left": "1", "right": null, "value": 2},
//       {"id": "1", "left": null, "right": null, "value": 1}
//     ],
//     "root": "10"
//   }
// }
