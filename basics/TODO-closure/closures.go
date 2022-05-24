package closure

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"
	"time"
)

// ClosureAsGenerator is an example of the idea of closures. The state of the
// function is sealed away (closed away) and it keeps that state even after the
// scope is destroyed. Think of a clam with a grain of sand, turning into a
// pearl.
func ClosureAsGenerator() func() int {
	startNum := 0
	return func() int {
		startNum++
		return startNum
	}
}

// ChangeValues is used for showing off examples with closures and accessing
// the values of ChangeValues from within the closures.
type ChangeValues struct {
	MyStr     string
	IsChanged bool
}

// ClosureAsMiddleware is an example of not changing the existing function, but
// only adding new features to it. In this example we can make a logger. It
// will check what the values passed in are before and after the function call
// without interfering with the original function! Pretty neat 💯
func ClosureAsMiddleware(myFunc func(cv *ChangeValues, n *int)) func(cv *ChangeValues, n *int) {
	return func(cv *ChangeValues, n *int) {
		fmt.Printf("this is a statement that happens **BEFORE** myFunc:\nHere are the values before changing them: %+v and %d\n", cv, *n)
		myFunc(cv, n)
		fmt.Printf("this is a statement that happens **AFTER** myFunc:\nHere are the values after changing them: %+v and %d\n", cv, *n)
	}
}

// ClosureForAccessingData is an example of feeding values to a function that
// does not accept that type or more values. This is very common with the
// `http.HandlerFunc` which is required to have exactly 2 parameters
// `http.ResponseWriter` and `*http.Request`. So how do we make a
// `http.HandlerFunc` have more parameters? Closures of course!
func ClosureForAccessingData(cv ChangeValues) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, cv.MyStr)
	}
}

// ClosureForStandardLibrary is an example of when the standard library will
// ask you to provide a closure in order to complete the functions parameters.
// In this we see two examples: `strings.Map` and `sort.Search`
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

// ClosureAvoidCallbackHell is an example of how, in Go, there is no such thing
// as "Callback Hell" because you can always call work synchronously and if you
// want to do it asynchronously you would use a goroutine with `go` to do it.
// If you are unfamiliar with "Callback Hell" it's not important to learn about
// it. You won't experience it in Go because it's much more well designed!
func ClosureAvoidCallbackHell() {
	// Simulate long enough time for main go routine to go to next line in test
	time.Sleep(3 * time.Millisecond)
	result1 := DoWork1(1, 2, 3, 4, 5, 6, 7, 8)
	result2 := DoWork2(result1)
	result3 := DoWork3(result2)
	fmt.Printf("Some padding for final result:\n%14s", result3)
}

// DoWork1 is a global private closure. It creates the sum of several numbers.
var DoWork1 = func(args ...int) int {
	sum := 0
	for _, n := range args {
		sum += n
	}
	return sum
}

// DoWork2 acts like it's doing some heavy number crunching and using the value
// that it obtains from DoWork1.
var DoWork2 = func(sum int) float32 {
	crunchNums := math.Pi * 2
	return float32(sum) * float32(crunchNums)
}

// DoWork3 changes the float32 data type into a string with 4 places of
// precision.
var DoWork3 = func(floaty float32) string {
	return fmt.Sprintf("%0.4f", floaty)
}
