package main

import "fmt"

func main() {
  //Enter your code here. Read input from STDIN. Print output to STDOUT
  var n int
  fmt.Scanln(&n)

  arr := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scanf("%d", &arr[i])
  }
  
  var numberOfSwaps int
  for i := 0; i < n; i++ {
    swapped := false

    for j := 0; j < n-1; j++ {
      if arr[j] > arr[j+1] {
        arr[j], arr[j+1] = arr[j+1], arr[j]
        numberOfSwaps++
        swapped = true
      }
    }

    if swapped == false {
      break
    }
  }

  fmt.Printf("Array is sorted in %d swaps.\n", numberOfSwaps)
  fmt.Printf("First Element: %d\n", arr[0])
  fmt.Printf("Last Element: %d\n", arr[n-1])
}
