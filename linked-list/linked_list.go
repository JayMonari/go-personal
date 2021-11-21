package linkedlist

import "errors"

// Node is what lists are made of
type Node struct {
	Val        interface{}
	prev, next *Node
}

// List is a structure of double linked list
type List struct{ head, tail *Node }

// ErrEmptyList occurs when the list is empty
var ErrEmptyList = errors.New("list is empty")

// Next returns the next node in the chain
func (e *Node) Next() *Node { return e.next }

// Prev returns the previous node in the chain
func (e *Node) Prev() *Node { return e.prev }

// NewList generates a new list with the given node values
func NewList(args ...interface{}) *List {
	l := &List{}
	for _, arg := range args {
		l.PushBack(arg)
	}
	return l
}

// PushFront adds values to the front of the list
func (l *List) PushFront(v interface{}) {
	n := &Node{Val: v}
	if l.head == nil {
		l.head, l.tail = n, n
	} else {
		n.next = l.head
		l.head.prev = n
		l.head = n
	}
}

// PushBack adds values to the end of the list
func (l *List) PushBack(v interface{}) {
	n := &Node{Val: v}
	if l.head == nil {
		l.head, l.tail = n, n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
}

// PopFront retrieves value from the beginning of the list
func (l *List) PopFront() (interface{}, error) {
	if l == nil || l.head == nil {
		return nil, ErrEmptyList
	}

	n, val := l.head, l.head.Val
	l.head = n.next
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	n = nil

	return val, nil
}

// PopBack retrieves value from the end of the list
func (l *List) PopBack() (interface{}, error) {
	if l == nil || l.tail == nil {
		return nil, ErrEmptyList
	}

	n, val := l.tail, l.tail.Val
	l.tail = n.prev
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	n = nil

	return val, nil
}

// Reverse reverses the list elements
func (l *List) Reverse() *List {
	for n := l.head; n != nil; n = n.prev {
		n.next, n.prev = n.prev, n.next
	}
	l.head, l.tail = l.tail, l.head
	return l
}

// First gives the first node in the list
func (l *List) First() *Node { return l.head }

// Last gives the last node in the list
func (l *List) Last() *Node { return l.tail }
