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
