package defers_test

import (
	"basics/defers"
	"fmt"
	"io"
	"os"
	"time"
)

func ExampleRunAtEnd() {
	defers.RunAtEnd()
	// Output:
	// This is the second line, but first to be printed.
	// A couple of more lines for good measure.
	// Never hurt anyone.
	// defer: this is the first line, but the last output.
}

func ExampleLIFO() {
	defers.LIFO()
	// Output:
	// Random line 1
	// Random line 2
	// Random line 3
	// Third In. First Out
	// Second In. Second Out
	// First In. Third Out
}

func ExampleArgumentsEvaluated() {
	defers.ArgumentsEvaluated()
	// Output:
	// Number at end of function: 9001
	// Number in defer: 42
}

func ExampleNamedReturn() {
	fmt.Println(defers.NamedReturn())
	// Output: Bag of sand
}

func ExampleReturn() {
	fmt.Println(defers.Return())
	// Output: Golden Idol
}

func ExampleRecoverPanic() {
	// NOTE(jay): Uncomment this if you're a brave soul 🫣
	// defers.RecoverPanic(false)
	defers.RecoverPanic(true)
	// Output: Recovered from: WE'RE ALL GOING DOWN, THIS IS THE END!!!
}

// Honorable Mentions /////////////////////////////////////////////////////////

func ExampleFileClose() {
	defers.FileClose()
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

func ExampleTempFileRemoveClose() {
	defers.TempFileRemoveClose()
	// Output:
	// closing
	// removing
}

func ExampleHTTPBodyClose() {
	defers.HTTPBodyClose()
	// Output: Does the first 4KiB of data have Gopher Go in it? true
}

func ExampleCancelContext() {
	defers.CancelContext(300 * time.Millisecond)
	defers.CancelContext(5 * time.Millisecond)
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
