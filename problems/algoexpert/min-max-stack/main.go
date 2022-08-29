package main

type entry struct{ min, max int }

// MinMaxStack records the minimum value in the entire stack and maximum value
// along with what you expect a stack to be able to do.
type MinMaxStack struct {
	stack   []int
	entries []entry
}

// Peek shows the number on the top of the stack without removing it.
func (s *MinMaxStack) Peek() int { return s.stack[len(s.stack)-1] }

// Pop returns the top number from the stack and removes it from the stack.
func (s *MinMaxStack) Pop() int {
	s.entries = s.entries[:len(s.entries)-1]
	out := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return out
}

// Push places a number onto the stack.
func (s *MinMaxStack) Push(n int) {
	e := entry{min: n, max: n}
	if len(s.entries) > 0 {
		last := s.entries[len(s.entries)-1]
		e.min = min(last.min, n)
		e.max = max(last.max, n)
	}
	s.entries = append(s.entries, e)
	s.stack = append(s.stack, n)
}

// GetMin ...
func (s *MinMaxStack) GetMin() int { return s.entries[len(s.entries)-1].min }

// GetMax ...
func (s *MinMaxStack) GetMax() int { return s.entries[len(s.entries)-1].max }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test Case 1
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [7],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [2],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     }
//   ]
// }
// Test Case 2
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [2],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [7],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [1],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [8],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [3],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [9],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     }
//   ]
// }
// Test Case 3
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [5],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [5],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [5],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [8],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [8],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [0],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [8],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [9],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [5],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     }
//   ]
// }
// Test Case 4
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [2],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [0],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [5],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [4],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [4],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [11],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [-11],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     },
//     {
//       "arguments": [6],
//       "method": "push"
//     },
//     {
//       "arguments": [],
//       "method": "getMin"
//     },
//     {
//       "arguments": [],
//       "method": "getMax"
//     },
//     {
//       "arguments": [],
//       "method": "peek"
//     },
//     {
//       "arguments": [],
//       "method": "pop"
//     }
//   ]
// }
