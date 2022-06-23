package diffsquares

import "math"

func SquareOfSum(number int) int {
	sum := 0
	for n := 0; n <= number; n++ {
		sum += n
	}
	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(number int) int {
	sum := 0
	for n := 0; n <= number; n++ {
		sum += int(math.Pow(float64(n), 2))
	}
	return sum
}

func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}
