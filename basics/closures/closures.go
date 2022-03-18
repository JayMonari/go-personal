package closure

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"
)

// ClosureAsGenerator
func ClosureAsGenerator() func() int {
	startNum := 0
	return func() int {
		startNum++
		return startNum
	}
}

// ChangeValues
type ChangeValues struct {
	MyStr     string
	IsChanged bool
}

// ClosureAsMiddleware
func ClosureAsMiddleware(myFunc func(cv *ChangeValues, n *int)) func(cv *ChangeValues, n *int) {
	return func(cv *ChangeValues, n *int) {
		fmt.Printf("this is a statement that happens **BEFORE** myFunc:\nHere are the values before changing them: %+v and %d\n", cv, *n)
		myFunc(cv, n)
		fmt.Printf("this is a statement that happens **AFTER** myFunc:\nHere are the values after changing them: %+v and %d\n", cv, *n)
	}
}

// ClosureForAccessingData
func ClosureForAccessingData(cv ChangeValues) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, cv.MyStr)
	}
}

// ClosureForStandardLibrary
func ClosureForStandardLibrary() {
	rot13 := "pybfherf ner pbby"
	mappedStr := strings.Map(func(r rune) rune {
		if r == ' ' {
			return r
		}
		r -= 13
		// We might go outside of alphabet range a-z, so we correct for it here.
		switch {
		case r < 'a':
			return r + 'z' - 'a' + 1
		case r > 'z':
			return r%'z' + 'a' - 1
		default:
			return r
		}
	}, rot13)
	fmt.Println(mappedStr)

	sortedInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	indexAfter13 := sort.Search(len(sortedInts), func(i int) bool {
		return sortedInts[i] > 13
	})
	indexBefore13 := sort.Search(len(sortedInts), func(i int) bool {
		return sortedInts[i] < 13
	})
	indexOf13 := sort.Search(len(sortedInts), func(i int) bool {
		return sortedInts[i] == 13
	})
	fmt.Printf("indices found from Binary Search of 13:\nAfter: %d\nBefore: %d\nEqual: %d",
		indexAfter13, indexBefore13, indexOf13)
}

var doWork1 = func(args ...int) int {
	sum := 0
	for _, n := range args {
		sum += n
	}
	return sum
}

var doWork2 = func(sum int) float32 {
	crunchNums := math.Pi * 2
	return float32(sum) * float32(crunchNums)
}

var doWork3 = func(floaty float32) string {
	return fmt.Sprintf("%0.4f", floaty)
}

// ClosureAvoidCallbackHell
func ClosureAvoidCallbackHell() {
	result1 := doWork1(1, 2, 3, 4, 5, 6, 7, 8)
	result2 := doWork2(result1)
	result3 := doWork3(result2)
	fmt.Printf("Some padding for final result:\n%14s", result3)
}
