package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	Bool()
	fmt.Println()
	String()
	fmt.Println()
	Int()
	fmt.Println()
	Rune()
	fmt.Println()
	Float()
	fmt.Println()
	Complex()
	fmt.Println()
}

func Bool() {
	fmt.Println("true || false = true. false is a bool Type")
	fmt.Printf("%t || %t = %t. %t is a %T Type", true, false, true || false, false, false)
}

func String() {
	fmt.Println("go", "+", "gopher", "=", `"gogopher"`, "and is of Type string")
	fmt.Printf("%s + %s = %q and is of Type %T", "go", "gopher", "go"+"gopher", "")
}

func Int() {
	fmt.Println("2 + 2 =", 2+2, "and is of Type int")
	fmt.Printf("%d + %d = %d and is of Type %T", 2, 2, 2+2, 0)
}

func Rune() {
	fmt.Println("'k' is an int32 Type. When strings are built, they use rune values. Another way to say rune is int32, they mean the same thing!\n'k' is actually 107")
	fmt.Printf("'%s' is an %T Type. When strings are built, they use rune values. Another way to say rune is int32, they mean the same thing!\n'%s' is actually %d", string('k'), 'k', string('k'), 'k')
}

func Float() {
	fmt.Println("1.23 + 4.56 =", 1.23+4.56, "and is of Type float64")
	fmt.Printf("%.2f + %.2f = %.2f and is of Type %T", 1.23, 4.56, 1.23+4.56, 0.0)
}

func Complex() {
	fmt.Println("(2.94-2.31i) + (1.43+2.65i) = (4.37+0.341i) and is of Type complex128")
	fmt.Printf("%.3g + %.3g = %.3g and is of Type %T", cmplx.Acos(-5+1i), cmplx.Acos(1+-7i), cmplx.Acos(-5+1i)+cmplx.Acos(1+-7i), cmplx.Acos(0))
}
