package constant_test

import (
	"basics/constant"
	"fmt"
)

func ExampleStuck() {
	fmt.Println(constant.Stuck)
	// Output:
	// This variable can never be reassigned.
}

func ExampleHeartEyes() {
	// We can convert constant just like variables
	fmt.Println(string(constant.HeartEyes))
	// Output:
	// 😍
}

func ExampleArithmetic() {
	fmt.Println(constant.Arithmetic)
	// Output:
	// 175.38731365097925
}

func ExampleAlwaysTrue() {
	fmt.Println(constant.AlwaysTrue)
	// Output:
	// true
}

func ExampleUntypedConst() {
	constant.UntypedConst()
	// Output:
	// false
	// false
}

func ExamplePrint() {
	constant.Print(constant.UntypedString)
	// Output:
	// I fit wherever the underlying type of something is a string!
}
