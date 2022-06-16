package goroutine

import (
	"fmt"
	"time"
)

// https://go.dev/talks/2012/concurrency.slide#53

func WillNotWait() {
	// NOTE(jay): This will be seen if we run `go test` for `ExampleWillNotWait`
	// it just won't be a part of the main goroutines output because it exits.
	go toofast()
}

func toofast() { fmt.Println("We'll never see this... without waiting") }

// SwitchToOther shows us how to artificially allow the goroutine we spawn to
// finish and exit, by slowing down the main goroutine. This is **not** how
// it's done in go. We use channels for true concurrency, but this is important
// to see before we introduce channels.
func SwitchToOther() {
	go toofast()
	// Make it wait 8 milliseconds to see separate goroutines output.
	time.Sleep(8 * time.Millisecond)
}

func main() {
	// This will always be first in the output because it is a blocking call.
	processData("main goroutine")
	// The `go` keyword starts a new goroutine. Here we start three new
	// goroutines, but if you run this over and over we'll see that they are
	// always change their order!
	go processData("new goroutine1")
	go processData("new goroutine2")
	go processData("new goroutine3")

	// The `go` keyword needs a function and that is all, even if it is an inline
	// anonymous function, it can still be used in a goroutine
	go func(comingFrom string) {
		fmt.Println("coming from:", comingFrom)
	}("anonymous inline goroutine")

	// We don't have an example_test.go file because goroutines are
	// non-deterministic by default, meaning the output for this function is
	// different everytime! 🤔 Can you think of a way to make this deterministic?
	go processData("new goroutine4")
	go processData("new goroutine5")
	go processData("new goroutine6")

	// We have to wait, because the main goroutine will shutdown other goroutines
	// and exit immediately. Comment out this line and the "time" package import
	// above and see what you get!
	time.Sleep(time.Second)
	fmt.Println("exiting main")
}

func processData(comingFrom string) {
	fmt.Println("coming from:", comingFrom)
}
