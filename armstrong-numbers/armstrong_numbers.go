package armstrong

import (
	"math"
	"strconv"
)

// IsNumber returns whether a given number satisfies the Armstrong number
// conjecture, which is a number that is the sum of its own digits each rasied
// to the power of the number of digits.
func IsNumber(maybeAN int) bool {
	n := maybeAN
	exponent := float64(len(strconv.Itoa(maybeAN)))
	aSum := 0.0
	for n > 0 {
		aSum += math.Pow(float64(n % 10), exponent)
		n /= 10
	}
	return int(aSum) == maybeAN
}
