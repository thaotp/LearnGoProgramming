package main

import "fmt"

func main() {
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	for index, number := range primeNumbers {
		fmt.Printf("PrimeNumber(%d) = %d\n", index+1, number)
	}
  // PrimeNumber(1) = 2
  // PrimeNumber(2) = 3
  // PrimeNumber(3) = 5
  // PrimeNumber(4) = 7
  // PrimeNumber(5) = 11
  // PrimeNumber(6) = 13
  // PrimeNumber(7) = 17
  // PrimeNumber(8) = 19
  // PrimeNumber(9) = 23
  // PrimeNumber(10) = 29
}
