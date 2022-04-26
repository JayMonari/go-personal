package operators_test

import "basics/operators"

func ExampleOperatorsComparison() {
	operators.OperatorsComparison()
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

func ExampleOperatorsString() {
	operators.OperatorsString()
	// Output:
	// Bacon🥓Lettuce🥬Tomato🍅 == Bacon🥓Lettuce🥬Tomato🍅: true
	// Is blt equal to bl+t? true
	// Bacon🥓Lettuce🥬Tomato🍅 != Bacon🥓Lettuce🥬: false
	// Is blt NOT equal to bl? true
}
