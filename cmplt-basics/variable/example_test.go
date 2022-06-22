package variable_test

import "basics/variable"

func ExampleExplicit() {
	variable.Explicit()
	// Output:
	// My Name's Jay! 😁,
	// From 0 to 1000
	// Given bool value: true
}

func ExampleImplicit() {
	variable.Implicit()
	// Output:
	// Does anyone have any room for 🥧,
	// More Pi 3.141590
	// Given bool value: false
	//
}

func ExampleZero() {
	variable.Zero()
	// Output:
	// string zero value is: ""
	// int zero value is: 0
	// float32 zero value is: 0.000000
	// bool zero value is: false
	// rune zero value is: 0 or "\x00"
	// slice zero value is: []
	// map zero value is: map[]
}

func ExampleAssignmentOperator() {
	variable.AssignmentOperator()
	// Output:
	// Short and Sweet. Very Nice! 👍
	// I really need c0ffee
	// Still works? true
}
