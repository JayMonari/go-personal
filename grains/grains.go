package grains

import (
	"errors"
	"math"
)

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("number must be between 1 and 64")
	}
	return uint64(math.Pow(2, float64(n-1))), nil
}

func Total() uint64 {
	return 18446744073709551615
}
