package main

import "strconv"

/*
Exercise 014:

Accept a sentence and count its upper- and lower-case letters.

Example:
  Ex014("Hello world!") -> "UPPER CASE 1, LOWER CASE 9"

Tip: run `go test ./...` from this folder.
*/

// Ex014 should return "UPPER CASE <u>, LOWER CASE <l>".
func Ex014(input string) string {

	upperCase, lowerCase := 0, 0

	for _, char := range input {
		if (char > 'a') && (char < 'z') {
			lowerCase++
		} else if (char > 'A') && (char < 'Z') {
			upperCase++
		} else {
			continue
		}
	}

	return FormatOutput(upperCase, lowerCase)
}

func FormatOutput(upperCase, lowerCase int) string {
	upperCaseString := strconv.Itoa(upperCase)
	lowerCaseString := strconv.Itoa(lowerCase)

	output := "UPPER CASE "
	output += upperCaseString
	output += ", LOWER CASE "
	output += lowerCaseString

	return output
}
