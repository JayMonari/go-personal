package collatzconjecture

import "errors"

// CollatzConjecture returns the total steps taken before the number reaches 1
// according to the collatz conjecture. If number < 1 an error is returned.
func CollatzConjecture(number int) (int, error) {
	if number < 1 {
		return 0, errors.New("number must be a positive integer")
	}

	steps := 0
	for number > 1 {
		if number%2 == 0 {
			number >>= 1
		} else {
			number = number*3 + 1
		}
		steps++
	}
	return steps, nil
}
