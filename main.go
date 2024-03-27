package main

import (
	"errors"
	"math"
)

var (
	ErrNegativeInput = errors.New("input is negative")
	ErrOverflow      = errors.New("factorial result overflow")
)

func Factorial(n int) (uint64, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}

	result := uint64(1)
	for i := 1; i <= n; i++ {
		result *= uint64(i)
		// Check for overflow
		if result == 0 {
			return 0, ErrOverflow
		}
	}
	return result, nil
}

func RecursiveFactorial(n int) (uint64, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}
	if n == 0 {
		return 1, nil
	}
	fact, err := RecursiveFactorial(n - 1)
	if err != nil {
		return 0, err
	}
	if fact > math.MaxUint64/uint64(n) {
		return 0, ErrOverflow
	}
	return fact * uint64(n), nil
}
