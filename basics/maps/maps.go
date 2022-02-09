package maps

import "fmt"

// https://en.wikipedia.org/wiki/Associative_array

// MapBasic shows how to make a new map and how to add and remove keys and
// values to it.
func MapBasic() {
	myMap := make(map[string]int)
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
