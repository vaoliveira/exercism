package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (err *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %v cows", err.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction && fodderAmount >= 0 {
		return 2 * (fodderAmount / float64(cows)), nil
	}
	if fodderAmount < 0 && (err == ErrScaleMalfunction || err == nil) {
		return 0, errors.New("negative fodder")
	}
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if cows < 0 {
		return 0, &SillyNephewError{cows: cows}
	}
	if err != nil {
		return 0, err
	}
	return fodderAmount / float64(cows), nil
}
