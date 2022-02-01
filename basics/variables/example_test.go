package main

func ExampleDeclareVarExplicit() {
	DeclareVarExplicit()
	// Output:
	// My Name's Jay! 😁,
	// From 0 to 1000
	// Given bool value: true
}

func ExampleDeclareVarImplicit() {
	DeclareVarImplicit()
	// Output:
	// Does anyone have any room for 🥧,
	// More Pi 3.141590
	// Given bool value: false
	//
}

func ExampleDeclareVarDefault() {
	DeclareVarDefault()
	// Output:
	// string default value is: ""
	// int default value is: 0
	// float32 default value is: 0.000000
	// bool default value is: false
	// rune default value is: 0 or "\x00"
}

func ExampleAssignmentOperator() {
	AssignmentOperator()
	// Output:
	// Short and Sweet. Very Nice! 👍
	// I really need c0ffee
	// Still works? true
}
