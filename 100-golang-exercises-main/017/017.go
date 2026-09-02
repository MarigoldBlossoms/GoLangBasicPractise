package practice

import (
	"strconv"
	"strings"
)

/*
Exercise 017:

Compute the net amount of a bank account from a transaction log. Each entry
is a string "<OP> <AMOUNT>" where OP is "D" (deposit) or "W" (withdrawal).
Empty or malformed entries should be skipped.

Example:
  Ex017([]string{"D 300", "D 300", "W 200", "D 100"}) -> 500

Tip: run `go test ./...` from this folder.
*/

// Ex017 should return the net account balance.
func Ex017(transactions []string) int {

	sum := 0

	for _, val := range transactions {
		transaction := strings.Split(val, " ")
		if len(transaction) >= 2 {
			amount, _ := strconv.Atoi(transaction[1])
			switch transaction[0] {
			case "D":
				sum += amount

			case "W":
				sum -= amount

			default:
				continue
			}
		}
	}

	return sum
}
