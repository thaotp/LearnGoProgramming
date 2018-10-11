package main

import "fmt"

func main() {
	s := [][]string{
		{"James Smith", "United States"},
		{"Maria Gracia", "England"},
		{"Sarah Johnson", "France"},
	}

	fmt.Println("Slice s = ", s) // Slice s =  [[James Smith United States] [Maria Gracia England] [Sarah Johnson France]]
	fmt.Println("length = ", len(s)) // length =  3
	fmt.Println("capacity = ", cap(s)) // capacity =  3
}
