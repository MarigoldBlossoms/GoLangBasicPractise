package practice

import (
	"strings"
)

/*
Exercise 008:

Accept a comma-separated sequence of words and return the words in a
comma-separated sequence after sorting them alphabetically.

Example:
  Ex008("without,hello,bag,world") -> "bag,hello,without,world"

Tip: run `go test ./...` from this folder.
*/

// Ex008 should split, sort alphabetically and re-join the words with commas.
func Ex008(input string) string {
	stringSlice := strings.Split(input, ",")

	sorting := true
	for sorting {
		sorting = false
		for i := 0; i < len(stringSlice)-1; i++ {
			if stringSlice[i] > stringSlice[i+1] {
				stringSlice[i], stringSlice[i+1] = stringSlice[i+1], stringSlice[i]
				sorting = true
			}
		}
	}

	return strings.Join(stringSlice, ",")
}
