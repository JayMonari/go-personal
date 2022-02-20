package methods_test

import (
	"basics/methods"
	"fmt"
)

func ExampleNew() {
	g := methods.New("Ghidra", false, 0xBEEF)
	fmt.Print(g)
	// Output: Hi! I'm Ghidra and my favorite number is 48879
}

func ExampleGopher_String() {
	g := methods.New("Gigi", true, 42)
	fmt.Println(g)

	g.StopCoding()
	fmt.Println(g)

	g.StartCoding()
	fmt.Println(g)
	// Output:
	// Let me get back to you after I'm done coding.
	// Hi! I'm Gigi and my favorite number is 42
	// Let me get back to you after I'm done coding.
}

func ExampleGopher_DoesChangeFavNumber() {
	g := methods.New("Gorple", false, 99)
	g.DoesChangeFavNumber(800)
	fmt.Println(g)
	g.DoesChangeFavNumber(-28298)
	fmt.Println(g)
	g.DoesChangeFavNumber(0xDEAD)
	fmt.Println(g)
	g.DoesChangeFavNumber(0xC0FFEE)
	fmt.Println(g)
	// Output:
	// Hi! I'm Gorple and my favorite number is 800
	// Hi! I'm Gorple and my favorite number is -28298
	// Hi! I'm Gorple and my favorite number is 57005
	// Hi! I'm Gorple and my favorite number is 12648430
}

func ExampleGopher_DoesNotChangeFavNumber() {
	g := methods.New("Galum", false, 99)
	g.DoesNotChangeFavNumber(800)
	fmt.Println(g)
	g.DoesNotChangeFavNumber(-28298)
	fmt.Println(g)
	g.DoesNotChangeFavNumber(0xDEAD)
	fmt.Println(g)
	g.DoesNotChangeFavNumber(0xC0FFEE)
	fmt.Println(g)
	// Output:
	// Hi! I'm Galum and my favorite number is 99
	// Hi! I'm Galum and my favorite number is 99
	// Hi! I'm Galum and my favorite number is 99
	// Hi! I'm Galum and my favorite number is 99
}
