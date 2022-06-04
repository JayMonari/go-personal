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

// RangeValues shows that we can ignore the index using a range loop if we
// don't need it.
func RangeValues() {
	friends := []string{"Gabby", "Gorm", "Gunter"}
	// You cannot have unused variables in Go, but you can use _ to ignore them.
	for _, f := range friends {
		fmt.Println("friend:", f)
	}
}

// RangeIndexAndValues shows how to grab both the index and the value of each
// element in a slice.
func RangeIndexAndValues() {
	nums := []int{1, 2, 3, 4, 5}
	for i, n := range nums {
		fmt.Printf("index: %d, access value: %d, range value: %d\n", i, nums[i], n)
		// This can also be written as: nums[i] = n * n
		nums[i] *= n
	}
	fmt.Println("nums:", nums)
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
		fmt.Println("index:", i, "rune:", r, "representation:", string(r))
	}
}

// RangeChannel shows that we can grab values from a channel until it is
// closed.
func RangeChannel() {
	ch := make(chan string, 6)
	ch <- "We can get"
	ch <- "values from a channel"
	ch <- "continuously."
	ch <- "Just make sure"
	ch <- "you close the channel"
	ch <- "at some time 😉"
	// NOTE(jay): Make sure the channel is closed!
	// You will be stuck in an infinite loop without it.
	close(ch)
	for val := range ch {
		fmt.Println(val)
	}
}

// RangeScopedValues shows that you get copies of the values of all of the
// collections (slice, string, map) that we can range over and changing those
// values does NOT change the collection.
func RangeScopedValues() {
	scopedSlice := []int{0, 1, 2, 3, 4}
	scopedString := "NOT changed"
	scopedMap := map[string]bool{"x": true, "y": true, "z": true}
	fmt.Println("Try to change by just the value")
	for _, num := range scopedSlice {
		b := num
		num = 9
		fmt.Println("before:", b, "after:", num)
		fmt.Println("Never changes:", scopedSlice)
	}
	fmt.Println("Same!", scopedSlice)
	for _, r := range scopedString {
		b := r
		r = 'X'
		fmt.Println("before:", b, "after:", r)
	}
	fmt.Println("Same!", scopedString)
	for _, val := range scopedMap {
		b := val
		val = false
		fmt.Println("before:", b, "after:", val)
	}
	fmt.Println("Same!", scopedMap)
	fmt.Println()
	fmt.Println("Change by index (by dereference)")
	for i, n := range scopedSlice {
		if n < 4 {
			scopedSlice[i] = 9
		}
	}
	fmt.Println("Changed!", scopedSlice)
	// NOTE(jay): Remember strings are immutable! You can't do this!
	// for i, _ := range scopedString {
	// 	scopedString[i] = byte('X')
	// 	scopedString[i] = 'X'
	// }
	byteSlice := []byte(scopedString)
	for i, b := range byteSlice {
		// if our byte is lowercase we change it.
		if b >= 'a' && b <= 'z' {
			byteSlice[i] = 'X'
		}
	}
	fmt.Println("Changed!", string(byteSlice))
	for key := range scopedMap {
		scopedMap[key] = false
	}
	fmt.Println("Changed!", scopedMap)
}
