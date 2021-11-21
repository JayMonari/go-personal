package linkedlist

import "fmt"

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(nums []int) *List {
	list := &List{}
	for _, n := range nums {
		list.Push(n)
	}
	return list
}

func (l *List) Size() int { return l.size }

func (l *List) Push(n int) {
	l.head = &Element{data: n, next: l.head}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size < 1 {
		return 0, fmt.Errorf("list is empty")
	}

	l.size--
	num, head := l.head.data, l.head
	l.head = head.next
	head.next = nil
	return num, nil
}

func (l *List) Array() []int {
	arr := make([]int, l.size)
	for i, n := len(arr)-1, l.head; n != nil; i, n = i-1, n.next {
		arr[i] = n.data
	}
	return arr
}

func (l *List) Reverse() *List {
	revList := &List{}
	for n := l.head; n != nil; n = n.next {
		revList.Push(n.data)
	}
	return revList
}
