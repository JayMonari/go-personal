package main

import "fmt"

func main() {
	ForLoop()
	WhileLoop()
	ForeverLoop()
	ContinueBreakLoop()
}

func ForLoop() {
	// We initialize a variable `i := 0`
	// We end the statement with a semicolon `;`
	// So that we can start a new statement.
	// We declare what condition we want to stop at `i <= 10;`
	// This means if `i` is less than or equal to 10 keep looping.
	// If it's 11 or higher. Stop!
	// We declare what values will be change on the next loop `i++`
	// This also could be read as `i = i + 1`
	for i := 0; i <= 10; i++ {
		fmt.Println("For loop i:", i)
	}
}

func WhileLoop() {
	// In other languages there is something called a while loop.
	// It only stops on a condition. Like the middle part of a `for` loop.
	i := 1
	for i <= 128 {
		fmt.Println("While loop i:", i)
		// This can also be written as
		// i = i + i
		i += i
	}
}

func ForeverLoop() {
	// A `for` loop with no terminating condition, will never stop!
	// Unless we force ourselves out of the loop by `break`ing it.
	fmt.Println("How many print statements do you want?\nAnother One,")
	for {
		fmt.Println("And Another One,")
		// XXX: This loop will not stop if you remove this. Try it out!
		// Use CTRL+C to stop the program.
		break
	}
	fmt.Println("We the best.")
}

func ContinueBreakLoop() {
	// We can also skip while looping if we don't want to perform any action.
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println("ContinueBreak loop i:", i)
		if i == 25 {
			break
		}
	}
}
