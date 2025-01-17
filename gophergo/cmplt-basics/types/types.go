package types

import (
	"fmt"
	"math/cmplx"
)

// Bool shows us the type bool which is short for boolean. A bool is either
// true or false.
func Bool() {
	fmt.Println("true || false = true. false is a bool Type")
	fmt.Printf("%t || %t = %t. %t is a %T Type",
		true, false, true || false, false, false)
}

// String shows us how to make a sequence of characters (or runes in Go) in a
// row, surrounded by double quote marks "".
func String() {
	fmt.Println("gopher", "+", "go", "=", `"gophergo"`, "and is of Type string")
	fmt.Printf("%s + %s = %q and is of Type %T",
		"gopher", "go", "gopher"+"go", "")
}

// Rune shows us how to represent and manipulate each value in a string.
func Rune() {
	fmt.Println("'k' is an int32 Type. When strings are built, they use rune " +
		"values.\nAnother way to say rune is int32, they mean the same thing!\n" +
		"'k' is actually 107")
	fmt.Printf("'%s' is an %T Type. When strings are built, they use rune "+
		"values.\nAnother way to say rune is int32, they mean the same thing!\n"+
		"'%s' is actually %d", string('k'), 'k', string('k'), 'k')
}

// Int short for integer shows us how to use the int type in Go. We can also do
// arithmetic like we would expect.
func Int() {
	// NOTE(jay): When we have big numbers we can separate them with an
	// underscore `_` like we do with comma `,` or period `.`
	// So 1,234 or 1.234 becomes 1_234 in Go
	fmt.Println("1234567 + 2 =", 1_234_567+2, "and is of Type int")
	fmt.Printf("%d + %d = %d and is of Type %T", 1_234_567, 2, 1_234_567+2, 0)
}

// Float short for floating point number shows how to represent numbers with
// decimal values in Go.
func Float() {
	// NOTE(jay): When we have big floating point numbers we can separate them
	// with an underscore `_` like we do with comma `,` or period `.`
	// For example 980,222.0123 or 980.222,0123 becomes 980_222.012_3
	fmt.Println("1_234.567_890_1 + 4.56 =", 1_234.567_890_1+4.56,
		"and is of Type float64")
	fmt.Printf("%.7f + %.2f = %.7f and is of Type %T",
		1_234.567_890_1, 4.56, 1_234.567_890_1+4.56, 0.0)
}

// Complex shows how to use complex numbers in Go... If you would ever need
// them ¯\_(ツ)_/¯
func Complex() {
	fmt.Println("(2.94-2.31i) + (1.43+2.65i) = (4.37+0.341i)" +
		" and is of Type complex128")
	fmt.Printf("%.3g + %.3g = %.3g and is of Type %T",
		cmplx.Acos(-5+1i), cmplx.Acos(1+-7i), cmplx.Acos(-5+1i)+cmplx.Acos(1+-7i),
		cmplx.Acos(0))
}
