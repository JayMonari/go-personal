package function_test

import (
	"basics/function"
	"fmt"
)

// This will not work, uncomment it and see what error it gives you.
// func ExampleprivateFunc() {
// 	// NOTE(jay): privateFunc not exported by package function
// 	function.privateFunc()
// }

func ExampleFuncPublic() {
	function.FuncPublic()
	// Output:
	// This function is exported and can be called anywhere.
}

func ExampleFuncWithParams() {
	function.FuncWithParams("Mechanical Arm", 9, '🦾')
	// Output:
	// Mechanical Arm looks like 🦾 and is a 9/10
}

func ExampleFuncWithReturn() {
	fmt.Println(function.FuncWithReturn())
	// Output:
	// It's just this easy to return a type
}

func ExampleFuncWithMultipleReturn() {
	fmt.Println(function.FuncWithMultipleReturn())
	// Output:
	// [1 2 3 4 5] true
}

func ExampleFuncWithNamedReturn() {
	fmt.Println(function.FuncWithNamedReturn("Gamba",
		"https://", "gophergo.dev", "/fun-with-funcs", "?isFun=yes&isEasy=yes"))
	// Output:
	// Gamba@gophergo.dev https://gophergo.dev/fun-with-funcs?isFun=yes&isEasy=yes
}

func ExampleFuncVariadic() {
	fmt.Println(function.FuncVariadic())
	fmt.Println(function.FuncVariadic(1, 2, 3))
	nums := []int{4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(function.FuncVariadic(nums...))
	// Output:
	// 0
	// 6
	// 72
}
