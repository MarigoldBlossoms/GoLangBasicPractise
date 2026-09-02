package main

/*
Exercise 003:

With a given integral number n, write a program to generate a map that contains
(i, i*i) such that i is an integral number between 1 and n (both included).

Example: Ex003(8) -> map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]

Tip: run `go test ./...` from this folder.
*/

// Ex003 should return a map where keys are 1..n and values are key*key.
func Ex003(n int) map[int]int {
	if n < 0 {
		return nil
	}

	var outputs = make(map[int]int)

	for i := 1; i <= n; i++ {
		outputs[i] = (i * i)
	}

	return outputs
}
