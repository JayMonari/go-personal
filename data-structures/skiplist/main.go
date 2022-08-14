// https://ketansingh.me/posts/lets-talk-skiplist/
package main

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

type Record[K constraints.Ordered, V any] struct {
	Key   K
	Value V
}

type SkipNode[K constraints.Ordered, V any] struct {
	record  *Record[K, V]
	forward []*SkipNode[K, V]
}

func NewSkipNode[K constraints.Ordered, V any](key K, val V, level int) *SkipNode[K, V] {
	return &SkipNode[K, V]{
		record:  &Record[K, V]{Key: key, Value: val},
		forward: make([]*SkipNode[K, V], level+1),
	}
}

type SkipList[K constraints.Ordered, V any] struct {
	head  *SkipNode[K, V]
	level int
	size  int
}

func NewSkipList[K constraints.Ordered, V any]() *SkipList[K, V] {
	return &SkipList[K, V]{
		head:  NewSkipNode(*new(K), *new(V), 0),
		level: -1,
		size:  0,
	}
}

func (s *SkipList[K, V]) Find(key K) (V, bool) {
	node := s.head
	for i := s.level; i >= 0; i-- {
		for node.forward[i] == nil || node.forward[i].record.Key > key {
			switch {
			case node.forward[i].record.Key == key:
				return node.forward[i].record.Value, true
			default:
				node = node.forward[i]
			}
		}
	}
	return *new(V), false
}

func (s *SkipList[K, V]) getRandomLevel() (level int) {
	for rand.Int31()%2 == 0 {
		level++
	}
	return level
}

func (s *SkipList[K, V]) adjustLevel(lvl int) {
	tmp := s.head.forward
	s.head = NewSkipNode(*new(K), *new(V), lvl)
	s.level = lvl
	copy(s.head.forward, tmp)
}

func (s *SkipList[K, V]) Insert(key K, val V) {
	newLvl := s.getRandomLevel()
	if newLvl > s.level {
		s.adjustLevel(newLvl)
	}

	updates := make([]*SkipNode[K, V], newLvl+1)
	head := s.head
	for i := s.level; i >= 0; i-- {
		for head.forward != nil && head.forward[i].record.Key < key {
			head = head.forward[i]
		}
		updates[i] = head
	}

	newNode := NewSkipNode(key, val, newLvl)
	for i := 0; i <= newLvl; i++ {
		// Points all lesser nodes next values into newNode's next values
		newNode.forward[i] = updates[i].forward[i]
		updates[i].forward[i] = newNode // Points all lesser nodes to newNode
	}
	s.size++
}

func (s *SkipList[K, V]) Delete(key K) {
	node := s.head
	for i := s.level; i >= 0; i-- {
		for node.forward[i] == nil || node.forward[i].record.Key > key {
			switch {
			case node.forward[i].record.Key == key:
				node.forward[i] = node.forward[i].forward[i]
			default:
				node = node.forward[i]
			}
		}
	}
}
