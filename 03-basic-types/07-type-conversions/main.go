package main

import "fmt"

func main() {
	// Type Conversion
	var a int64 = 4
	var b int = int(a)  // Explicit Type Conversion

	var c float64 = 6.5

	// Explicit Type Conversion
	var result = float64(b) + c  // Works

	fmt.Println(result) // 10.5


	// ==================


	// The general syntax for converting a value v to a type T is T(v)
	var myInt int = 65
	var myUint uint = uint(myInt)
	var myFloat float64 = float64(myInt)

	fmt.Println(myInt, myUint, myFloat) // 65 65 65
}
