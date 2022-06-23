package sublist

// Relation is the relationship between two slices of integer.
type Relation string

const (
	Eq    Relation = "equal"
	Sub   Relation = "sublist"
	Super Relation = "superlist"
	Uneq  Relation = "unequal"
)

// Sublist returns a Relation of two slices depending on what relationship they
// have.
func Sublist(slOne, slTwo []int) Relation {
	switch {
	case areEqual(slOne, slTwo):
		return Eq
	case isSub(slOne, slTwo):
		return Sub
	case isSub(slTwo, slOne):
		return Super
	default:
		return Uneq
	}
}

// isSub returns whether a is a sublist of b.
func isSub(a, b []int) bool {
	if len(a) >= len(b) {
		return false
	}

	for j := 0; j <= len(b) - len(a); j++ {
		isMatch := true
		for i, n := range a {
			if n != b[i+j] {
				isMatch = false
				break
			}
		}
		if isMatch {
			return isMatch
		}
	}
	return false
}

// areEqual returns whether two integer slices have equal length and the same
// ordered values.
func areEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, n := range a {
		if n != b[i] {
			return false
		}
	}
	return true
}
