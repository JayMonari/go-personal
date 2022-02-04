package main

func ExampleSwitchBasic() {
	SwitchBasic()
	// Output:
	// i does not have a matching case.
}

func ExampleSwitchMultiple() {
	SwitchMultiple()
	// Output:
	// Must have some Summer flare.
}

func ExampleSwitchType() {
	SwitchType()
	// Output:
	// You bool! I knew it was you all along.
	// You seem like an int-eresting type.
	// Hey, hey. Save me a slice! []string
	// We've never seen a struct {} like this.
}

func ExampleSwitchNoValue() {
	SwitchNoValue()
	// Output:
	// Looks like it's Spring!
}

func ExampleSwitchFallthrough() {
	SwitchFallthrough()
	// Output:
	// Floor number three
	// Floor number two
	// Floor number one
	// Now arriving bottom floor.
}
