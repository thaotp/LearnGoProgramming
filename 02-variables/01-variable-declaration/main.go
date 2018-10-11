package main

import "fmt"

func main() {
	// Declaring Variables
	var myStr string = "Golang"
	var myInt int = 26
	var myFloat float64 = 30.15
	fmt.Println(myStr, myInt, myFloat) // Golang 26 30.15

	// Multiple Declarations
	var (
		employeeId int = 1
		firstName, lastName string = "Tran", "Phuoc Thao"
	)
	fmt.Println(employeeId, firstName, lastName) // 1 Tran Phuoc Thao

	// Short variable declaration syntax
	name := "Tran Phuoc Thao"
	age, salary, isProgrammer := 26, 10000, true

	fmt.Println(name, age, salary, isProgrammer) // 26 10000 true
}