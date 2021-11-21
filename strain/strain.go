package strain

type Ints []int
type Lists [][]int
type Strings []string

type IntPredicate func(int) bool
type ListPredicate func([]int) bool
type StringPredicate func(string) bool

func (i Ints) Keep(pred IntPredicate) (o Ints) {
	for _, v := range i {
		if pred(v) {
			o = append(o, v)
		}
	}
	return
}

func (i Ints) Discard(pred IntPredicate) (o Ints) {
	for _, v := range i {
		if !pred(v) {
			o = append(o, v)
		}
	}
	return
}

func (l Lists) Keep(pred ListPredicate) (o Lists) {
	for _, v := range l {
		if pred(v) {
			o = append(o, v)
		}
	}
	return
}

func (s Strings) Keep(pred StringPredicate) (o Strings) {
	for _, v := range s {
		if pred(v) {
			o = append(o, v)
		}
	}
	return
}
