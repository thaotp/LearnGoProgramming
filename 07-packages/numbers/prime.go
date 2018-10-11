package numbers

import "math"

// Checks if a number is prime or not
func IsPrime(num int) bool {
  for i := 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
    if num%i == 0 {
      return false
    }
	}
  return num > 1
}

func IsPrime2(num int) bool{
  if num <= 1 {
    return false
  }
  if num <= 3 {
    return true
  }
  if num%2==0 || num%3==0 {
    return false
  }
  for i:=5; i*i <= num; i+=6{
    if num%i == 0 || num % (i + 2)==0 {
      return false
    }
  }
  return true
}