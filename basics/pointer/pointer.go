package pointer

import "fmt"

// PassByValue passes an int by value. When an integer is passed by value it is
// copied into the function meaning whatever we do to this `intValue` it will
// never change the original because it's a copy.
func PassByValue(intValue int) {
	intValue = 100
}

// PassByReference passes a pointer to the function. This value is an address,
// like your street address. If we derefence it `*intPointer` we can change the
// original value it was pointing to.
func PassByReference(intPointer *int) {
	*intPointer = 100
}

// PassMoreByReferences shows you that you can pass all of the other primitive
// data types and derefence '*' them and change their values.
func PassMoreByReferences(sPtr *string, bPtr *bool, rPtr *rune, fPtr *float64) {
	*sPtr = "Dereferenced and changed"
	*bPtr = true
	*rPtr = '🤡'
	*fPtr = 3.14159
}

type ChangeThings struct {
	Int  int
	Str  string
	Rune rune
}

// PassCollections shows what happens when we try to dereference slices, maps,
// and structs. Structs are special in that they are NOT a pointer like slices
// and maps and therefore need to have a pointer passed in if we want to change
// their inner values.
func PassCollections(slice []string, mp map[string]rune, ctCopy ChangeThings,
	ctPtr *ChangeThings) {

	// With fmt.Printf("%p, %p, %p, %p", slice, mp, ctCopy, ctPtr) you can see
	// the address of all of the arguments, we don't do it here because every
	// time the function executes the address will change! 🤯 and therefore the
	// test would, Never! pass.
	for i := range slice {
		if i == 0 {
			slice[i] = "Dereferenced by `[]` operator. It acts just like `*` operator"
		} else {
			slice[i] = fmt.Sprintf("me %d", i+1)
		}
	}
	for k, v := range mp {
		mp[k] = v % 7 // Dereferenced again with `[]` operator
	}
	ctCopy.Int = 0xF4B1E
	ctCopy.Str = "Changing a copies value doesn't work on the original."
	ctCopy.Rune = '🌞'
	ctPtr.Int = 0xF4B1E
	ctPtr.Str = "Dereferenced with the `.` operator, like `*` and `[]`."
	ctPtr.Rune = '🌞'
	// These lines do nothing, as the function makes a copy of everything it
	// receives. What we really wanted was to release the original collection
	// types to garbage collection. That must be done where it is instantiated.
	// Not when it is passed to a function as a copy.
	slice = nil
	mp = nil
	ctPtr = nil
	// XXX: won't work
	// ctCopy = nil
	//
	// This makes it more obvious that we have don't have a pointer. A pointer
	// can always be set to nil. We set it to a zero valued ChangeThings struct.
	ctCopy = ChangeThings{}
}
