package operators

import (
	"fmt"
)

// Comparison shows how to compare values in programming. This is
// essential part of development and for control flow.
func Comparison() {
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

// Logical shows how to introduce more granular logic into your
// program with the logical operators.
func Logical() {
	theAnswer := '*'
	toTheUniverse := 42
	percentC := 42
	percentD := 42
	// BOTH statements must evaluate to the `bool` true or it will output false.
	fmt.Printf("Is the answer to the universe as rune %c or as int %d? %t\n",
		percentC, percentD, theAnswer == 42 && toTheUniverse == '*')
	fmt.Printf("Is 84 NOT equal to 84 AND something true? %t\n",
		84 != 84 && true)
	fmt.Println()

	this := "Some Value"
	that := "Other Value"
	lastRuneThis := this[len(this)-1]
	lastRuneThat := that[len(that)-1]
	// One of the statements must be true for the output to be true.
	fmt.Printf("Is one of these statements true?\n"+
		`1. %q == "Some Value" OR 2. %q == "DOESN'T MATTER"`+
		"?\nAnswer: %t\n", this, that,
		this == "Some Value" || that == "DOESN'T MATTER")

	fmt.Printf("If the second one is true?\n"+
		`1. %q == "FALSE" OR 2. %q == "Other Value"`+
		"?\nAnswer: %t\n", this, that,
		this == "Some Value" || that == "DOESN'T MATTER")

	fmt.Printf("How many can we use?\n"+
		"1. %q == %q OR 2. %c == %c OR 3. %d == %d OR 4. %c == %c\n Answer: %t\n",
		this, that, this[0], that[0], 64, 640, lastRuneThis, lastRuneThat,
		this == that || this[0] == that[0] || 42+24 == 640 || lastRuneThis == lastRuneThat)
	fmt.Println()

	goodFood := true
	badSmell := false
	fmt.Printf("Would you eat goodFood? %t\nWhat if it smelled bad? %t\n",
		goodFood, badSmell)
	fmt.Printf(
		"Would you eat NOT goodFood? %t\nWhat if it did NOT smell bad? %t\n",
		!goodFood, !badSmell)

	fmt.Printf(
		"🤔 What if it was NOT goodFood, but it did NOT smellBad? %t\n",
		!goodFood || !badSmell)

	fmt.Printf(
		"🤔 What if it was goodFood AND it did NOT smellBad? %t\n",
		goodFood && !badSmell)
}

// Arithmetic shows that all of the operators you know and love ❤️ from
// math class 🫠 still work the same as in class.
func Arithmetic() {
	var (
		myInt     = 21
		myFloat   = 6.28
		myComplex = -5 + 2i
	)
	fmt.Printf("nine plus ten does NOT equal twenty-one\n"+
		"9 + 10 != %d: %t\n", myInt, 9+10 != myInt)
	fmt.Printf("Is pi greater than 3?\n"+
		"%.2f - 3.14 >= 3: %t\n", myFloat, myFloat-3.14 >= 3)
	fmt.Printf("my head hurts 🥴\n%.3g * %.3g = %.3g\n",
		myComplex, myComplex, myComplex*myComplex)
	fmt.Printf("tau divided by tau is one\n"+
		"%.4f / %.4f: %.4f", myFloat, myFloat, myFloat/myFloat)
	fmt.Println()

	// ONLY FOR THE INTS!!
	// 21 % 5 == 21 - 5 - 5 - 5 - 5 == 1
	fmt.Printf("%d %% 5 == 1\n"+
		"What whole number remains when I try to divide %d by 5? %d\n",
		myInt, myInt, myInt%5)

	// 21 % 4 == 21 - 4 - 4 - 4 - 4 - 4 == 1
	fmt.Printf("%d %% 4 == 1\n"+
		"What whole number remains when I try to divide %d by 4? %d\n",
		myInt, myInt, myInt%4)

	// 21 % 3 == 21 - 3 - 3 - 3 - 3 - 3 - 3 - 3 == 0
	fmt.Printf("%d %% 3 == 0\n"+
		"What whole number remains when I try to divide %d by 3? %d\n",
		myInt, myInt, myInt%3)

	// 21 % 21 == 21 - 21 == 0
	fmt.Printf("%d %% 21 == 0\n"+
		"What whole number remains when I try to divide %d by 21? %d\n",
		myInt, myInt, myInt%21)
}

// String shows the only operator that works with the `string` type,
// which is the `+` operator or concatenation
func String() {
	const blt = "Bacon🥓Lettuce🥬Tomato🍅"
	var (
		t  = "Tomato🍅"
		bl = "Bacon🥓"
		l  = "Lettuce🥬"
	)
	// bl = bl + l
	bl += l
	fmt.Printf("%s == %s: %t\nIs blt equal to bl+t? true\n",
		blt, bl+t, blt == bl+t)
	fmt.Printf("%s != %s: %t\nIs blt NOT equal to bl? true\n",
		blt, bl, blt == bl)
}

// Bitwise shows how to manipulate the bits of integer types:
// `int`, int8`, `int16`, `int32`, `int64`,
// `uint`, uint8`, `uint16`, `uint32`, `uint64`,
func Bitwise() {
	// bitwise AND -- & -- works just like logical AND -- &&
	// BOTH statements must be true to output true otherwise it's false
	// BOTH numbers must be one to output one otherwise it's zero
	fmt.Printf("1011111101 &\n1101010011 ==\n%b\nBonus as an int %d\n",
		0b1011111101&0b1101010011, 0b1011111101&0b1101010011)
	fmt.Println()

	// bitwise OR -- | -- works just like logical OR -- ||
	// One of the statements must be true to output true otherwise it's false
	// One of the numbers must be one to output one otherwise it's zero
	fmt.Printf("1111100000 |\n0000011111 ==\n%b\nBonus as an int %d\n",
		0b1111100000|0b0000011111, 0b1111100000|0b0000011111)
	fmt.Println()

	// ^    bitwise XOR            integers
	// bitwise XOR -- ^ -- or more formally eXclusive OR
	// works like the logical NOT -- ! -- operator with two statements.
	// Logically it is the equivalent of saying:
	//   One of the statements is false AND one of the statements is true
	//   Then I will output true; i.e. if both statements are NOT exclusive
	//   Then I'm going to output false.
	// If both numbers are 0 --> 0
	// If both numbers are 1 --> 0
	// If one number is 1 and the other number is 0 --> 1
	fmt.Printf("010111110000 ^\n101011110000 ==\n%b\nBonus as an int %d\n",
		0b010111110000^0b101011110000, 0b010111110000^0b101011110000)
	fmt.Println()

	// bit clear -- &^ -- is used to NUKE 💣 a bit.
	// When you want a specific bit or set of bits gone you can use bit clear
	// It is logically the same as NOT -- ! -- and AND -- && -- together
	// ! (true && true) --> ! (true) --> false
	// ! (1 && 1) --> ! (1) --> 0
	fmt.Printf("111111111111 &^\n000001111000 ==\n%b\nBonus as an int %d\n",
		0b111111111111&^0b000001111000, 0b111111111111&^0b000001111000)
	fmt.Println()

	// shifting bits with left shift -- << -- and right shift -- >> --
	// moves all bits by a given amount.
	fmt.Printf("Shift all bits one left\n"+
		"0100100101 << 1 ==\n%b\nBonus as an int %d\n",
		0b0100100101<<1, 0b0100100101<<1)
	fmt.Printf("Shift all bits two right\n"+
		"0100100101 >> 2 ==\n%010b\nBonus as an int %d\n",
		0b0100100101>>2, 0b0100100101>>2)
}
