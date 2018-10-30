package main

import "fmt"

func main() {
	var employeeSalary = map[string]float64{
		"A":  78000.00,
		"B": 160000.50,
		"C": 85000.00,
	}

	var salary = employeeSalary["C"]
	fmt.Printf("C's Salary = %f\n", salary)

	// If a key doesn't exist in the map, we get the zero value of the value type
	salary = employeeSalary["B"]
	fmt.Printf("B's Salary = %f\n", salary)
}
