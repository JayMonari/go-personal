package main

import "math"

func main() {
}

func declareConst() {
	// const is used exactly the same as var with limitations on types.
	const stuck = "This variable can never be reassigned."
	// XXX: Compiler will not allow this.
	// stuck = "This won't work, stuck is constant!"
	const heartEyes = '😍'
	// Constant expressions will do math for you!
	const aritmatic = 600 / 3.421
	const alwaysTrue = true

	// However you **cannot** declare arrays, slices, maps, or structs constant.
	// XXX: Will give error: InvalidConstInit
	// const myArray = [2]string{"won't", "work"}
	// const mySlice = []string{"still", "doesn't", "work"}
	// const myMap = map[string]int{}
	// const me = struct{ name string }{name: "Jay"}
}

func untypedConst() {
	// const values do not have a type and therefore are very useful when you
	// don't want to have to do explicit casting.
	const untyped = 42
	// We don't care what this function does, we only care what it looks like
	// math.IsInf(float64, int)
	math.IsInf(untyped, untyped)
	// If we try this with typed int we have to cast it.
	var typed int = 42 // or typed := 42
	math.IsInf(float64(typed), typed)
}
