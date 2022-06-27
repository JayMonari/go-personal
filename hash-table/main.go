package main

// ArraySize is the size of the hash table array
const ArraySize = 7

// HashTable is a data structure that can insert, search, and delete items in
// constant time.
type HashTable struct {
	array [ArraySize]*bucket
}

// New initializes a HashTable to have buckets.
func New() (ht HashTable) {
	for i := range ht.array {
		ht.array[i] = &bucket{}
	}
	return ht
}

// Insert places a key into the HashTable in constant time.
func (h *HashTable) Insert(key string) {
	h.array[hash(key)].insert(key)
}

// Search looks through the HashTable for the key and returns whether it was
// found or not.
func (h *HashTable) Search(key string) bool {
	return h.array[hash(key)].search(key)
}

// Delete will remove a key from the HashTable. It returns true if the key was
// found and successfully removed, false otherwise.
func (h *HashTable) Delete(key string) bool {
	return h.array[hash(key)].delete(key)
}

// bucket is a linked list in each slot of a HashTable.
type bucket struct {
	head *node
}

// insert places a key into the bucket in constant time.
func (b *bucket) insert(k string) {
	newNode := &node{key: k}
	newNode.next, b.head = b.head, newNode
}

// search looks through the bucket for the key and returns whether it was found
// or not.
func (b *bucket) search(key string) bool {
	currNode := b.head
	for currNode != nil && currNode.key != key {
		currNode = currNode.next
	}
	return currNode != nil
}

// delete will remove a key from the bucket.
func (b *bucket) delete(key string) bool {
	if b.head == nil {
		return false
	}
	if b.head.key == key {
		b.head = b.head.next
		return true
	}

	prevNode := b.head
	for currNode := prevNode.next; currNode != nil; currNode = currNode.next {
		if currNode.key == key {
			currNode.next = currNode.next.next
			return true
		}
	}
	return false
}

// node ...
type node struct {
	key  string
	next *node
}

func main() {
}

func hash(key string) int {
	sum := 0
	for _, r := range key {
		sum += int(r)
	}
	return sum % ArraySize
}
