package pointer_test

import (
	"basics/pointer"
	"fmt"
)

func ExamplePassByValue() {
	intVal := 8
	pointer.PassByValue(intVal)
	fmt.Println("intVal after passing value:", intVal)
	// Output:
	// intVal in function: 100
	// intVal after passing value: 8
}

func ExamplePassByReference() {
	intVal := 8
	intPtr := &intVal
	pointer.PassByReference(intPtr)
	// pointer.PassByReference(&intVal) <-- Also works!
	fmt.Println("intVal after derefence:", intVal)
	// Output:
	// intVal in function: 100
	// intVal after derefence: 100
}

func ExamplePassMoreByReferences() {
	s := "This is going to change"
	b := false
	r := '🔥'
	f := 2.139284094893
	fmt.Printf(`Before changing values:
string value: %q
bool value: %t
rune as emoji value: %s
float value: %f

`, s, b, string(r), f)

	pointer.PassMoreByReferences(&s, &b, &r, &f)
	fmt.Printf(`After changing values:
string value: %q
bool value: %t
rune as emoji value: %s
float value: %f

`, s, b, string(r), f)
	// Output:
	// Before changing values:
	// string value: "This is going to change"
	// bool value: false
	// rune as emoji value: 🔥
	// float value: 2.139284
	//
	// After changing values:
	// string value: "Dereferenced and changed"
	// bool value: true
	// rune as emoji value: 🤡
	// float value: 3.141590
}

func ExamplePassCollections() {
	sl := []string{"Look", "at", "all", "my", "values"}
	mp := map[string]rune{
		"boxing":     '🥊',
		"chestnut":   '🌰',
		"ocean":      '🌊',
		"heart-eyes": '😍',
		"microphone": '🎤',
		"hibiscus":   '🌺',
	}
	asCopy := pointer.ChangeThings{
		Int:  777,
		Str:  "Will I change?",
		Rune: '🌔',
	}
	asCopy2 := pointer.ChangeThings{
		Int:  777,
		Str:  "Will I change?",
		Rune: '🌔',
	}
	asPtr := &asCopy2
	/////////////////////////////////////////////////////////////////////////////
	fmt.Printf(`Before changing values:
slice value: %#v
map value: %#v
copy of struct value: %#v
pointer to struct value: %#v
copy2 of struct value: %#v

`, sl, mp, asCopy, asPtr, asCopy2)

	pointer.PassCollections(sl, mp, asCopy, asPtr)
	/////////////////////////////////////////////////////////////////////////////
	fmt.Printf(`After changing values:
slice value: %#v
map value: %#v
copy of struct value: %#v
pointer to struct value: %#v
copy2 of struct value: %#v

`, sl, mp, asCopy, asPtr, asCopy2)

	/////////////////////////////////////////////////////////////////////////////
	sl = nil
	mp = nil
	asPtr = nil
	asCopy = pointer.ChangeThings{}
	asCopy2 = pointer.ChangeThings{}
	fmt.Printf(`Actually remove all elements from collection types:
pointer: %#v
pointer: %#v
pointer: %#v
literal: %#v
literal: %#v
`, sl, mp, asPtr, asCopy, asCopy2)
	// Output:
	// Before changing values:
	// slice value: []string{"Look", "at", "all", "my", "values"}
	// map value: map[string]int32{"boxing":129354, "chestnut":127792, "heart-eyes":128525, "hibiscus":127802, "microphone":127908, "ocean":127754}
	// copy of struct value: pointer.ChangeThings{Int:777, Str:"Will I change?", Rune:127764}
	// pointer to struct value: &pointer.ChangeThings{Int:777, Str:"Will I change?", Rune:127764}
	// copy2 of struct value: pointer.ChangeThings{Int:777, Str:"Will I change?", Rune:127764}
	//
	// After changing values:
	// slice value: []string{"Dereferenced by `[]` operator. It acts just like `*` operator", "me 2", "me 3", "me 4", "me 5"}
	// map value: map[string]int32{"boxing":1, "chestnut":0, "heart-eyes":5, "hibiscus":3, "microphone":4, "ocean":4}
	// copy of struct value: pointer.ChangeThings{Int:777, Str:"Will I change?", Rune:127764}
	// pointer to struct value: &pointer.ChangeThings{Int:1002270, Str:"Dereferenced with the `.` operator, like `*` and `[]`.", Rune:127774}
	// copy2 of struct value: pointer.ChangeThings{Int:1002270, Str:"Dereferenced with the `.` operator, like `*` and `[]`.", Rune:127774}
	//
	// Actually remove all elements from collection types:
	// pointer: []string(nil)
	// pointer: map[string]int32(nil)
	// pointer: (*pointer.ChangeThings)(nil)
	// literal: pointer.ChangeThings{Int:0, Str:"", Rune:0}
	// literal: pointer.ChangeThings{Int:0, Str:"", Rune:0}
}
