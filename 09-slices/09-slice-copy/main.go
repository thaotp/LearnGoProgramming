package main

import "fmt"

func main() {
	/*
		The copy() function copies elements from one slice to another
		func copy(dst, src []T) int
	*/

	src := []string{"Sublime", "VSCode", "IntelliJ", "Eclipse"}
	dest := make([]string, 2)

	numElementsCopied := copy(dest, src)

	fmt.Println("src = ", src) // src =  [Sublime VSCode IntelliJ Eclipse]
	fmt.Println("dest = ", dest) // dest =  [Sublime VSCode]
	fmt.Println("Number of elements copied from src to dest = ", numElementsCopied) // Number of elements copied from src to dest =  2
}
