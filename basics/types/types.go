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
	fmt.Println("go", "+", "gopher", "=", `"gogopher"`, "and is of Type string")
	fmt.Printf("%s + %s = %q and is of Type %T",
		"go", "gopher", "go"+"gopher", "")
}

// Int short for integer shows us how to use the int type in Go. We can also do
// arithmetic like we would expect.
func Int() {
	fmt.Println("2 + 2 =", 2+2, "and is of Type int")
	fmt.Printf("%d + %d = %d and is of Type %T", 2, 2, 2+2, 0)
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

// Float short for floating point number shows how to represent numbers with
// decimal values in Go.
func Float() {
	fmt.Println("1.23 + 4.56 =", 1.23+4.56, "and is of Type float64")
	fmt.Printf("%.2f + %.2f = %.2f and is of Type %T", 1.23, 4.56, 1.23+4.56, 0.0)
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
