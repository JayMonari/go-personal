package method_test

import (
	"basics/method"
	"fmt"
)

func ExampleNew() {
	g := method.New("Ghidra", false, 0xBEEF)
	fmt.Print(g)
	// Output: Hi! I'm Ghidra and my favorite number is 48879
}

func ExampleGopher_DoesChangeFavNumber() {
	g := method.New("Gorple", false, 99)
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
	g := method.New("Galum", false, 99)
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

func ExampleStartCoding() {
	g := method.New("Geany", false, 3333)
	method.StartCoding(g)
	fmt.Println(g)
	// Output:
	// Let me get back to you after I'm done coding.
}

func ExampleStopCoding() {
	g := method.New("Ghjil", true, 0b010101011110111111)
	method.StopCoding(g)
	fmt.Println(g)
	// Output:
	// Hi! I'm Ghjil and my favorite number is 87999
}

func ExampleGopher_StartCoding() {
	g := method.New("Garkicye", true, 0xBA5ED)
	g.StartCoding()
	fmt.Println(g)
	// Output:
	// Let me get back to you after I'm done coding.
}

func ExampleGopher_StopCoding() {
	g := method.New("Gremeri", true, 87999)
	g.StopCoding()
	fmt.Println(g)
	// Output:
	// Hi! I'm Gremeri and my favorite number is 87999
}

func ExampleGopher_String() {
	fmt.Println(method.New("Gasitti", false, 42))
	fmt.Println(method.New("Gon", true, 42))
	// Output:
	// Hi! I'm Gasitti and my favorite number is 42
	// Let me get back to you after I'm done coding.
}
