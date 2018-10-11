package main

import "fmt"

func main() {
	// Untyped Constant
	const myFavLanguage = "Python"
	const sunRisesInTheEast = true

	// Multiple declaration
	const country, code = "Viet Nam", 81

	const (
		employeeId string = "E101"
		salary float64 = 50000.0
	)

	// Typed Constant
	const typedInt int = 1234
	const typedStr string = "Hi"

	fmt.Println(myFavLanguage, sunRisesInTheEast, country, code, employeeId, salary, typedInt, typedStr) // Python true Viet Nam 81 E101 50000 1234 Hi
}
