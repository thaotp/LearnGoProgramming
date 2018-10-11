package main

import "fmt"

func main() {
	// Appending to a slice that has enough capacity to accommodate new elements
	slice1 := make([]string, 3, 10)
	copy(slice1, []string{"C", "C++", "Java"})

	slice2 := append(slice1, "Python", "Ruby", "Go")

	fmt.Printf("slice1 = %v, len = %d, cap = %d\n", slice1, len(slice1), cap(slice1)) // slice1 = [C C++ Java], len = 3, cap = 10
	fmt.Printf("slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2)) // slice2 = [C C++ Java Python Ruby Go], len = 6, cap = 10

	slice1[0] = "C#"
	fmt.Println("\nslice1 = ", slice1) // slice1 =  [C# C++ Java]
	fmt.Println("slice2 = ", slice2) // slice2 =  [C# C++ Java Python Ruby Go]

	/*
		In this case, no new array would be allocated, and the elements would be added to the same underlying array.
	*/
}
