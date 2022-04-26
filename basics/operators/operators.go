package operators

import (
	"fmt"
)

// OperatorsComparison shows how to compare values in programming. This is
// essential part of development and for control flow.
func OperatorsComparison() {
	const (
		trueVal  = true
		falseVal = false
	)
	fmt.Printf("trueVal == falseVal: %t\n"+
		"Is trueVal equal to falseVal? false\n",
		trueVal == falseVal)
	fmt.Printf("trueVal != falseVal: %t\n"+
		"Is trueVal NOT equal to falseVal? true\n",
		trueVal != falseVal)
	fmt.Println()

	const (
		name1 = "Goober"
		name2 = "Goomba"
	)
	fmt.Printf("name1 == name2: %t\nIs name1 equal to name2? false\n",
		name1 == name2)
	fmt.Printf("name1 != name2: %t\nIs name1 NOT equal to name2? true\n",
		name1 != name2)
	fmt.Println()

	const (
		bigNum  = 100000
		lilNum1 = 10
		lilNum2 = 10
	)
	fmt.Printf("lilNum1 == lilNum2: %t\nIs lilNum1 equal to lilNum2? true\n",
		lilNum1 == lilNum2)
	fmt.Printf("lilNum1 != lilNum2: %t\nIs lilNum1 NOT equal to lilNum2? false\n",
		lilNum1 != lilNum2)
	fmt.Println()

	fmt.Printf("lilNum1 < bigNum: %t\nIs lilNum1 less than bigNum? true\n",
		lilNum1 < bigNum)
	fmt.Printf("bigNum > lilNum2: %t\nIs bigNum greater than lilNum2? true\n",
		bigNum > lilNum2)
	fmt.Printf(
		"lilNum1 <= lilNum2: %t\nIs lilNum1 less than or equal to lilNum2? true\n",
		lilNum1 <= lilNum2)
	fmt.Printf("lilNum1 >= lilNum2: %t\n"+
		"Is lilNum1 greater than or equal to lilNum2? true\n",
		lilNum1 >= lilNum2)
}

// OperatorsLogical shows how to introduce more granular logic into your
// program with the logical operators.
func OperatorsLogical() {
	// &&    conditional AND    p && q  is  "if p then q else false"
	// ||    conditional OR     p || q  is  "if p then true else q"
	// !     NOT                !p      is  "not p"
}

// OperatorsArithmetic shows that all of the operators you know and love ❤️ from
// math class 🫠 still work the same as in class.
func OperatorsArithmetic() {
	var (
		myInt     = 21
		myFloat   = 6.28
		myComplex = -5 + 2i
	)
	fmt.Printf("9 + 10 != %d: %t\n", myInt, 9+10 != myInt)
	fmt.Printf("%.2f - 3.14 >= 3: %t\n", myFloat, myFloat-3.14 >= 3)
	fmt.Printf("%.3g * %.3g = %.3g\n",
		myComplex, myComplex, myComplex*myComplex)
	fmt.Printf("%.4f / %.4f: %.4f", myFloat, myFloat, myFloat/myFloat)

	fmt.Println()
	// ONLY FOR THE INTS!!

	// 21 % 4 == 21 - 4 - 4 - 4 - 4 == 1
	fmt.Printf("%d %% 4 == 1\n"+
		"What whole number is left when I divide %d by 4? %d\n",
		myInt, myInt, myInt%4)

	// 21 % 3 == 21 - 3 - 3 - 3 - 3 - 3 - 3 - 3 == 0
	fmt.Printf("%d %% 3 == 0\n"+
		"What whole number is left when I divide %d by 3? %d\n",
		myInt, myInt, myInt%3)

	// 21 % 21 == 21 - 21 == 0
	fmt.Printf("%d %% 21 == 0\n"+
		"What whole number is left when I divide %d by 21? %d\n",
		myInt, myInt, myInt%21)
}

// OperatorsString shows the only operator that works with the `string` type,
// which is the `+` operator or concatenation
func OperatorsString() {
	const blt = "Bacon🥓Lettuce🥬Tomato🍅"
	var (
		t  = "Tomato🍅"
		bl = "Bacon🥓"
		l  = "Lettuce🥬"
	)
	bl += l
	fmt.Printf("%s == %s: %t\nIs blt equal to bl+t? true\n", blt, bl+t, blt == bl+t)
	fmt.Printf("%s != %s: %t\nIs blt NOT equal to bl? true\n", blt, bl, blt == bl)
}

// OperatorBitwise shows how to manipulate the bits of integer types:
// `int`, int8`, `int16`, `int32`, `int64`,
// `uint`, uint8`, `uint16`, `uint32`, `uint64`,
func OperatorBitwise() {
	// &    bitwise AND            integers
	// |    bitwise OR             integers
	// ^    bitwise XOR            integers
	// &^   bit clear (AND NOT)    integers

	// <<   left shift             integer << integer >= 0
	// >>   right shift            integer >> integer >= 0
}
