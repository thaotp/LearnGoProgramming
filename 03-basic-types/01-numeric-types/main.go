package main

import "fmt"

func main() {
  var myInt8 int8 = 27

  /*
    When you don't declare any type explicitly, the type inferred is `int`
    (The default type for integers)
  */
  var myInt = 3113

  var myUint uint = 256

  var myHexNumber = 0xF2  // Use prefix '0x' or '0X' for declaring hexadecimal numbers
  var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers

  var myFloat32 float32 = 3.14
  var myFloat = 6.12 // // Type inferred as `float64`

  fmt.Printf("%d, %d, %d, %#x, %#o %f %f\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber, myFloat32, myFloat) // 27, 3113, 256, 0xf2, 034 3.140000 6.120000
}
