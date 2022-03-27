package constants_test

import (
	"basics/constants"
	"fmt"
)

func ExampleStuck() {
	fmt.Println(constants.Stuck)
	// Output:
	// This variable can never be reassigned.
}

func ExampleHeartEyes() {
	// We can convert constants just like variables
	fmt.Println(string(constants.HeartEyes))
	// Output:
	// 😍
}

func ExampleAritmatic() {
	fmt.Println(constants.Arithmetic)
	// Output:
	// 175.38731365097925
}

func ExampleAlwaysTrue() {
	fmt.Println(constants.AlwaysTrue)
	// Output:
	// true
}

func ExampleUntypedConst() {
	constants.UntypedConst()
	// Output:
	// false
	// false
}
