package main

import (
	"container/heap"
	"fmt"
	"stdlib/container"
)

// This example creates a PriorityQueue with some items, adds and manipulates
// an item, and then removes the items in Priority order.
func main() {
	items := map[string]int{"banana": 3, "apple": 2, "pear": 4}
	pq := make(container.PriorityQueue, len(items))
	i := 0
	for v, p := range items {
		pq[i] = &container.Item{
			Value:    v,
			Priority: p,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	item := &container.Item{Value: "orange", Priority: 1}
	heap.Push(&pq, item)
	pq.Update(item, item.Value, 5)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*container.Item)
    fmt.Printf("Priority: %02d Fruit: %s\n", item.Priority, item.Value)
	}
	// h := &container.IntHeap{2, 1, 5, 9, 1, 3, 4}
	// heap.Init(h)
	// fmt.Println("min:", (*h)[0])
	// for h.Len() > 0 {
	// 	fmt.Println(heap.Pop(h))
	// }
}
