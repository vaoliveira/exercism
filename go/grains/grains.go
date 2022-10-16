package grains

import "errors"

// Using bitwise operation, it is possible to shift left to quickly obtain the value of a certain square
func Square(number int) (uint64, error) {
	if number > 64 || number < 1 {
		return 0, errors.New("Invalid chess square")
	}
	return 1 << (number - 1), nil
}

// Knowing that the grains in square N is the sum of the grains in N-1 squares minus 1, we can shift left to get the value
// of a theoretical 65th square and subtract one to obtain the sum of all 64 squares
func Total() uint64 {
	return 1<<64 - 1
}
