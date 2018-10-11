package main

import "fmt"

func main() {
	// Iterating over an array using range form of for loop
	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for index, value := range daysOfWeek {
		fmt.Printf("Day %d of week = %s\n", index, value)
    // Day 0 of week = Mon
    // Day 1 of week = Tue
    // Day 2 of week = Wed
    // Day 3 of week = Thu
    // Day 4 of week = Fri
    // Day 5 of week = Sat
    // Day 6 of week = Sun
	}

	// Finding the sum of an array
	a := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for _, value := range a {
		sum = sum + value
	}

	fmt.Printf("Sum of all the elements in array %v = %f\n", a, sum) // Sum of all the elements in array [3.5 7.2 4.8 9.5] = 25.000000
}
