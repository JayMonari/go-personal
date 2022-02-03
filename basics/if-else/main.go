package main

import "fmt"

func main() {
	IfElse()
	ElseIf()
	DeclareInIf()
}

// There is no ternary if in Go. If you don't know what that means, then don't
// worry it doesn't exist and you don't need to care about it.
func IfElse() {
	i := 5
	if i < 4 {
		fmt.Println("This statement will not be printed.")
	}
	if i > 4 {
		fmt.Println("This statement will print because i > 4 ==", i > 4)
	}

	if i > 10 {
		fmt.Println("This statement will not be printed.")
	} else {
		fmt.Println("We will always reach and print this statement.")
	}
}

func ElseIf() {
	i := 8
	if i > 8 {
		fmt.Println("This statement will not be printed.")
	} else if i == 8 {
		fmt.Println("i == 8 so we will reach into this else if block!")
	} else {
		fmt.Println("This statement will not be printed.")
	}
}

func DeclareInIf() {
	if i := 5; i < 4 {
		fmt.Println("This statement will not be printed.")
	} else if i == 5 {
		fmt.Println("We can reuse i for the entire if-else statement!")
	}
	// XXX: Since `i` was declared in the scope of the `if` statement, it doesn't
	// exist outside of that scope, so if we uncomment this, we will get an
	// error: "Undeclared name: i"
	// if i == 5 {
	// 	fmt.Println("This will never work.")
	// }
}
