package main

import "math"

func DijkstrasAlgorithm(start int, edges [][][]int) []int {
	foundDists := make([]int, len(edges))
	for i := range edges {
		foundDists[i] = math.MaxInt32
	}
	foundDists[start] = 0

	pairs := make([]item, len(edges))
	for i := range edges {
		pairs[i] = item{i, math.MaxInt32}
	}
	minDists := NewMinHeap(pairs)
	minDists.Update(start, 0)

	for !minDists.IsEmpty() {
		node, currMinDist := minDists.Remove()
		if currMinDist == math.MaxInt32 {
			break
		}

		for _, e := range edges[node] {
			dest, dist := e[0], e[1]
			newPathDistance := currMinDist + dist
			if newPathDistance < foundDists[dest] {
				foundDists[dest] = newPathDistance
				minDists.Update(dest, newPathDistance)
			}
		}
	}

	res := make([]int, len(edges))
	for i, d := range foundDists {
		if d == math.MaxInt32 {
			res[i] = -1
		} else {
			res[i] = d
		}
	}
	return res
}

type item struct{ node, distance int }

type MinHeap struct {
	items []item
	nodes map[int]int
}

func NewMinHeap(items []item) *MinHeap {
	nodes := map[int]int{}
	for _, item := range items {
		nodes[item.node] = item.node
	}
	heap := &MinHeap{items: items, nodes: nodes}
	heap.buildHeap()
	return heap
}

func (h *MinHeap) IsEmpty() bool { return h.length() == 0 }

func (h *MinHeap) Remove() (int, int) {
	l := h.length()
	h.swap(0, l-1)
	peeked := h.items[l-1]
	h.items = h.items[0 : l-1]
	delete(h.nodes, peeked.node)
	h.siftDown(0, l-2)
	return peeked.node, peeked.distance
}

func (h *MinHeap) Update(vertex int, value int) {
	h.items[h.nodes[vertex]] = item{vertex, value}
	h.siftUp(h.nodes[vertex])
}

func (h MinHeap) swap(i, j int) {
	h.nodes[h.items[i].node] = j
	h.nodes[h.items[j].node] = i
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h MinHeap) length() int { return len(h.items) }

func (h *MinHeap) buildHeap() {
	first := (len(h.items) - 2) / 2
	for i := first + 1; i >= 0; i-- {
		h.siftDown(i, len(h.items)-1)
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
		if c2Idx > -1 && h.items[c2Idx].distance < h.items[c1Idx].distance {
			toSwap = c2Idx
		}

		if h.items[toSwap].distance >= h.items[currIdx].distance {
			return
		}
		h.swap(currIdx, toSwap)
		currIdx = toSwap
		c1Idx = currIdx*2 + 1
	}
}

func (h *MinHeap) siftUp(currIdx int) {
	pIdx := (currIdx - 1) / 2
	for currIdx > 0 && h.items[currIdx].distance < h.items[pIdx].distance {
		h.swap(currIdx, pIdx)
		currIdx = pIdx
		pIdx = (currIdx - 1) / 2
	}
}

// Test Case 1
// {
//   "start": 0,
//   "edges": [
//     [
//       [1, 7]
//     ],
//     [
//       [2, 6],
//       [3, 20],
//       [4, 3]
//     ],
//     [
//       [3, 14]
//     ],
//     [
//       [4, 2]
//     ],
//     [],
//     []
//   ]
// }
// Test Case 2
// {
//   "start": 1,
//   "edges": [
//     [],
//     [],
//     [],
//     []
//   ]
// }
// Test Case 3
// {
//   "start": 7,
//   "edges": [
//     [
//       [1, 1],
//       [3, 1]
//     ],
//     [
//       [2, 1]
//     ],
//     [
//       [6, 1]
//     ],
//     [
//       [1, 3],
//       [2, 4],
//       [4, 2],
//       [5, 3],
//       [6, 5]
//     ],
//     [
//       [5, 1]
//     ],
//     [
//       [4, 1]
//     ],
//     [
//       [5, 2]
//     ],
//     [
//       [0, 7]
//     ]
//   ]
// }
// Test Case 4
// {
//   "start": 4,
//   "edges": [
//     [
//       [1, 3],
//       [2, 2]
//     ],
//     [
//       [3, 7]
//     ],
//     [
//       [1, 2],
//       [3, 4],
//       [4, 1]
//     ],
//     [],
//     [
//       [0, 2],
//       [1, 8],
//       [3, 1]
//     ]
//   ]
// }
// Test Case 5
// {
//   "start": 1,
//   "edges": [
//     [
//       [1, 2]
//     ],
//     [
//       [0, 1]
//     ],
//     [
//       [3, 1]
//     ],
//     [
//       [2, 2]
//     ]
//   ]
// }
// Test Case 6
// {
//   "start": 0,
//   "edges": [
//     [
//       [1, 1],
//       [7, 8]
//     ],
//     [
//       [2, 1]
//     ],
//     [
//       [3, 1]
//     ],
//     [
//       [4, 1]
//     ],
//     [
//       [5, 1]
//     ],
//     [
//       [6, 1]
//     ],
//     [
//       [7, 1]
//     ],
//     []
//   ]
// }
// Test Case 7
// {
//   "start": 3,
//   "edges": [
//     [
//       [1, 2],
//       [3, 3],
//       [4, 2]
//     ],
//     [
//       [0, 1],
//       [6, 3]
//     ],
//     [
//       [3, 9]
//     ],
//     [
//       [0, 3],
//       [1, 4],
//       [4, 4],
//       [8, 7]
//     ],
//     [
//       [0, 1],
//       [10, 3]
//     ],
//     [
//       [7, 1],
//       [8, 4]
//     ],
//     [
//       [8, 1]
//     ],
//     [],
//     [
//       [7, 1]
//     ],
//     [
//       [10, 2]
//     ],
//     []
//   ]
// }
// Test Case 8
// {
//   "start": 8,
//   "edges": [
//     [
//       [1, 4],
//       [7, 11]
//     ],
//     [
//       [0, 4],
//       [2, 11],
//       [7, 14]
//     ],
//     [
//       [1, 11],
//       [3, 10],
//       [5, 7],
//       [8, 5]
//     ],
//     [
//       [2, 10],
//       [4, 12],
//       [5, 17]
//     ],
//     [
//       [3, 12],
//       [5, 13],
//       [6, 3]
//     ],
//     [
//       [2, 7],
//       [3, 17],
//       [4, 13],
//       [6, 5]
//     ],
//     [
//       [4, 3],
//       [5, 6],
//       [7, 4],
//       [9, 8]
//     ],
//     [
//       [0, 11],
//       [1, 14],
//       [6, 4],
//       [8, 10]
//     ],
//     [
//       [2, 5],
//       [6, 9],
//       [7, 10]
//     ],
//     []
//   ]
// }
// Test Case 9
// {
//   "start": 3,
//   "edges": [
//     [
//       [2, 4]
//     ],
//     [
//       [0, 2]
//     ],
//     [
//       [1, 1],
//       [3, 2]
//     ],
//     [
//       [0, 3]
//     ]
//   ]
// }
