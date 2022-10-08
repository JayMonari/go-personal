package main

func SortKSortedArray(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nums
	}
	slc4heap := make([]int, min(k+1, len(nums)))
	copy(slc4heap, nums[0:min(k+1, len(nums))])
	kMinHeap := NewMinHeap(slc4heap)

	toInsertIdx := 0
	for i := k + 1; i < len(nums); i++ {
		nums[toInsertIdx] = kMinHeap.Remove()
		toInsertIdx++

		kMinHeap.Insert(nums[i])
	}

	for !kMinHeap.IsEmpty() {
		nums[toInsertIdx] = kMinHeap.Remove()
		toInsertIdx++
	}
	return nums
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type MinHeap []int

func NewMinHeap(nums []int) *MinHeap {
	heap := MinHeap(nums)
	ptr := &heap
	ptr.BuildHeap(nums)
	return ptr
}

func (h *MinHeap) BuildHeap(array []int) {
	first := (len(array) - 2) / 2
	for i := first + 1; i >= 0; i-- {
		h.siftDown(i, len(array)-1)
	}
}

func (h *MinHeap) siftDown(currIdx, endIdx int) {
	c1Idx := currIdx*2 + 1
	for c1Idx <= endIdx {
		c2Idx := -1
		if currIdx*2+2 <= endIdx {
			c2Idx = currIdx*2 + 2
		}
		toSwap := c1Idx
		if c2Idx > -1 && (*h)[c2Idx] < (*h)[c1Idx] {
			toSwap = c2Idx
		}
		if (*h)[toSwap] < (*h)[currIdx] {
			(*h)[currIdx], (*h)[toSwap] = (*h)[toSwap], (*h)[currIdx]
			currIdx = toSwap
			c1Idx = currIdx*2 + 1
		} else {
			return
		}
	}
}

func (h *MinHeap) siftUp() {
	i := len(*h) - 1
	pIdx := (i - 1) / 2
	for i > 0 {
		if (*h)[i] >= (*h)[pIdx] {
			return
		}
		(*h)[i], (*h)[pIdx] = (*h)[pIdx], (*h)[i]
		i = pIdx
		pIdx = (i - 1) / 2
	}
}

func (h MinHeap) Peek() int {
	if len(h) == 0 {
		return -1
	}
	return h[0]
}

func (h *MinHeap) Remove() int {
	l := len(*h)
	(*h)[0], (*h)[l-1] = (*h)[l-1], (*h)[0]
	peeked := (*h)[l-1]
	*h = (*h)[0 : l-1]
	h.siftDown(0, l-2)
	return peeked
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	h.siftUp()
}

func (h *MinHeap) IsEmpty() bool { return len(*h) == 0 }

// Test Case 1
// {
//   "array": [3, 2, 1, 5, 4, 7, 6, 5],
//   "k": 3
// }
// Test Case 2
// {
//   "array": [-1, -3, -4, 2, 1, 3],
//   "k": 2
// }
// Test Case 3
// {
//   "array": [1, 2, 3, 4, 5],
//   "k": 0
// }
// Test Case 4
// {
//   "array": [],
//   "k": 5
// }
// Test Case 5
// {
//   "array": [4, 3, 2, 1, 2, 5, 6],
//   "k": 4
// }
// Test Case 6
// {
//   "array": [3, 2, 1, 0, 4, 7, 6, 5, 9, 8, 7],
//   "k": 3
// }
// Test Case 7
// {
//   "array": [2, 1, 4, 3, 5, 6, 8, 7],
//   "k": 1
// }
// Test Case 8
// {
//   "array": [1, 0, 1, 1, 1, 1, 1, 1],
//   "k": 1
// }
// Test Case 9
// {
//   "array": [5, 4, 3, 2, -100],
//   "k": 5
// }
// Test Case 10
// {
//   "array": [3, 3, 2, 1, 6, 4, 4, 5, 9, 7, 8, 11, 12],
//   "k": 3
// }
// Test Case 11
// {
//   "array": [1],
//   "k": 1
// }
// Test Case 12
// {
//   "array": [-1, -5],
//   "k": 1
// }
// Test Case 13
// {
//   "array": [-2, -3, 1, 2, 3, 1, 1, 2, 3, 8, 100, 130, 9, 12],
//   "k": 4
// }
// Test Case 14
// {
//   "array": [1, 2, 3, 4, 5, 6, 1],
//   "k": 8
// }
