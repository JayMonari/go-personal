package closure_test

import (
	"basics/closure"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func ExampleClosureAsGenerator() {
	countUp := closure.ClosureAsGenerator()
	fmt.Print(countUp(), countUp(), countUp(), countUp(), countUp(), countUp())
	startOver := closure.ClosureAsGenerator()
	fmt.Println()
	fmt.Print("New number generator starting over: ", startOver())
	// Output:
	// 1 2 3 4 5 6
	// New number generator starting over: 1
}

func ExampleClosureAsMiddleware() {
	changesStuff := func(cv *closure.ChangeValues, n *int) {
		cv.MyStr = "A Whole New String from a closure! 🤯"
		cv.IsChanged = true
		*n = 9999
	}
	logsStuff := closure.ClosureAsMiddleware(changesStuff)
	cv := &closure.ChangeValues{MyStr: "My cool String 😎", IsChanged: false}
	intPtr := new(int)
	logsStuff(cv, intPtr)
	// Output:
	// this is a statement that happens **BEFORE** myFunc:
	// Here are the values before changing them: &{MyStr:My cool String 😎 IsChanged:false} and 0
	// this is a statement that happens **AFTER** myFunc:
	// Here are the values after changing them: &{MyStr:A Whole New String from a closure! 🤯 IsChanged:true} and 9999
}

func ExampleClosureForAccessingData() {
	// This is how to test HTTP Handlers, it will be covered much later
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	// Create a handler that now has access to outside resources **without**
	// adding additonal parameters to the original function. Super Cool! 😁
	handler := closure.ClosureForAccessingData(closure.ChangeValues{MyStr: "Accessed this data from an input to function without closure knowing where it came from!"})
	http.HandlerFunc(handler).ServeHTTP(rr, req)
	// What we want to see.
	fmt.Println(rr.Body.String())
	// Output:
	//  Accessed this data from an input to function without closure knowing where it came from!
}

func ExampleClosureForStandardLibrary() {
	closure.ClosureForStandardLibrary()
	// Output:
	// closures are cool
	// indices found from Binary Search of 13:
	// After: 14
	// Before: 0
	// Equal: 13
}

func ExampleClosureAvoidCallbackHell() {
	// `go` keyword creates a new goroutine it will be covered in a later lesson.
	// Just know it let's us multi-task.
	go func() {
		// Simulate long enough time for main go routine to go to next line in test
		time.Sleep(3 * time.Millisecond)
		result1 := closure.DoWork1(1, 2, 3, 4, 5, 6, 7, 8)
		result2 := closure.DoWork2(result1)
		result3 := closure.DoWork3(result2)
		fmt.Printf("Some padding for final result:\n%14s", result3)
	}()

	fmt.Println("Done asynchronously! If you're not FIRST you're LAST")
	time.Sleep(5 * time.Millisecond)
	// Output:
	// Done asynchronously! If you're not FIRST you're LAST
	// Some padding for final result:
	//       226.1947
}
