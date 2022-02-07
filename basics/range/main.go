package main

import "fmt"

func main() {
	RangeIndex()
	RangeValues()
	RangeIndexAndValues()
	RangeMap()
	RangeString()
}

func RangeIndex() {
	uselessSlice := make([]struct{}, 10)
	for i := range uselessSlice {
		fmt.Println("index:", i)
	}
	// You can also inline this! Though it doesn't ever make sense to use this.
	// for i := range make([]struct{}, 10) {
	// 	fmt.Println("index:", i)
	// }
}

// This is also called enumeration of a slice.
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

func RangeValues() {
	friends := []string{"Gabby", "Gorm", "Gunter"}
	// You cannot have unused variables in go, but you can use _ to ignore them.
	for _, f := range friends {
		fmt.Println("friend:", f)
	}
}

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

func RangeString() {
	for i, r := range "gophergo.dev" {
		fmt.Println("index:", i, "rune:", string(r))
	}
}
