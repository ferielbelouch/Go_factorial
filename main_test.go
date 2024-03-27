package main

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		input    int
		expected uint64
		err      error
	}{
		{0, 1, nil},
		{1, 1, nil},
		{5, 120, nil},
		{-1, 0, ErrNegativeInput},
		{20, 0, ErrOverflow}, // 20! overflows uint64
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Factorial(%d)", tc.input), func(t *testing.T) {
			result, err := Factorial(tc.input)
			if err != tc.err {
				t.Errorf("Expected error: %v, got: %v", tc.err, err)
			}
			if result != tc.expected {
				t.Errorf("Expected result: %d, got: %d", tc.expected, result)
			}
		})
	}
}
