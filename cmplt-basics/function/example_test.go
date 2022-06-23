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

func ExamplePublic() {
	function.Public()
	// Output:
	// This function is exported and can be called anywhere.
}

func ExampleWithParams() {
	function.WithParams("Mechanical Arm", 9, '🦾')
	// Output:
	// Mechanical Arm looks like 🦾 and is a 9/10
}

func ExampleWithReturn() {
	fmt.Println(function.WithReturn())
	// Output:
	// It's just this easy to return a type
}

func ExampleWithMultipleReturn() {
	fmt.Println(function.WithMultipleReturn())
	// Output:
	// [1 2 3 4 5] true
}

func ExampleWithNamedReturn() {
	fmt.Println(function.WithNamedReturn("Gamba",
		"https://", "gophergo.dev", "/fun-with-funcs", "?isFun=yes&isEasy=yes"))
	// Output:
	// Gamba@gophergo.dev https://gophergo.dev/fun-with-funcs?isFun=yes&isEasy=yes
}

func ExampleVariadic() {
	fmt.Println(function.Variadic())
	fmt.Println(function.Variadic(1, 2, 3))
	nums := []int{4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(function.Variadic(nums...))
	// Output:
	// 0
	// 6
	// 72
}
