package goroutine_test

import "basics/goroutine"

func ExampleWillNotWait() {
	goroutine.WillNotWait()
	// Output:
}

func ExampleMakeItWait() {
	goroutine.MakeItWait()
	// Output:
	// We'll never see this...
}
