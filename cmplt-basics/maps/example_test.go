package maps_test

import "basics/maps"

func ExampleBasic() {
	maps.Basic()
	// Output:
	// print myMap: map[key1:100 key2:20 key3:-3 key4:12648430]
	// empty after deleting myMap keys: map[]
}

func ExampleValueExists() {
	maps.ValueExists()
	// Output:
	// value: VALUE1 is present? true
	// value:  is present? false
	// value: VALUE2 is there do something with it.
}

func ExampleAsSet() {
	maps.AsSet()
	// Output:
	// We found the key we were looking for!
}
