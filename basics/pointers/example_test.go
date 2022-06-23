package pointers_test

import (
	"basics/pointers"
	"fmt"
)

func ExamplePassByValue() {
	val := 8
	pointers.PassByValue(val)
	fmt.Println("After passing value:", val)
	// Output: After passing value: 8
}

func ExamplePassByReference() {
	val := 8
	pointers.PassByReference(&val)
	fmt.Println("After derefence:", val)
	// Output: After derefence: 100
}

func ExamplePassMoreByReferences() {
	s := "This is going to change"
	b := false
	r := '🔥'
	f := 2.139284094893
	fmt.Printf("Before changing values:\n%q\n%t\n%s\n%f\n", s, b, string(r), f)
	pointers.PassMoreByReferences(&s, &b, &r, &f)
	fmt.Printf("After changing values:\n%q\n%t\n%s\n%f", s, b, string(r), f)
	// Output:
	// Before changing values:
	// "This is going to change"
	// false
	// 🔥
	// 2.139284
	// After changing values:
	// "Dereferenced and changed"
	// true
	// 🤡
	// 3.141590
}
