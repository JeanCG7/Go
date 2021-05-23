// Package collatzconjecture contains simple functions to calculate Collatz Conjecture
package collatzconjecture

import (
	"errors"
)

// CollatzConjecture returns the number of steps to reach 1, given an positive integer
func CollatzConjecture(input int) (int, error) {
	steps := 0
	if input <= 0 {
		return 0, errors.New("input must be a positive integer number")
	}

	for input != 1 {
		if input%2 == 0 {
			input = input / 2
		} else {
			input = input*3 + 1
		}
		steps = steps + 1
	}

	return steps, nil
}
