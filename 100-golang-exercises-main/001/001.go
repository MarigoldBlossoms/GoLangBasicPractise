package main

import (
	"strconv"
	"strings"
)

/*
Exercise 001:

Write a program which will find all such numbers which are divisible by 7 but
are not a multiple of 5, between 2000 and 3200 (both included). The numbers
obtained should be printed in a comma-separated sequence on a single line.

Hint: consider using strconv and strings.Join.

Tip: run `go test ./...` from this folder.
*/

// Ex001 should return all integers in [low, high] that are divisible by 7
// but not by 5, joined by commas (e.g. "112,119,126,...").
func FactorOf(num, divisor int) bool {
	return (num%divisor == 0)
}

func Ex001(min, max int) string {
	var outputs []string
	for i := min; i <= max; i++ {
		if FactorOf(i, 7) && !(FactorOf(i, 5)) {
			outputs = append(outputs, strconv.Itoa(i))
		}
	}

	return strings.Join(outputs, ",")
}
