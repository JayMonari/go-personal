package hamming

import (
	"errors"
	"reflect"
)

// Distance returns the Hamming difference between two strands of DNA
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("the two strands must be the same length")
	}

	count := 0
	zippedStrands := [][2]rune{}
	genericZip([]rune(a), []rune(b), &zippedStrands)
	for _, strand := range zippedStrands {
		if strand[0] != strand[1] {
			count++
		}
	}
	return count, nil
}

type nucTuple struct {
	a, b rune
}

// zip takes in two rune slices and forms one slice with a tuple of the
// elements in their original order paired together. Slower than the
// genericZip somehow....
func zip(a, b []rune) ([]nucTuple, error) {
	if len(a) != len(b) {
		return nil, errors.New("zip: arguments must be of same length")
	}

	r := make([]nucTuple, len(a))
	for i, e := range a {
		r[i] = nucTuple{e, b[i]}
	}
	return r, nil
}

// zip takes in two slices of the same type and same length and converts them
// into a single slice with combined elements for each index.
// e.g. []int{1,2,3} + []int{4,5,6} -> []int{{1,4},{2,5},{3,6}}
func genericZip(a, b, c interface{}) error {
	ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)
	if ta.Kind() != reflect.Slice || tb.Kind() != reflect.Slice || ta != tb {
		return errors.New("zip: first two arguments must be slices of the same type")
	}
	if tc.Kind() != reflect.Ptr {
		return errors.New("zip: third argument must be pointer to slice")
	}
	for tc.Kind() == reflect.Ptr {
		tc = tc.Elem()
	}
	if tc.Kind() != reflect.Slice {
		return errors.New("zip: third argument must be pointer to slice")
	}

	eta, _, etc := ta.Elem(), tb.Elem(), tc.Elem()
	if etc.Kind() != reflect.Array || etc.Len() != 2 {
		return errors.New("zip: third argument's elements must be an array of length 2")
	}
	if etc.Elem() != eta {
		return errors.New("zip: third argument's elements must be an array of elements of the same type that the first two arguments are slices of")
	}

	va, vb, vc := reflect.ValueOf(a), reflect.ValueOf(b), reflect.ValueOf(c)
	for vc.Kind() == reflect.Ptr {
		vc = vc.Elem()
	}
	if va.Len() != vb.Len() {
		return errors.New("zip: first two arguments must have same length")
	}

	for i := 0; i < va.Len(); i++ {
		ea, eb := va.Index(i), vb.Index(i)
		tt := reflect.New(etc).Elem()
		tt.Index(0).Set(ea)
		tt.Index(1).Set(eb)
		vc.Set(reflect.Append(vc, tt))
	}
	return nil
}
