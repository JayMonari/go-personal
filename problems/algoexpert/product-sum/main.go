package main

type SpecialArray []interface{}

func ProductSum(array []interface{}) int {
	return productSum(array, 1)
}

func productSum(vv []interface{}, multiplier int) int {
	sum := 0
	for _, v := range vv {
		switch t := v.(type) {
		case SpecialArray:
			sum += productSum(t, multiplier+1)
		case int:
			sum += t
		}
	}
	return sum * multiplier
}

// Test Case 1
// {
//   "array": [5, 2, [7, -1], 3, [6, [-13, 8], 4]]
// }
// Test Case 2
// {
//   "array": [1, 2, 3, 4, 5]
// }
// Test Case 3
// {
//   "array": [1, 2, [3], 4, 5]
// }
// Test Case 4
// {
//   "array": [
//     [1, 2],
//     3,
//     [4, 5]
//   ]
// }
// Test Case 5
// {
//   "array": [
//     [
//       [
//         [
//           [5]
//         ]
//       ]
//     ]
//   ]
// }
// Test Case 6
// {
//   "array": [9, [2, -3, 4], 1, [1, 1, [1, 1, 1]], [[[[3, 4, 1]]], 8], [1, 2, 3, 4, 5, [6, 7], -7], [1, [2, 3, [4, 5]], [6, 0, [7, 0, -8]], -7], [1, -3, 2, [1, -3, 2, [1, -3, 2], [1, -3, 2, [1, -3, 2]], [1, -3, 2]]], -3]
// }
