package main
import (
  "fmt"
)


func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
  var n int
  fmt.Scanln(&n)

  // arr := make([]int, n)
  for i := 0; i < n; i++ {
    var num int
    fmt.Scanf("%d", &num)
    if isPrime(num){
      fmt.Println("Prime")
    }else{
      fmt.Println("Not prime")
    }
    
  }
}
func isPrime(n int) bool {
  if n <= 1 {
    return false
  }
    
  if n <= 3 {
    return true
  }
  if n%2==0 || n%3==0 {
    return false
  }
  for i:=5; i*i <= n; i+=6{
    if n%i == 0 || n % (i + 2)==0 {
      return false
    }
  }
  return true
} 
