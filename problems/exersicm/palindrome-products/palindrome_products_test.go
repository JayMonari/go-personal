package palindrome

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var testData = []struct {
	// input to Products(): range limits for factors of the palindrome
	fmin, fmax int
	// output from Products():
	pmin, pmax Product // min and max palandromic products
	errPrefix  string  // start of text if there is an error, "" otherwise
}{
	{1, 9,
		Product{}, // zero value means don't bother to test it
		Product{9, [][2]int{{1, 9}, {3, 3}}},
		""},
	{10, 99,
		Product{121, [][2]int{{11, 11}}},
		Product{9009, [][2]int{{91, 99}}},
		""},
	{100, 999,
		Product{10201, [][2]int{{101, 101}}},
		Product{906609, [][2]int{{913, 993}}},
		""},
	{4, 10, Product{}, Product{}, "no palindromes"},
	{10, 4, Product{}, Product{}, "fmin > fmax"},
	{-99, -10,
		Product{121, [][2]int{{-11, -11}}},
		Product{9009, [][2]int{{-99, -91}}},
		""},
	{-2, 2,
		Product{-4, [][2]int{{-2, 2}}},
		Product{4, [][2]int{{-2, -2}, {2, 2}}},
		""},
}

func TestPalindromeProducts(t *testing.T) {
	for _, test := range testData {
		ret := fmt.Sprintf("Products(%d, %d) returned",
			test.fmin, test.fmax)
		pmin, pmax, err := Products(test.fmin, test.fmax)
		var _ error = err
		switch {
		case err == nil:
			if test.errPrefix > "" {
				t.Fatalf(ret+" err = nil, want %q", test.errPrefix+"...")
			}
		case test.errPrefix == "":
			t.Fatalf(ret+" err = %q, want nil", err)
		case !strings.HasPrefix(err.Error(), test.errPrefix):
			t.Fatalf(ret+" err = %q, want %q", err, test.errPrefix+"...")
		default:
			continue
		}
		matchProd := func(ww string, rp, wp Product) {
			if len(wp.Factorizations) > 0 &&
				!reflect.DeepEqual(rp, wp) {
				t.Fatal(ret, ww, "=", rp, "want", wp)
			}
		}
		matchProd("pmin", pmin, test.pmin)
		matchProd("pmax", pmax, test.pmax)
	}
}

func BenchmarkPalindromeProducts(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testData {
			Products(test.fmin, test.fmax)
		}
	}
}
