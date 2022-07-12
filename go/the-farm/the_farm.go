package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	err  error
	cows int
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction && fodderAmount >= 0 {
		return 2 * (fodderAmount / float64(cows)), nil
	} else if fodderAmount < 0 && (err == ErrScaleMalfunction || err == nil) {
		return 0, errors.New("negative fodder")
	} else if cows == 0 {
		return 0, errors.New("division by zero")
	} else if cows < 0 {
		value := fmt.Sprintf("silly nephew, there cannot be %v cows", cows)
		nephew := SillyNephewError{
			cows: cows,
			err:  errors.New(value),
		}
		return 0, nephew.err
	} else if err != nil {
		return 0, err
	}
	return fodderAmount / float64(cows), nil
}
