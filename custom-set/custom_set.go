package stringset

import (
	"fmt"
	"strings"
)

// Set is a collection of unique strings.
type Set map[string]bool

// New creates a new Set.
func New() Set { return Set{} }

// NewFromSlice creates a new set from a slice of strings.
func NewFromSlice(sl []string) Set {
	set := Set{}
	for _, s := range sl {
		set.Add(s)
	}
	return set
}

// String returns the Set formatted as: {"a", "b", "c"}
func (s Set) String() string {
	vals := make([]string, 0, len(s))
	for k := range s {
		vals = append(vals, fmt.Sprintf("\"%s\"", k))
	}
	return fmt.Sprintf("{%s}", strings.Join(vals, ", "))
}

// IsEmpty
func (s *Set) IsEmpty() bool { return len(*s) == 0 }

// Has returns whether the set contains a value in constant time.
func (s *Set) Has(str string) bool {
	_, ok := (*s)[str]
	return ok
}

// Add adds a unique string to the Set.
func (s *Set) Add(str string) { (*s)[str] = true }

// Subset returns whether all values of s1 are in s2, if s1 is empty true is
// always returned.
func Subset(s1, s2 Set) bool {
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

// Disjoint returns whether s1 and s2 share any elements, if both Sets are
// empty true is always returned.
func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

// Equal returns whether all elements of s1 and s2 are found in both Sets.
func Equal(s1, s2 Set) bool { return len(s1) == len(s2) && Subset(s1, s2) }

// Intersection returns a new Set of the values shared by s1 and s2.
func Intersection(s1, s2 Set) Set {
	inter := Set{}
	for k := range s1 {
		if s2.Has(k) {
			inter.Add(k)
		}
	}
	return inter
}

// Difference returns a new Set of the values that are only in s1.
func Difference(s1, s2 Set) Set {
	dif := Set{}
	for k := range s1 {
		if !s2.Has(k) {
			dif.Add(k)
		}
	}
	return dif
}

// Union returns a new Set of all elements in s1 and s2.
func Union(s1, s2 Set) Set {
	u := make(Set, len(s1) + len(s1))
	for _, s := range []Set{s1, s2} {
		for k := range s {
			u.Add(k)
		}
	}
	return u
}
