package defers_test

import (
	"basics/defers"
	"fmt"
	"io"
	"os"
	"time"
)

func ExampleDeferRunAtEnd() {
	defers.DeferRunAtEnd()
	// Output:
	// This is the second line, but first to be printed.
	// A couple of more lines for good measure.
	// Never hurt anyone.
	// defer: this is the first line, but the last output.
}

func ExampleDeferLIFO() {
	defers.DeferLIFO()
	// Output:
	// Random line 1
	// Random line 2
	// Random line 3
	// Third In. First Out
	// Second In. Second Out
	// First In. Third Out
}

func ExampleDeferArgumentsEvaluated() {
	defers.DeferArgumentsEvaluated()
	// Output:
	// Number at end of function: 9001
	// Number in defer: 42
}

func ExampleDeferNamedReturn() {
	fmt.Println(defers.DeferNamedReturn())
	// Output: Bag of sand
}

func ExampleDeferReturn() {
	fmt.Println(defers.DeferReturn())
	// Output: Golden Idol
}

func ExampleDeferRecoverPanic() {
	// XXX: Uncomment this if you're a brave soul 🫣
	// defers.DeferRecoverPanic(false)
	defers.DeferRecoverPanic(true)
	// Output: Recovered from: WE'RE ALL GOING DOWN, THIS IS THE END!!!
}

// Honorable Mentions /////////////////////////////////////////////////////////

func ExampleDeferFileClose() {
	defers.DeferFileClose()
	f, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Reading contents of example.txt: ")
	io.Copy(os.Stdout, f)
	// Output:
	// closing
	// Reading contents of example.txt: 👋 🌏
}

func ExampleDeferTempFileRemoveClose() {
	defers.DeferTempFileRemoveClose()
	// Output:
	// closing
	// removing
}

func ExampleDeferHTTPBodyClose() {
	defers.DeferHTTPBodyClose()
	// Output: Does the first 4KiB of data have Gopher Go in it? true
}

func ExampleDeferCancelContext() {
	defers.DeferCancelContext(300 * time.Millisecond)
	defers.DeferCancelContext(5 * time.Millisecond)
	// Output:
	// You're too slow!!!
	// Don't forget to cancel!
}

func ExampleNewAccount() {
	a := defers.NewAccount(1000)
	a.Deposit(100)
	fmt.Println("Balance is:", a.Balance())
	// Output: Balance is: 1100
}
