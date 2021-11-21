package wordy

type collection interface {
	createIterator() iterator
}

type IdxOperPair [2]int

type iterator interface {
	hasNext() bool
	next() IdxOperPair
}

type IdxOperCollection []IdxOperPair

func (i IdxOperCollection) createIterator() iterator {
	return &operIterator{ops: i}
}

type operIterator struct {
	index int
	ops   IdxOperCollection
}

// hasNext
func (o *operIterator) hasNext() bool {
	if o.index < len(o.ops) {
		return true
	}
	return false
}

// next hands back the next array
func (o *operIterator) next() IdxOperPair {
	if o.hasNext() {
		n := o.ops[o.index]
		o.index++
		return n
	}
	panic("How to array correctly?")
}
