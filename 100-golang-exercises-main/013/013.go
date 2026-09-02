package main

import (
	"strconv"
	"strings"
)

/*
Exercise 013:

Accept a sentence and return the number of letters and digits in the format:

    LETTERS x, DIGITS y

Example: Ex013("hello world! 123") -> "LETTERS 10, DIGITS 3"

Tip: run `go test ./...` from this folder.
*/

// Ex013 should count letters and digits and return them formatted as
// "LETTERS <n>, DIGITS <m>".
func Ex013(input string) string {

	input = strings.ToLower(input)
	letters, digits := 0, 0

	//iterate
	for _, char := range input {
		//is it a letter or number
		if (char >= 'a') && (char <= 'z') {
			letters++
		} else if (char >= '0') && (char <= '9') {
			digits++
		} else {
			continue
		}
	}

	return FormatOutput(letters, digits)
}

func FormatOutput(letters, digits int) string {
	letterCount := strconv.Itoa(letters)
	digitsCount := strconv.Itoa(digits)

	output := "LETTERS "
	output += letterCount
	output += ", DIGITS "
	output += digitsCount

	return output

}
