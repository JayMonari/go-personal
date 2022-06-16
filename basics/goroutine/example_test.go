package goroutine_test

import "basics/goroutine"

func ExampleWillNotWait() {
	goroutine.WillNotWait()
	// Output:
}

func ExampleSwitchToOther() {
	goroutine.SwitchToOther()
	// Output:
	// We'll never see this...
}
