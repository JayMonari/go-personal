package if_else

import "fmt"

// IfElse shows you how to control the flow of logic in your application using
// if and else statements. It's also good to be aware that there is no ternary
// operator in Go.
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

// ElseIf shows you that you can branch you logic as many times as you want
// with an `else if` block
func ElseIf() {
	i := 8
	if i > 8 {
		fmt.Println("This statement will not be printed.")
	} else if i == 6 {
		fmt.Println("This statement will not be printed.")
	} else if i == 7 {
		fmt.Println("This statement will not be printed.")
	} else if i == 8 {
		fmt.Println("i == 8 so we will reach into this else if block!")
	} else {
		fmt.Println("This statement will not be printed.")
	}
}

// DeclareInIf shows you that if you want your variable to be scoped to just an
// `if` block you can do it in Go!
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
