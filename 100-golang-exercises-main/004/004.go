package main

import (
	"strconv"
	"strings"
)

/*
Exercise 004:

Write a program which accepts a sequence of comma-separated numbers and
generates a slice out of them.

Example:
  Ex004("12, 23, 34, 45") -> []int{12, 23, 34, 45}

Tip: run `go test ./...` from this folder.
*/

// Ex004 should parse a comma-separated list of integers (whitespace is allowed
// around the numbers) and return them as []int.
func Ex004(input string) []int {
	var splitString = strings.Split(input, ",")
	var outputInts []int
	for i := 0; i < len(splitString); i++ {
		num, _ := strconv.Atoi(strings.TrimSpace(splitString[i]))
		outputInts = append(outputInts, num)
	}

	return outputInts
}
