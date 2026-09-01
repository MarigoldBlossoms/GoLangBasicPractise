package main

import "fmt"

func myFunction(fname string) {
	x := 4

	fmt.Printf("%v Doe", fname)
}

func main() {
	myFunction("John")
}
