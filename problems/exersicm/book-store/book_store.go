package bookstore

var prices = map[int]int{
	1: 800,
	2: 1520,
	3: 2160,
	4: 2560,
	5: 3000,
}

func Cost(books []int) int {
	nSets := setLens(books)
	for contains(nSets, 3) && contains(nSets, 5) {
		nSets = remove(nSets, 3)
		nSets = remove(nSets, 5)
		nSets = append(nSets, 4, 4)
	}
	c := 0
	for _, n := range nSets {
		c += prices[n]
	}
	return c
}

func remove(s []int, v int) []int {
	for i, n := range s {
		if v == n {
			s[i] = s[len(s)-1]
			s = s[:len(s)-1]
			break
		}
	}
	return s
}

func contains(s []int, v int) bool {
	for _, n := range s {
		if v == n {
			return true
		}
	}
	return false
}

type set map[int]struct{}

func newSet(nums []int) set {
	s := make(map[int]struct{})
	for _, n := range nums {
		s[n] = struct{}{}
	}
	return s
}

func setLens(books []int) []int {
	co := make([]int, len(books))
	copy(co, books)
	nSets := make([]int, 0, 3)
	s := newSet(co)
	for len(co) != 0 {
		nSets = append(nSets, len(s))
		for k := range s {
			for i, n := range co {
				if k == n {
					co[i] = co[len(co)-1]
					co = co[:len(co)-1]
					break
				}
			}
		}
		s = newSet(co)
	}
	return nSets
}
