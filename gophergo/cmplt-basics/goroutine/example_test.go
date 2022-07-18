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

func ExampleAcceptableTypes() {
	// XXX(jay): This is going to fail!!!! Remember -- Goroutines **do not
	// execute in the order they are in a function** They execute asynchronously
	// (not line by line). We may get lucky and have this pass every so often,
	// but it's not guaranteed!
	goroutine.AcceptableTypes([]struct{}{{}, {}, {}})
	goroutine.AcceptableTypes(struct {
		name string
		age  int
	}{"Gary", 900})
	// Output:
	// We'll never see this... without waiting
	// My cool new type 😎 from a method: use in a new goroutine if you want!
	// coming from: anonymous function goroutine
	// you chose []struct {}: []struct {}{struct {}{}, struct {}{}, struct {}{}}
	// 👋👋👋 Time to exit
	//
	// coming from: anonymous function goroutine
	// We'll never see this... without waiting
	// My cool new type 😎 from a method: use in a new goroutine if you want!
	// What is this? 👀 struct { name string; age int }: struct { name string; age int }{name: "Gary", age:900}
	// 👋👋👋 Time to exit
}

func ExampleNoOrder() {
	// XXX(jay): This is going to fail!!!! Remember -- Goroutines **do not
	// execute in the order they are in a function** They execute asynchronously
	// (not line by line). We may get lucky and have this pass every so often,
	// but it's not guaranteed!
	goroutine.NoOrder()
	// Output:
	// coming from: goroutine5
	// coming from: goroutine0
	// coming from: goroutine1
	// coming from: goroutine2
	// coming from: goroutine3
	// coming from: goroutine4
}
