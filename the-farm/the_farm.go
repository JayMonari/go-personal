package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct{ amt int }

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.amt)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(wf WeightFodder, cows int) (float64, error) {
	switch {
	case cows == 0:
		return 0.0, errors.New("Division by zero")
	case cows < 0:
		return 0.0, &SillyNephewError{amt: cows}
	}

	fodder, err := wf.FodderAmount()
	switch err {
	case ErrScaleMalfunction:
		fodder *= 2
	case nil:
		// noop
	default:
		return 0.0, err
	}
	switch {
	case fodder < 0.0:
		return 0.0, errors.New("Negative fodder")
	default:
		return fodder / float64(cows), nil
	}
}
