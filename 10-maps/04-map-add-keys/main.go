package main

import "fmt"

func main() {
	// Initializing a map
	var tinderMatch = make(map[string]string)

	// Adding keys to a map
	tinderMatch["A"] = "A"
	tinderMatch["B"] = "B"
	tinderMatch["C"] = "C"

	fmt.Println(tinderMatch)

	/*
	  Adding a key that already exists will simply override
	  the existing key with the new value
	*/
	tinderMatch["C"] = "D"
	fmt.Println(tinderMatch)
}
