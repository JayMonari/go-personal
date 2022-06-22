package switches

import (
	"fmt"
	"time"
)

// SwitchBasic shows how to switch what logic to perform depending on a case
// and if that none of the criteria are met we can perform some default logic.
func SwitchBasic() {
	i := 0
	switch i {
	case 1:
		fmt.Println("i is one")
	case 2:
		fmt.Println("i is two")
	case 3:
		fmt.Println("i is three")
	default:
		fmt.Println("i does not have a matching case.")
	}
}

// SwitchMultiple shows that we can perform the same logic for multiple cases
// using the same `case` keyword by comma separating the values.
func SwitchMultiple() {
	month := time.August
	switch month {
	case time.December, time.January, time.February:
		fmt.Println("Winter is here.")
	case time.June, time.July, time.August:
		fmt.Println("Must have some Summer flare.")
	case time.October, time.November, time.September:
		fmt.Println("Autumn is in the air.")
	case time.March, time.April, time.May:
		fmt.Println("Looks like it's Spring!")
	}
}

// SwitchType shows us that we can do type assertions using switch statements!
// This is particularly useful when getting JSON data with no idea what's
// inside.
func SwitchType(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Printf("You seem like an %T-eresting type.\n", t)
	case bool:
		fmt.Printf("You %T! I knew it was you all along.\n", t)
	case []string:
		fmt.Printf("Hey, hey. Save me a slice! %T\n", t)
	default:
		fmt.Printf("We've never seen a %T like this.\n", t)
	}
}

// SwitchNoValue shows that you don't have to give a value to the `switch`
// statement, but instead perform true or false (bool) assertions on a given
// value.
func SwitchNoValue() {
	// This will get the current month. So this test may fail, see if you can't
	// update the test with the correct string to make it pass! 😁
	t := time.Now().Month()
	switch {
	case t <= time.February || t == time.December:
		fmt.Println("Winter is here.")
	case t <= time.May:
		fmt.Println("Looks like it's Spring!")
	case t <= time.August:
		fmt.Println("Must have some Summer flare.")
	case t <= time.November:
		fmt.Println("Autumn is in the air.")
	}
}

// SwitchFallthrough shows off the **very** rarely used `fallthrough` keyword
// in Go. If you're using `fallthrough` there's probably a better solution to
// your problem.
func SwitchFallthrough() {
	switch "three" {
	case "three":
		fmt.Println("Floor number three")
		fallthrough
	case "two":
		fmt.Println("Floor number two")
		fallthrough
	case "one":
		fmt.Println("Floor number one")
		fallthrough
	default:
		fmt.Println("Now arriving bottom floor.")
	}
}
