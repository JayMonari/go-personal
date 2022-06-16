package goroutine_test

import "basics/goroutine"

func ExampleWillNotWait() {
	goroutine.WillNotWait()
	// Output:
}

func ExampleSwitchToOther() {
	goroutine.SwitchToOther()
	// Output:
	// We'll never see this... without waiting
}

func ExampleAnonymousFunctions() {
	// XXX(jay): This is going to fail!!!! Remember -- Goroutines **do not
	// execute in the order they are in a function** They execute asynchronously.
	// We may get lucky and have this pass every so often, but it's not
	// guaranteed!
	goroutine.AnonymousFunctions(1)
	goroutine.AnonymousFunctions([]struct{}{{}, {}, {}})
	goroutine.AnonymousFunctions("Look a string!")
	goroutine.AnonymousFunctions(struct {
		name string
		age  int
	}{"Gary", 900})
	// Output:
	// you chose int: 1
	// coming from: anonymous function goroutine
	// 👋👋👋 Time to exit
	// you chose []struct {}: []struct {}{struct {}{}, struct {}{}, struct {}{}}
	// coming from: anonymous function goroutine
	// 👋👋👋 Time to exit
	// you chose string: Look a string!
	// coming from: anonymous function goroutine
	// 👋👋👋 Time to exit
	// what is this 👀 struct { name string; age int }: struct { name string; age int }{name:"Gary", age:900}
	// coming from: anonymous function goroutine
	// 👋👋👋 Time to exit
}
