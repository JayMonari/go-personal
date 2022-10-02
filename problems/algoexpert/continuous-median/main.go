package main

type ContinuousMedianHandler struct {
	Median float64

	lowers, highers *Heap
}

func NewContinuousMedianHandler() *ContinuousMedianHandler {
	return &ContinuousMedianHandler{
		Median:  0,
		lowers:  NewHeap(func(a, b int) bool { return a > b }),
		highers: NewHeap(func(a, b int) bool { return a < b }),
	}
}

func (h *ContinuousMedianHandler) GetMedian() float64 {
	return h.Median
}

func (h *ContinuousMedianHandler) Insert(number int) {
	if h.lowers.Length() == 0 || number < h.lowers.Peek() {
		h.lowers.Insert(number)
	} else {
		h.highers.Insert(number)
	}
	h.rebalanceHeaps()
	h.updateMedian()
}

func (h *ContinuousMedianHandler) rebalanceHeaps() {
	if h.lowers.Length()-h.highers.Length() == 2 {
		h.highers.Insert(h.lowers.Remove())
	} else if h.highers.Length()-h.lowers.Length() == 2 {
		h.lowers.Insert(h.highers.Remove())
	}
}

func (h *ContinuousMedianHandler) updateMedian() {
	if h.lowers.Length() == h.highers.Length() {
		sum := (h.lowers.Peek() + h.highers.Peek())
		h.Median = float64(sum) / 2
	} else if h.lowers.Length() > h.highers.Length() {
		h.Median = float64(h.lowers.Peek())
	} else {
		h.Median = float64(h.highers.Peek())
	}
}

type Heap struct {
	comp   ComparisonFunc
	values []int
}

type ComparisonFunc func(int, int) bool

func NewHeap(fn ComparisonFunc) *Heap {
	return &Heap{
		comp:   fn,
		values: []int{},
	}
}

func (h *Heap) Length() int { return len(h.values) }

func (h *Heap) Peek() int {
	if len(h.values) == 0 {
		return -1
	}
	return h.values[0]
}

func (h *Heap) Insert(value int) {
	h.values = append(h.values, value)
	h.siftUp()
}

func (h *Heap) Remove() int {
	l := h.Length()
	h.values[0], h.values[l-1] = h.values[l-1], h.values[0]
	peeked := h.values[l-1]
	h.values = h.values[0 : l-1]
	h.siftDown()
	return peeked
}

func (h *Heap) siftUp() {
	i := h.Length() - 1
	parentIdx := (i - 1) / 2
	for i > 0 {
		current, parent := h.values[i], h.values[parentIdx]
		if h.comp(current, parent) {
			h.values[i], h.values[parentIdx] = h.values[parentIdx], h.values[i]
			i = parentIdx
			parentIdx = (i - 1) / 2
		} else {
			return
		}
	}
}

func (h *Heap) siftDown() {
	i := 0
	endIdx := h.Length() - 1
	c1Idx := i*2 + 1
	for c1Idx <= endIdx {
		c2Idx := -1
		if i*2+2 <= endIdx {
			c2Idx = i*2 + 2
		}
		toSwap := c1Idx
		if c2Idx > -1 && h.comp(h.values[c2Idx], h.values[c1Idx]) {
			toSwap = c2Idx
		}
		if h.comp(h.values[toSwap], h.values[i]) {
			h.values[i], h.values[toSwap] = h.values[toSwap], h.values[i]
			i = toSwap
			c1Idx = i*2 + 1
		} else {
			return
		}
	}
}

