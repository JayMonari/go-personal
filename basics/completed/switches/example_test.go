package switches_test

import "basics/switches"

func ExampleSwitchBasic() {
	switches.SwitchBasic()
	// Output:
	// i does not have a matching case.
}

func ExampleSwitchMultiple() {
	switches.SwitchMultiple()
	// Output:
	// Must have some Summer flare.
}

func ExampleSwitchType() {
	switches.SwitchType(true)
	switches.SwitchType(8)
	switches.SwitchType([]string{"some", "strings"})
	switches.SwitchType(struct{}{})
	// Output:
	// You bool! I knew it was you all along.
	// You seem like an int-eresting type.
	// Hey, hey. Save me a slice! []string
	// We've never seen a struct {} like this.
}

func ExampleSwitchNoValue() {
	switches.SwitchNoValue()
	// Output:
	// Looks like it's Spring!
}

func ExampleSwitchFallthrough() {
	switches.SwitchFallthrough()
	// Output:
	// Floor number three
	// Floor number two
	// Floor number one
	// Now arriving bottom floor.
}
