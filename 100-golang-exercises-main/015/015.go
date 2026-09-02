package main

import "math"

/*
Exercise 015:

Compute a + aa + aaa + aaaa for a given digit a.

Example:
  Ex015(9) -> 11106  (9 + 99 + 999 + 9999)
  Ex015(1) -> 1234   (1 + 11 + 111 + 1111)

Tip: run `go test ./...` from this folder.
*/

// Ex015 should return the sum a + aa + aaa + aaaa.
func Ex015(a int) int {
	sum := 0

	for i := 0; i < 4; i++ {
		sum += IncrementDigits(a, i)
	}

	return sum
}

func IncrementDigits(digit, rounds int) int {
	if rounds <= 0 {
		return digit
	}

	return (digit * int(math.Pow(10, float64(rounds)))) + IncrementDigits(digit, rounds-1)
}
