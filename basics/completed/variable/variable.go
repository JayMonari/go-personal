package variable

import "fmt"

// DeclareVarZero shows off all of the zero values or what might also be
// thought of as default values for each type when they are not assigned
// anything.
func DeclareVarZero() {
	var name string
	var x int
	var f float32
	var isHard bool
	var r rune
	var slice []int
	var map_ map[int]int
	fmt.Printf(`
string zero value is: %q
int zero value is: %d
float32 zero value is: %f
bool zero value is: %t
rune zero value is: %d or %q
slice zero value is: %+v
map zero value is: %+v`,
		name, x, f, isHard, r, string(r), slice, map_)
}

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
	var question = "Does anyone have any room for"
	// Emojis are unicode characters, they are supported in Go through runes
	var emoji = '🥧'
	var pi = 3.14159
	var isTrue = false
	fmt.Printf("%s %s,\nMore Pi %f\nGiven bool value: %t",
		question, string(emoji),
		pi,
		isTrue)
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
