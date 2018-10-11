package main

import "fmt"

func main() {
	var (
		firstName, lastName string
		age int
		salary float64
		isConfirmed bool
	)

	fmt.Printf("firstName: %s, lastName: %s, age: %d, salary: %f, isConfirmed: %t\n",
		firstName, lastName, age, salary, isConfirmed) // firstName: , lastName: , age: 0, salary: 0.000000, isConfirmed: false

}