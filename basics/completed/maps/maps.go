package maps

import "fmt"

// https://en.wikipedia.org/wiki/Associative_array

// MapBasic shows how to make a new map and how to add and remove keys and
// values to it.
func MapBasic() {
	myMap := make(map[string]int, 4) // We can leave out the size too!
	myMap["key1"] = 100
	myMap["key2"] = 20
	myMap["key3"] = -3
	myMap["key4"] = 0xc0ffee
	fmt.Println("print myMap:", myMap)

	delete(myMap, "key1")
	delete(myMap, "key2")
	delete(myMap, "key3")
	delete(myMap, "key4")
	fmt.Println("empty after deleting myMap keys:", myMap)
}

// MapValueExists shows how to check if a value exists in a given map or not.
func MapValueExists() {
	inlineMap := map[string]string{"key1": "VALUE1", "key2": "VALUE2"}
	// Will return zero values if the value is not present.
	val1, present := inlineMap["key1"]
	fmt.Printf("value: %v is present? %t\n", val1, present)

	none, found := inlineMap["non-existent"]
	fmt.Printf("value: %v is present? %t\n", none, found)

	if val, exists := inlineMap["key2"]; exists {
		fmt.Println("value:", val, "is there do something with it.")
	}
	if val, ok := inlineMap["non-existent"]; ok {
		fmt.Println("This statement will never be reached!", val)
	}
}

// MapAsSet demonstrates how to create a `set` data structure -- an unordered
// collection of some type with very quick lookup and insertion -- the idomatic
// Go way.
//
// The time to use a `set` is when you don't care about the order of your
// values and you want to be able to add, delete, or get values instantly!
func MapAsSet() {
	type important string
	// struct{} takes up no space. When we see it we can think of a typed `nil`
	mySet := make(map[important]struct{})
	// We don't have to explicitly type cast these as "[ANYTHING HERE]" is an
	// untyped string and will type itself to what it needs to at runtime.
	mySet["fast lookup"] = struct{}{}
	mySet["unordered"] = struct{}{}
	mySet["fast insert"] = struct{}{}
	mySet["collection"] = struct{}{}

	// Simulate getting a string from some other function. We want to check if
	// our `set` has that key, so we type cast it `important(...)` and do
	// something with the information that our `set` has that key.
	gotFromOtherFunc := "fast insert"
	if _, found := mySet[important(gotFromOtherFunc)]; found {
		fmt.Println("We found the key we were looking for!")
		delete(mySet, important(gotFromOtherFunc))
	} else {
		fmt.Println("We didn't find that key, so let's do something else!")
		mySet[important(gotFromOtherFunc)] = struct{}{}
	}
}
