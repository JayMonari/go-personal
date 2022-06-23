package variable_test

import "basics/variable"

func ExampleDeclareVarExplicit() {
	variable.DeclareVarExplicit()
	// Output:
	// My Name's Jay! 😁,
	// From 0 to 1000
	// Given bool value: true
}

func ExampleDeclareVarImplicit() {
	variable.DeclareVarImplicit()
	// Output:
	// Does anyone have any room for 🥧,
	// More Pi 3.141590
	// Given bool value: false
	//
}

func ExampleDeclareVarZero() {
	variable.DeclareVarZero()
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
