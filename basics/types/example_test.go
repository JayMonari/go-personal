package types_test

import "basics/types"

func ExampleBool() {
	types.Bool()
	// Output:
	// true || false = true. false is a bool Type
	// true || false = true. false is a bool Type
}

func ExampleString() {
	types.String()
	// Output:
	// gopher + go = "gophergo" and is of Type string
	// gopher + go = "gophergo" and is of Type string
}

func ExampleInt() {
	types.Int()
	// Output:
	// 2 + 2 = 4 and is of Type int
	// 2 + 2 = 4 and is of Type int
}

func ExampleRune() {
	types.Rune()
	// Output:
	// 'k' is an int32 Type. When strings are built, they use rune values.
	// Another way to say rune is int32, they mean the same thing!
	// 'k' is actually 107
	// 'k' is an int32 Type. When strings are built, they use rune values.
	// Another way to say rune is int32, they mean the same thing!
	// 'k' is actually 107
}

func ExampleFloat() {
	types.Float()
	// Output:
	// 1.23 + 4.56 = 5.79 and is of Type float64
	// 1.23 + 4.56 = 5.79 and is of Type float64
}

func ExampleComplex() {
	types.Complex()
	// Output:
	// (2.94-2.31i) + (1.43+2.65i) = (4.37+0.341i) and is of Type complex128
	// (2.94-2.31i) + (1.43+2.65i) = (4.37+0.341i) and is of Type complex128
}
