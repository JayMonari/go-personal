package main

func KnapsackProblem(items [][]int, capacity int) []interface{} {
	vals := make([][]int, len(items)+1)
	for i := range vals {
		vals[i] = make([]int, capacity+1)
	}

	for i := 1; i < len(items)+1; i++ {
		currVal := items[i-1][0]
		currWeight := items[i-1][1]
		for c := 0; c < capacity+1; c++ {
			vals[i][c] = vals[i-1][c]
			if currWeight <= c {
				vals[i][c] = max(vals[i-1][c], vals[i-1][c-currWeight]+currVal)
			}
		}
	}
	return []interface{}{vals[len(items)][capacity], getKnapsackItems(vals, items)}
}

func getKnapsackItems(values [][]int, items [][]int) []int {
	var sequence []int
	i, c := len(values)-1, len(values[0])-1
	for i > 0 {
		if values[i][c] == values[i-1][c] {
			i--
		} else {
			sequence = append(sequence, i-1)
			c -= items[i-1][1]
			i--
		}
		if c == 0 {
			break
		}
	}

	for i, j := 0, len(sequence)-1; i < j; i, j = i+1, j-1 {
		sequence[i], sequence[j] = sequence[j], sequence[i]
	}
	return sequence
}

func max(first int, rest ...int) int {
	curr := first
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}

// Test Case 1
// {
//   "items": [
//     [1, 2],
//     [4, 3],
//     [5, 6],
//     [6, 7]
//   ],
//   "capacity": 10
// }
// Test Case 2
// {
//   "items": [
//     [1, 2],
//     [4, 3],
//     [5, 6],
//     [6, 9]
//   ],
//   "capacity": 11
// }
// Test Case 3
// {
//   "items": [
//     [465, 100],
//     [400, 85],
//     [255, 55],
//     [350, 45],
//     [650, 130],
//     [1000, 190],
//     [455, 100],
//     [100, 25],
//     [1200, 190],
//     [320, 65],
//     [750, 100],
//     [50, 45],
//     [550, 65],
//     [100, 50],
//     [600, 70],
//     [240, 40]
//   ],
//   "capacity": 200
// }
// Test Case 4
// {
//   "items": [
//     [465, 100],
//     [400, 85],
//     [255, 55],
//     [350, 45],
//     [650, 130],
//     [1000, 190],
//     [455, 100],
//     [100, 25],
//     [1200, 190],
//     [320, 65],
//     [750, 100],
//     [50, 45],
//     [550, 65],
//     [100, 50],
//     [600, 70],
//     [255, 40]
//   ],
//   "capacity": 200
// }
// Test Case 5
// {
//   "items": [
//     [2, 1],
//     [70, 70],
//     [30, 30],
//     [69, 69],
//     [100, 100]
//   ],
//   "capacity": 100
// }
// Test Case 6
// {
//   "items": [
//     [1, 2],
//     [70, 70],
//     [30, 30],
//     [69, 69],
//     [99, 100]
//   ],
//   "capacity": 100
// }
// Test Case 7
// {
//   "items": [
//     [1, 2],
//     [70, 70],
//     [30, 30],
//     [69, 69],
//     [100, 100]
//   ],
//   "capacity": 0
// }
