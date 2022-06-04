package closure_test

import (
	"basics/closure"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func ExampleAsGenerator() {
	countUp := closure.AsGenerator()
	fmt.Print(countUp(), countUp(), countUp(), countUp(), countUp(), countUp())
	startOver := closure.AsGenerator()
	fmt.Println()
	fmt.Println("New number generator starting over:", startOver())
	fmt.Print("Old number generator still going: ", countUp())
	// Output:
	// 1 2 3 4 5 6
	// New number generator starting over: 1
	// Old number generator still going: 7
}

func ExampleForAccessingData() {
	closure.ForAccessingData()
	// Output:
	// This is a function, that is allowed to see local variables 👉 How does it know I exist?
}

func ExampleAsMiddleware() {
	changesStuff := func(strct *closure.MyS, n *int) {
		strct.MyStr = "A Whole New String from a closure! 🤯"
		strct.IsChanged = true
		*n = 9999
	}
	logsStuff := closure.AsMiddleware(changesStuff)
	strct := &closure.MyS{MyStr: "My cool String 😎", IsChanged: false}
	intPtr := new(int)
	logsStuff(strct, intPtr)
	// Output:
	// this is a statement that happens **BEFORE** myFunc:
	// Here are the values before changing them: &{MyStr:My cool String 😎 IsChanged:false} and 0
	// this is a statement that happens **AFTER** myFunc:
	// Here are the values after changing them: &{MyStr:A Whole New String from a closure! 🤯 IsChanged:true} and 9999
}

func ExampleForAccessingMoreData() {
	// Create a handler that now has access to outside resources **without**
	// adding additonal parameters to the original function. Super Cool! 😁
	handler := closure.ForAccessingMoreData(closure.MyS{
		MyStr: "Accessed this data from an input to function without closure" +
			" knowing where it came from!"})

	// This is how to test HTTP Handlers, it will be covered much later
	rr := httptest.NewRecorder()
	http.HandlerFunc(handler).ServeHTTP(rr, nil)
	// What we want to see coming from ResponseRecorder aka our Response
	fmt.Println(rr.Body.String())
	// Output:
	//  Accessed this data from an input to function without closure knowing where it came from!
}

func ExampleForStandardLibrary() {
	closure.ForStandardLibrary()
	// Output:
	// closures are cool
	// indices found from Binary Search of 13:
	// After: 14
	// Before: 0
	// Equal: 13
}

func ExampleAvoidCallbackHell() {
	// `go` keyword creates a new goroutine it will be covered in a later lesson.
	// Just know it let's us multi-task or better put -- introduce concurrency.
	go func() {
		result1 := closure.DoWork1(1, 2, 3, 4, 5, 6, 7, 8)
		result2 := closure.DoWork2(result1)
		result3 := closure.DoWork3(result2)
		fmt.Printf("\nSome padding for final result:\n%14s", result3)
	}()
	go closure.AvoidCallbackHell()

	fmt.Println("Done asynchronously! If you're not FIRST you're LAST")
	// Simulate long enough time for goroutine to finish above.
	time.Sleep(3 * time.Millisecond)
	// Output:
	// Done asynchronously! If you're not FIRST you're LAST
	// Some padding for final result:
	//       226.1947
	// Some padding for final result:
	//       226.1947
}
