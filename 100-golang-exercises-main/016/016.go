package main

import (
	"strconv"
	"strings"
)

/*
Exercise 016:

Accept a comma-separated sequence of numbers and keep only the odd ones,
preserving their order.

Example:
  Ex016("1,2,3,4,5,6,7,8,9") -> "1,3,5,7,9"

Tip: run `go test ./...` from this folder.
*/

// Ex016 should return only odd numbers from the comma-separated input.
func Ex016(input string) string {
	splitInput := strings.Split(input, ",")
	var output []string

	for _, val := range splitInput {
		if num, _ := strconv.Atoi(val); num%2 != 0 {
			output = append(output, val)
		}
	}

	return strings.Join(output, ",")
}
