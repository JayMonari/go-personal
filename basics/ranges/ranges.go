package ranges

import "fmt"

// RangeIndex shows us how to grab the index of each element in a slice.
func RangeIndex() {
	uselessSlice := make([]struct{}, 10)
	for i := range uselessSlice {
		fmt.Println("index:", i)
	}
	// You can also inline this!
	// for i := range make([]struct{}, 10) {
	// 	fmt.Println("index:", i)
	// }
}

// RangeIndexAndValues shows how to grab both the index and the value of each
// element in a slice.
func RangeIndexAndValues() {
	nums := []int{1, 2, 3, 4, 5}
	for i, n := range nums {
		fmt.Printf("index: %d, access value: %d, range value: %d\n", i, nums[i], n)
		// This can also be written as:
		// nums[i] = n * n
		nums[i] *= n
	}
	fmt.Println("nums:", nums)
}

// RangeValues shows that we can ignore the index using a range loop if we
// don't need it.
func RangeValues() {
	friends := []string{"Gabby", "Gorm", "Gunter"}
	// You cannot have unused variables in Go, but you can use _ to ignore them.
	for _, f := range friends {
		fmt.Println("friend:", f)
	}
}

// RangeMap shows that we can loop through the entries of a map (key and value)
// using range. Maps are not ordered in Go!
func RangeMap() {
	isMarried := map[string]bool{"Gaph": true, "Gene": false, "Gable": false}
	for key, val := range isMarried {
		if val == true {
			fmt.Println(key, "is married.")
		} else {
			fmt.Println(key, "is not married.")
		}
	}
}

// RangeString shows that we can get the index and rune value of each character
// in a string.
func RangeString() {
	for i, r := range "gophergo.dev" {
		fmt.Println("index:", i, "rune:", string(r))
	}
}
