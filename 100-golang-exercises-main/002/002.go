package main

import (
	"errors"
)

/*
Exercise 002:

Write a program which can compute the factorial of a given number.

  - Ex002(8) -> 40320, nil
  - Ex002(0) -> 1,     nil   (by definition)
  - Ex002(-3) -> 0,    error (negative input is not allowed)

Tip: run `go test ./...` from this folder.
*/

// Ex002 should return n! as a uint64 (or an error for negative input).
func Ex002(n int) (uint64, error) {
	if n < 0 {
		return 0, errors.New("negative input is allowed")
	}

	return Factorial(n), nil
}

func Factorial(num int) uint64 {
	if num <= 1 {
		return 1
	}

	return uint64(num) * Factorial(num-1)
}
