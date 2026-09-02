package main

import "math"

/*
Exercise 006:

Compute, for each D in the input slice:

    Q = round( sqrt( (2 * C * D) / H ) )

with C = 50 and H = 30.

Example: Ex006([]int{100, 150, 180}) -> []int{18, 22, 24}

Tip: run `go test ./...` from this folder.
*/

// Ex006 should apply the formula to every value in d and return the rounded
// integer results.
func Ex006(d []int) []int {
	var outputs []int = make([]int, len(d))
	for x := range d {
		pdt := 2 * 50 * d[x] / 30
		sq := math.Sqrt(float64(pdt))

		outputs[x] = int(math.Round(sq))
	}

	return outputs
}
