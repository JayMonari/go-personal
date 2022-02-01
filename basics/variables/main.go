package main

import "fmt"

func main() {
	DeclareVarExplicit()
	fmt.Println()
	DeclareVarImplicit()
	fmt.Println()
	DeclareVarDefault()
	fmt.Println()
	AssignmentOperator()
	fmt.Println()
}

func DeclareVarExplicit() {
	// It works for all types: string, bool, int, float, rune...
	var name string = "Jay"
	// Emojis are unicode characters, they are supported in Go through runes
	// I got this one from https://emojipedia.org/beaming-face-with-smiling-eyes/
	var emoji rune = '😁'
	// You can do multiple declarations
	var min, max int = 0, 1000
	var isTrue bool = true
	fmt.Printf("My Name's %s! %s,\nFrom %d to %d\nGiven bool value: %t",
		name, string(emoji),
		min, max,
		isTrue)
}

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

func AssignmentOperator() {
	// It works for all types: string, bool, int, float, rune...
	statement := "Short and Sweet. Very Nice!"
	emoji := '👍'
	mySpecialNumber := 12648430
	isQuick := true
	fmt.Printf("%s %s\nI really need %x\nStill works? %t",
		statement, string(emoji), mySpecialNumber, isQuick)
}
