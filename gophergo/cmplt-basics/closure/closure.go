package closure

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"
)

// AsGenerator is an example of the idea of closures. The state of the
// function is sealed away (closed away) and it keeps that state even after the
// scope is destroyed. Think of an oyster 🦪 with a grain of sand, turning into
// a pearl. ⚪
func AsGenerator() func() int {
	startNum := 0
	return func() int {
		startNum++
		return startNum
	}
}

// ForAccessingData shows we can access variables that do not belong to the
// inner function aka our closure.
func ForAccessingData() {
	notInClosure := "How does it know I exist?"
	func() {
		fmt.Println("This is a function, that is allowed to see local variables 👉",
			notInClosure)
	}()
	// NOTE(jay): This is just like how we can grab a global slice and change the
	// innards of it. -- `GlobalSlice[0] = "Some other value"` Because the
	// `GlobalSlice` is in scope of the function we can change it. Essentially
	// the "global" area has expanded for our closure allowing us access to the
	// `notInClosure` variable.
}

// NOTE(jay): This cannot be done outside of the 👆 above function. Closures
// are given special permission to access variables in the current scope.
// func doesNotWork() {
// 	fmt.Println("This is a function, that is allowed to see local variables 👉",
// 		notInClosureScope)
// }

// MyS is used for showing off examples with closures and accessing
// the values of MyS from within the closures.
type MyS struct {
	MyStr     string
	IsChanged bool
}

// AsMiddleware is an example of not changing the existing function, but
// only adding new features to it. In this example we can make a logger. It
// will check what the values passed in are before and after the function call
// without interfering with the original function! Pretty neat 💯
func AsMiddleware(myFunc func(strct *MyS, n *int)) func(strct *MyS, n *int) {
	return func(strct *MyS, n *int) {
		fmt.Printf("this is a statement that happens **BEFORE** myFunc:\n"+
			"Here are the values before changing them: %+v and %d\n", strct, *n)
		myFunc(strct, n)
		fmt.Printf("this is a statement that happens **AFTER** myFunc:\n"+
			"Here are the values after changing them: %+v and %d\n", strct, *n)
	}
}

// ForAccessingMoreData is an example of feeding values to a function that
// does not accept that type or more values. This is very common with the
// `http.HandlerFunc` which is required to have exactly 2 parameters
// `http.ResponseWriter` and `*http.Request`. So how do we make a
// `http.HandlerFunc` have more parameters? Closures of course!
func ForAccessingMoreData(strct MyS) func(http.ResponseWriter, *http.Request) {
	// This func here 👇matches with this func here👆
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, strct.MyStr)
	}
}

// ForStandardLibrary is an example of when the standard library will
// ask you to provide a closure in order to complete the functions parameters.
// In this we see two examples: `strings.Map` and `sort.Search`
func ForStandardLibrary() {
	rot13 := "pybfherf ner pbby"
	// NOTE(jay): Even though this is an anonymous function it's not a true
	// closure as we don't use any variables from outside of the inner functions
	// scope. We **could** do that and it would become a closure, but it is
	// important to notice we don't have to make an anonymous function a closure.
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
	fmt.Printf("indices found from Binary Search of 13:\nAfter: %d\n"+
		"Before: %d\nEqual: %d", indexAfter13, indexBefore13, indexOf13)
}

// AvoidCallbackHell is an example of how, in Go, there is no such thing
// as "Callback Hell" because you can always call work synchronously and if you
// want to do it asynchronously you would use a goroutine with `go` to do it.
// If you are unfamiliar with "Callback Hell" it's not important to learn about
// it. You won't experience it in Go because it's much more well designed!
func AvoidCallbackHell() {
	// result1 is DoWork1 as a closure
	result1 := func(args ...int) int {
		sum := 0
		for _, n := range args {
			sum += n
		}
		return sum
	}(1, 2, 3, 4, 5, 6, 7, 8)

	// result2 is DoWork2 as a closure
	result2 := func(sum int) float32 {
		crunchNums := math.Pi * 2
		return float32(sum) * float32(crunchNums)
	}(result1)

	// result3 is DoWork3 as a closure
	result3 := func(f float32) string {
		return fmt.Sprintf("%0.4f", f)
	}(result2)

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
var DoWork3 = func(f float32) string {
	return fmt.Sprintf("%0.4f", f)
}

// Gotcha shows that closures will grab from outside of their environment, but
// closures have no control of those variables having their values changed.
func Gotcha() {
	limit := 4
	funcs := make([]func(), limit)
	for gotcha := 0; gotcha < limit; gotcha++ {
		funcs[gotcha] = func() {
			fmt.Printf("number: %d and pointer: %p\n", gotcha, &gotcha)
		}
	}
	for i := 0; i < limit; i++ {
		funcs[i]()
	}
}

// GotchaFix shows a deeper understanding of closures and how we can avoid the
// very common beginner mistake of
func GotchaFix() {
	limit := 4
	funcs := make([]func(), limit)
	for gotcha := 0; gotcha < limit; gotcha++ {
		// NOTE(jay): this is known as shadowing. This is confusing and is normally
		// avoided if possible. The explanation is we are creating a new variable
		// scoped inside of the for loop and therefore this `gotcha` won't change
		// only the `gotcha` on the for loop scope will change.
		//
		// This also works and you could put gotcha2 in place of gotcha in the
		// closure if it helps your understanding.
		//
		// gotcha2 := gotcha
		gotcha := gotcha

		// 👇 This also works and it's less confusing for the uninitiated. Again,
		// we are creating a new variable `gotcha2` inside of the for loop, that is
		// the same as the `gotcha` on the for loop scope.
		// gotcha2 := gotcha
		funcs[gotcha] = func() {
			fmt.Printf("number: %d and pointer: %p\n", gotcha, &gotcha)
		}
	}

	// Instead of using a closure we can use an anonymous function which has a
	// parameter passed in.
	funcs2 := make([]func(int), limit)
	for gotcha := 0; gotcha < limit; gotcha++ {
		funcs2[gotcha] = func(i int) {
			fmt.Printf("number: %d and pointer: %p\n", i, &i)
		}
	}

	fmt.Println("fix 1:")
	for i := 0; i < limit; i++ {
		funcs[i]()
	}
	fmt.Println("fix 2:")
	for i := 0; i < limit; i++ {
		funcs2[i](i)
	}
}
