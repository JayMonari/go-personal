package if_else_test

import if_else "basics/if-else"

func ExampleIfElse() {
	if_else.IfElse()
	// Output:
	// This statement will print because i > 4 == true
	// We will always reach and print this statement.
}

func ExampleElseIf() {
	if_else.ElseIf()
	// Output:
	// i == 8 so we will reach into this else if block!
}

func ExampleDeclareInIf() {
	if_else.DeclareInIf()
	// Output:
	// We can reuse i for the entire if-else statement!
}
