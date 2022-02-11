package pointers

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
