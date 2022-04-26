package functions_test

import (
	"basics/functions"
	"fmt"
)

// This will not work, uncomment it and see what error it gives you.
// func ExampleprivateFunc() {
// 	// XXX: privateFunc not exported by package functions
// 	functions.privateFunc()
// }

func ExampleFuncPublic() {
	functions.FuncPublic()
	// Output:
	// This function is exported and can be called anywhere.
}

func ExampleFuncWithParams() {
	functions.FuncWithParams("Mechanical Arm", 9, '🦾')
	// Output:
	// Mechanical Arm looks like 🦾 and is a 9/10
}

func ExampleFuncWithReturn() {
	fmt.Println(functions.FuncWithReturn())
	// Output:
	// It's just this easy to return a type
}

func ExampleFuncWithMultipleReturn() {
	fmt.Println(functions.FuncWithMultipleReturn())
	// Output:
	// [1 2 3 4 5] true
}

func ExampleFuncWithNamedReturn() {
	fmt.Println(functions.FuncWithNamedReturn("Gamba",
		"https://", "gophergo.dev", "/fun-with-funcs", "?isFun=yes&isEasy=yes"))
	// Output:
	// Gamba@gophergo.dev https://gophergo.dev/fun-with-funcs?isFun=yes&isEasy=yes
}

func ExampleFuncVariadic() {
	fmt.Println(functions.FuncVariadic())
	fmt.Println(functions.FuncVariadic(1, 2, 3))
	nums := []int{4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(functions.FuncVariadic(nums...))
	// Output:
	// 0
	// 6
	// 72
}