package main

import "fmt"

func main() {
	// Creating a slice using a slice literal
	var s = []int{3, 5, 7, 9, 11, 13, 17} // Creates an array, and returns a slice reference to the array

	// Short hand declaration
	t := []int{2, 4, 8, 16, 32, 64}

	fmt.Println("s = ", s) // s =  [3 5 7 9 11 13 17]
	fmt.Println("t = ", t) // t =  [2 4 8 16 32 64]
}
