package practice

import (
	"strconv"
	"strings"
)

/*
Exercise 012:

Find all numbers in [start, end) where every digit is even and return them
as a comma-separated string.

For the standard challenge call Ex012(100, 300) (the expected result is in
the test file).

Tip: run `go test ./...` from this folder.
*/

// Ex012 should return numbers in [start, end) whose digits are all even,
// joined by commas.
func Ex012(start, end int) string {
	var output []string

	for num := start; num <= end; num++ {
		numStr := strconv.Itoa(num)
		good := true
		for _, char := range numStr {
			if val := int(char); val%2 != 0 {
				good = false
			}
		}
		if good {
			output = append(output, strconv.Itoa(num))
		}
	}

	return strings.Join(output, ",")
}
