package main

import (
	"fmt"
	"time"
)

func main() {
	SwitchBasic()
	SwitchMultiple()
	SwitchType()
	SwitchNoValue()
	SwitchFallthrough()
}

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

func SwitchType() {
	findType := func(i interface{}) {
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
	findType(true)
	findType(8)
	findType([]string{"some", "strings"})
	findType(struct{}{})
}

func SwitchNoValue() {
	t := time.Now().Month()
	switch {
	case t < time.February || t == time.December:
		fmt.Println("Winter is here.")
	case t <= time.May:
		fmt.Println("Looks like it's Spring!")
	case t <= time.August:
		fmt.Println("Must have some Summer flare.")
	case t <= time.November:
		fmt.Println("Autumn is in the air.")
	}
}

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