// Test Case 1
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     },
//     {
//       "arguments": [100],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     }
//   ]
// }
// Test Case 2
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     },
//     {
//       "arguments": [10],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     }
//   ]
// }
// Test Case 3
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "insert"
//     },
//     {
//       "arguments": [100],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     },
//     {
//       "arguments": [200],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     }
//   ]
// }
// Test Case 4
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "insert"
//     },
//     {
//       "arguments": [100],
//       "method": "insert"
//     },
//     {
//       "arguments": [200],
//       "method": "insert"
//     },
//     {
//       "arguments": [6],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     },
//     {
//       "arguments": [13],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     }
//   ]
// }
// Test Case 5
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "insert"
//     },
//     {
//       "arguments": [100],
//       "method": "insert"
//     },
//     {
//       "arguments": [200],
//       "method": "insert"
//     },
//     {
//       "arguments": [6],
//       "method": "insert"
//     },
//     {
//       "arguments": [13],
//       "method": "insert"
//     },
//     {
//       "arguments": [14],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     },
//     {
//       "arguments": [50],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     }
//   ]
// }
// Test Case 6
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "insert"
//     },
//     {
//       "arguments": [100],
//       "method": "insert"
//     },
//     {
//       "arguments": [200],
//       "method": "insert"
//     },
//     {
//       "arguments": [6],
//       "method": "insert"
//     },
//     {
//       "arguments": [13],
//       "method": "insert"
//     },
//     {
//       "arguments": [14],
//       "method": "insert"
//     },
//     {
//       "arguments": [50],
//       "method": "insert"
//     },
//     {
//       "arguments": [51],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     },
//     {
//       "arguments": [52],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     }
//   ]
// }
// Test Case 7
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "insert"
//     },
//     {
//       "arguments": [100],
//       "method": "insert"
//     },
//     {
//       "arguments": [200],
//       "method": "insert"
//     },
//     {
//       "arguments": [6],
//       "method": "insert"
//     },
//     {
//       "arguments": [13],
//       "method": "insert"
//     },
//     {
//       "arguments": [14],
//       "method": "insert"
//     },
//     {
//       "arguments": [50],
//       "method": "insert"
//     },
//     {
//       "arguments": [51],
//       "method": "insert"
//     },
//     {
//       "arguments": [52],
//       "method": "insert"
//     },
//     {
//       "arguments": [1000],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     },
//     {
//       "arguments": [10000],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     }
//   ]
// }
// Test Case 8
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "insert"
//     },
//     {
//       "arguments": [100],
//       "method": "insert"
//     },
//     {
//       "arguments": [200],
//       "method": "insert"
//     },
//     {
//       "arguments": [6],
//       "method": "insert"
//     },
//     {
//       "arguments": [13],
//       "method": "insert"
//     },
//     {
//       "arguments": [14],
//       "method": "insert"
//     },
//     {
//       "arguments": [50],
//       "method": "insert"
//     },
//     {
//       "arguments": [51],
//       "method": "insert"
//     },
//     {
//       "arguments": [52],
//       "method": "insert"
//     },
//     {
//       "arguments": [1000],
//       "method": "insert"
//     },
//     {
//       "arguments": [10000],
//       "method": "insert"
//     },
//     {
//       "arguments": [10001],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     },
//     {
//       "arguments": [10002],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     }
//   ]
// }
// Test Case 9
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "insert"
//     },
//     {
//       "arguments": [100],
//       "method": "insert"
//     },
//     {
//       "arguments": [200],
//       "method": "insert"
//     },
//     {
//       "arguments": [6],
//       "method": "insert"
//     },
//     {
//       "arguments": [13],
//       "method": "insert"
//     },
//     {
//       "arguments": [14],
//       "method": "insert"
//     },
//     {
//       "arguments": [50],
//       "method": "insert"
//     },
//     {
//       "arguments": [51],
//       "method": "insert"
//     },
//     {
//       "arguments": [52],
//       "method": "insert"
//     },
//     {
//       "arguments": [1000],
//       "method": "insert"
//     },
//     {
//       "arguments": [10000],
//       "method": "insert"
//     },
//     {
//       "arguments": [10001],
//       "method": "insert"
//     },
//     {
//       "arguments": [10002],
//       "method": "insert"
//     },
//     {
//       "arguments": [10003],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     },
//     {
//       "arguments": [10004],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     }
//   ]
// }
// Test Case 10
// {
//   "classMethodsToCall": [
//     {
//       "arguments": [5],
//       "method": "insert"
//     },
//     {
//       "arguments": [10],
//       "method": "insert"
//     },
//     {
//       "arguments": [100],
//       "method": "insert"
//     },
//     {
//       "arguments": [200],
//       "method": "insert"
//     },
//     {
//       "arguments": [6],
//       "method": "insert"
//     },
//     {
//       "arguments": [13],
//       "method": "insert"
//     },
//     {
//       "arguments": [14],
//       "method": "insert"
//     },
//     {
//       "arguments": [50],
//       "method": "insert"
//     },
//     {
//       "arguments": [51],
//       "method": "insert"
//     },
//     {
//       "arguments": [52],
//       "method": "insert"
//     },
//     {
//       "arguments": [1000],
//       "method": "insert"
//     },
//     {
//       "arguments": [10000],
//       "method": "insert"
//     },
//     {
//       "arguments": [10001],
//       "method": "insert"
//     },
//     {
//       "arguments": [10002],
//       "method": "insert"
//     },
//     {
//       "arguments": [10003],
//       "method": "insert"
//     },
//     {
//       "arguments": [10004],
//       "method": "insert"
//     },
//     {
//       "arguments": [75],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     },
//     {
//       "arguments": [80],
//       "method": "insert"
//     },
//     {
//       "arguments": [],
//       "method": "getMedian"
//     }
//   ]
// }
