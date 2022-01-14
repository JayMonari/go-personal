package main

import (
	"container/heap"
	"fmt"
	"stdlib/container"
)

func main() {
	h := &container.IntHeap{2, 1, 5, 9, 1, 3, 4}
	heap.Init(h)
	fmt.Println("min:", (*h)[0])
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
