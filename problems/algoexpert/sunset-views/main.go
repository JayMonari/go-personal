package main

const (
	w = "WEST"
	e = "EAST"
)

func SunsetViews(buildings []int, direction string) []int {
	idxs := []int{}
	idx, step := len(buildings)-1, -1
	if direction == w {
		idx, step = 0, 1
	}

	for max := 0; idx >= 0 && idx < len(buildings); idx += step {
		if b := buildings[idx]; b > max {
			idxs = append(idxs, idx)
			max = b
		}
	}

	if direction == e {
		for i, j := 0, len(idxs)-1; i < j; i, j = i+1, j-1 {
			idxs[i], idxs[j] = idxs[j], idxs[i]
		}
	}
	return idxs
}

func SunsetViewsStack(buildings []int, direction string) []int {
	idx, step := len(buildings)-1, -1
	if direction == e {
		idx, step = 0, 1
	}

	stack := []int{}
	for ; idx >= 0 && idx < len(buildings); idx += step {
		for len(stack) > 0 && buildings[stack[len(stack)-1]] <= buildings[idx] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, idx)
	}

	if direction == w {
		for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
			stack[i], stack[j] = stack[j], stack[i]
		}
	}
	return stack
}

// Test Case 1
// {
//   "buildings": [3, 5, 4, 4, 3, 1, 3, 2],
//   "direction": "EAST"
// }
// Test Case 2
// {
//   "buildings": [3, 5, 4, 4, 3, 1, 3, 2],
//   "direction": "WEST"
// }
// Test Case 3
// {
//   "buildings": [10, 11],
//   "direction": "EAST"
// }
// Test Case 4
// {
//   "buildings": [2, 4],
//   "direction": "WEST"
// }
// Test Case 5
// {
//   "buildings": [1],
//   "direction": "EAST"
// }
// Test Case 6
// {
//   "buildings": [1],
//   "direction": "WEST"
// }
// Test Case 7
// {
//   "buildings": [],
//   "direction": "EAST"
// }
// Test Case 8
// {
//   "buildings": [],
//   "direction": "WEST"
// }
// Test Case 9
// {
//   "buildings": [7, 1, 7, 8, 9, 8, 7, 6, 5, 4, 2, 5],
//   "direction": "EAST"
// }
// Test Case 10
// {
//   "buildings": [1, 2, 3, 4, 5, 6],
//   "direction": "EAST"
// }
// Test Case 11
// {
//   "buildings": [1, 2, 3, 4, 5, 6],
//   "direction": "WEST"
// }
// Test Case 12
// {
//   "buildings": [1, 2, 3, 1, 5, 6, 9, 1, 9, 9, 11, 10, 9, 12, 8],
//   "direction": "WEST"
// }
// Test Case 13
// {
//   "buildings": [20, 2, 3, 1, 5, 6, 9, 1, 9, 9, 11, 10, 9, 12, 8],
//   "direction": "EAST"
// }
