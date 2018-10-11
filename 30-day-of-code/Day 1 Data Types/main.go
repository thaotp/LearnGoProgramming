package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

  var i uint64 = 4
  var d float64 = 4.0
  var s string = "HackerRank "

  scanner := bufio.NewReader(os.Stdin)

  var (
      firstLine uint64
      secondLine float64
      thirdLine string
  )
  // Read and save an integer, double, and String to your variables.
  fmt.Scan(&firstLine)
  fmt.Scan(&secondLine)

  thirdLine, _ = scanner.ReadString('\n')
  // Print the sum of both integer variables on a new line.
  sum := i + firstLine
  fmt.Println(sum)
  // Print the sum of the double variables on a new line.
  sumDub := d + secondLine
  fmt.Printf("%.1f \n", sumDub)
  // Concatenate and print the String variables on a new line
  // The 's' variable above should be printed first.
  fmt.Println(s + "" + thirdLine)
}
