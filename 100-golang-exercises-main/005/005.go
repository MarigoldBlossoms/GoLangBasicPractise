package main

import (
	"bufio"
	"os"
	"strings"
)

/*
Exercise 005:

Create a module which has at least two methods:

  - ReadString: read a string from console input.
  - PrintString: return the previously read string in upper case.

The original exercise wires these together from a main.go file. For the
practice version we keep them as plain package-level functions so they can be
unit-tested without stdin.

Tip: run `go test ./...` from this folder.
*/

var input string

// SetInput is a small helper so the tests can seed the string without using
// real stdin. Your ReadString should write into the same storage that
// PrintString reads from.
func SetInput(s string) {
	input = s
}

// ReadString should read a line from stdin and store it for PrintString.
// (In the test we call SetInput instead, but the real program would call
// this from main.)
func ReadString() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
}

// PrintString should return the stored string converted to upper case.
func PrintString() string {
	return strings.ToUpper(input)
}
