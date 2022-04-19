package for_loop

import "fmt"

// ForLoop shows how to initialize and run through a for loop. The structure of
// a for loop looks like `for initialize; stop condition; increment {`
func ForLoop() {
	// We initialize a variable `i := 0`
	// We end the statement with a semicolon `;`
	// So that we can start a new statement.
	// We declare what condition we want to stop at `i <= 10;`
	// This means if `i` is less than or equal to 10 keep looping.
	// If it's 11 or higher. Stop!
	// We increment the values that will change on the next loop `i++`
	// This also could be read as `i = i + 1`
	for i := 0; i <= 10; i++ {
		fmt.Println("For loop i:", i)
	}
}

// WhileLoop shows how to create a while loop in Go by just using a for loop. A
// while loop is just a for loop with only the condition to stop looping.
// `for stop condition {`
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

// ForeverLoop shows how to create a loop that will never stop running if there
// is no break condition inside of it.
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

// ContinueBreakLoop shows how we can skip logic downwards by using a continue
// statement and how we can break out of a loop on a certain condition.
func ContinueBreakLoop() {
	// We can also skip while looping if we don't want to perform any action.
	i := 0
	for {
		i++
		if i%2 == 0 {
			// This will make sure the logic below is not ran. It will make us go
			// back to the start of the loop where `i++` is.
			continue
		}
		fmt.Println("ContinueBreak loop i:", i)
		if i == 25 {
			// This will break us out from the forever loop
			break
		}
	}
}
