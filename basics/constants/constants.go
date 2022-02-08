package constants

import (
	"fmt"
	"math"
)

// Stuck is an untyped string
const Stuck = "This variable can never be reassigned."

// Stuck = "This won't work, stuck is constant!"
// XXX: cannot assign to Stuck (untyped string constant "This variable can
// never be reassigned.")

// HeartEyes is an untyped rune
const HeartEyes = '😍'

// Arithmetic is an untyped float
const Arithmetic = 600 / 3.421

// AlwaysTrue is an untyped bool
const AlwaysTrue = true

// However you **cannot** declare arrays, slices, maps, or structs constant.
// XXX: (value of type [2]string) is not constant
// const myArray = [2]string{"won't", "work"}
// const mySlice = []string{"still", "doesn't", "work"}
// const myMap = map[string]int{}
// const me = struct{ name string }{name: "Jay"}

// UntypedConst shows that constants can have values that will be automatically
// converted to the necessary type that the function needs at runtime.
func UntypedConst() {
	// const values do not have a type and therefore are very useful when you
	// don't want to have to do explicit casting.
	const untyped = 42
	// We don't care what this function does, we only care what it looks like
	// math.IsInf(float64, int)
	fmt.Println(math.IsInf(untyped, untyped))
	// If we try this with typed int we have to cast it.
	var typed int = 42 // or typed := 42
	fmt.Println(math.IsInf(float64(typed), typed))
}
