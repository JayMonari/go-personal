package operators_test

import "basics/operators"

func ExampleComparison() {
	operators.Comparison()
	// Output:
	// trueVal == falseVal: false
	// Is trueVal equal to falseVal? false
	// trueVal != falseVal: true
	// Is trueVal NOT equal to falseVal? true
	//
	// name1 == name2: false
	// Is name1 equal to name2? false
	// name1 != name2: true
	// Is name1 NOT equal to name2? true
	//
	// lilNum1 == lilNum2: true
	// Is lilNum1 equal to lilNum2? true
	// lilNum1 != lilNum2: false
	// Is lilNum1 NOT equal to lilNum2? false
	//
	// lilNum1 < bigNum: true
	// Is lilNum1 less than bigNum? true
	// bigNum > lilNum2: true
	// Is bigNum greater than lilNum2? true
	// lilNum1 <= lilNum2: true
	// Is lilNum1 less than or equal to lilNum2? true
	// lilNum1 >= lilNum2: true
	// Is lilNum1 greater than or equal to lilNum2? true
}

func ExampleLogical() {
	operators.Logical()
	// Output:
	// Is the answer to the universe as rune * or as int 42? true
	// Is 84 NOT equal to 84 AND something true? false
	//
	// Is one of these statements true?
	// 1. "Some Value" == "Some Value" OR 2. "Other Value" == "DOESN'T MATTER"?
	// Answer: true
	// If the second one is true?
	// 1. "Some Value" == "FALSE" OR 2. "Other Value" == "Other Value"?
	// Answer: true
	// How many can we use?
	// 1. "Some Value" == "Other Value" OR 2. S == O OR 3. 64 == 640 OR 4. e == e
	//  Answer: true
	//
	// Would you eat goodFood? true
	// What if it smelled bad? false
	// Would you eat NOT goodFood? false
	// What if it did NOT smell bad? true
	// 🤔 What if it was NOT goodFood, but it did NOT smellBad? true
	// 🤔 What if it was goodFood AND it did NOT smellBad? true
}

func ExampleArithmetic() {
	operators.Arithmetic()
	// Output:
	// nine plus ten does NOT equal twenty-one
	// 9 + 10 != 21: true
	// Is pi greater than 3?
	// 6.28 - 3.14 >= 3: true
	// my head hurts 🥴
	// (-5+2i) * (-5+2i) = (21-20i)
	// tau divided by tau is one
	// 6.2800 / 6.2800: 1.0000
	// 21 % 5 == 1
	// What whole number remains when I try to divide 21 by 5? 1
	// 21 % 4 == 1
	// What whole number remains when I try to divide 21 by 4? 1
	// 21 % 3 == 0
	// What whole number remains when I try to divide 21 by 3? 0
	// 21 % 21 == 0
	// What whole number remains when I try to divide 21 by 21? 0
}

func ExampleString() {
	operators.String()
	// Output:
	// Bacon🥓Lettuce🥬Tomato🍅 == Bacon🥓Lettuce🥬Tomato🍅: true
	// Is blt equal to bl+t? true
	// Bacon🥓Lettuce🥬Tomato🍅 != Bacon🥓Lettuce🥬: false
	// Is blt NOT equal to bl? true
}

func ExampleBitwise() {
	operators.Bitwise()
	//Output:
	// 1011111101 &
	// 1101010011 ==
	// 1001010001
	// Bonus as an int 593
	//
	// 1111100000 |
	// 0000011111 ==
	// 1111111111
	// Bonus as an int 1023
	//
	// 010111110000 ^
	// 101011110000 ==
	// 111100000000
	// Bonus as an int 3840
	//
	// 111111111111 &^
	// 000001111000 ==
	// 111110000111
	// Bonus as an int 3975
	//
	// Shift all bits one left
	// 0100100101 << 1 ==
	// 1001001010
	// Bonus as an int 586
	// Shift all bits two right
	// 0100100101 >> 2 ==
	// 0001001001
	// Bonus as an int 73
}
