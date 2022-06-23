package switches_test

import "basics/switches"

func ExampleBasic() {
	switches.Basic()
	// Output:
	// i does not have a matching case.
}

func ExampleMultiple() {
	switches.Multiple()
	// Output:
	// Must have some Summer flare.
}

func ExampleType() {
	switches.Type(true)
	switches.Type(8)
	switches.Type([]string{"some", "strings"})
	switches.Type(struct{}{})
	// Output:
	// You bool! I knew it was you all along.
	// You seem like an int-eresting type.
	// Hey, hey. Save me a slice! []string
	// We've never seen a struct {} like this.
}

func ExampleNoValue() {
	switches.NoValue()
	// Output:
	// Looks like it's Spring!
}

func ExampleFallthrough() {
	switches.Fallthrough()
	// Output:
	// Floor number three
	// Floor number two
	// Floor number one
	// Now arriving bottom floor.
}
