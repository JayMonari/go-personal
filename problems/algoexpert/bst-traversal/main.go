package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (t BST) InOrderTraverse(nums []int) []int {
	if t.Left != nil {
		nums = t.Left.InOrderTraverse(nums)
	}
	nums = append(nums, t.Value)
	if t.Right != nil {
		nums = t.Right.InOrderTraverse(nums)
	}
	return nums
}

func (t BST) PreOrderTraverse(nums []int) []int {
	nums = append(nums, t.Value)
	if t.Left != nil {
		nums = t.Left.PreOrderTraverse(nums)
	}
	if t.Right != nil {
		nums = t.Right.PreOrderTraverse(nums)
	}
	return nums
}

func (t BST) PostOrderTraverse(nums []int) []int {
	if t.Left != nil {
		nums = t.Left.PostOrderTraverse(nums)
	}
	if t.Right != nil {
		nums = t.Right.PostOrderTraverse(nums)
	}
	nums = append(nums, t.Value)
	return nums
}

// Test Case 1
// {
//   "tree": {
//     "nodes": [
//       {"id": "10", "left": "5", "right": "15", "value": 10},
//       {"id": "15", "left": null, "right": "22", "value": 15},
//       {"id": "22", "left": null, "right": null, "value": 22},
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
//       {"id": "100", "left": "5", "right": "502", "value": 100},
//       {"id": "502", "left": "204", "right": "55000", "value": 502},
//       {"id": "55000", "left": "1001", "right": null, "value": 55000},
//       {"id": "1001", "left": null, "right": "4500", "value": 1001},
//       {"id": "4500", "left": null, "right": null, "value": 4500},
//       {"id": "204", "left": "203", "right": "205", "value": 204},
//       {"id": "205", "left": null, "right": "207", "value": 205},
//       {"id": "207", "left": "206", "right": "208", "value": 207},
//       {"id": "208", "left": null, "right": null, "value": 208},
//       {"id": "206", "left": null, "right": null, "value": 206},
//       {"id": "203", "left": null, "right": null, "value": 203},
//       {"id": "5", "left": "2", "right": "15", "value": 5},
//       {"id": "15", "left": "5-2", "right": "22", "value": 15},
//       {"id": "22", "left": null, "right": "57", "value": 22},
//       {"id": "57", "left": null, "right": "60", "value": 57},
//       {"id": "60", "left": null, "right": null, "value": 60},
//       {"id": "5-2", "left": null, "right": null, "value": 5},
//       {"id": "2", "left": "1", "right": "3", "value": 2},
//       {"id": "3", "left": null, "right": null, "value": 3},
//       {"id": "1", "left": "-51", "right": "1-2", "value": 1},
//       {"id": "1-2", "left": null, "right": "1-3", "value": 1},
//       {"id": "1-3", "left": null, "right": "1-4", "value": 1},
//       {"id": "1-4", "left": null, "right": "1-5", "value": 1},
//       {"id": "1-5", "left": null, "right": null, "value": 1},
//       {"id": "-51", "left": "-403", "right": null, "value": -51},
//       {"id": "-403", "left": null, "right": null, "value": -403}
//     ],
//     "root": "100"
//   }
// }
// Test Case 3
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
// Test Case 4
// {
//   "tree": {
//     "nodes": [
//       {"id": "10", "left": null, "right": null, "value": 10}
//     ],
//     "root": "10"
//   }
// }
// Test Case 5
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
// Test Case 6
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
