package main

import (
	"fmt"
	"time"
)

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
