package variables

import "fmt"

// DeclareVarExplicit shows that we can explicitly type a variable to a
// specific type if we wanted to.
func DeclareVarExplicit() {
	// It works for all types: string, bool, int, float, rune...
	var name string = "Jay"
	// Emojis are unicode characters, they are supported in Go through runes
	var emoji rune = '😁'
	// You can do multiple declarations
	var min, max int = 0, 1000
	var isTrue bool = true
	fmt.Printf("My Name's %s! %s,\nFrom %d to %d\nGiven bool value: %t",
		name, string(emoji),
		min, max,
		isTrue)
}

// DeclareVarImplicit shows that we don't need to declare what the type is as
// long as we initialize it on the right.
func DeclareVarImplicit() {
	// It works for all types: string, bool, int, float, rune...
	var name = "Does anyone have any room for"
	// Emojis are unicode characters, they are supported in Go through runes
	var emoji = '🥧'
	var pi = 3.14159
	var isTrue = false
	fmt.Printf("%s %s,\nMore Pi %f\nGiven bool value: %t",
		name, string(emoji),
		pi,
		isTrue)
}

// DeclareVarDefault shows off all of the zero values or what might also be
// thought of as default values for each type when they are not assigned
// anything.
func DeclareVarDefault() {
	var name string
	var x int
	var f float32
	var isHard bool
	var r rune
	fmt.Printf(`
string default value is: %q
int default value is: %d
float32 default value is: %f
bool default value is: %t
rune default value is: %d or %q`,
		name, x, f, isHard, r, string(r))
}

// AssignmentOperator shows the preferred and idomatic way of declaring
// variables in Go. With the `:=` operator
func AssignmentOperator() {
	// It works for all types: string, bool, int, float, rune...
	statement := "Short and Sweet. Very Nice!"
	emoji := '👍'
	mySpecialNumber := 12648430
	isQuick := true
	fmt.Printf("%s %s\nI really need %x\nStill works? %t",
		statement, string(emoji), mySpecialNumber, isQuick)
}
