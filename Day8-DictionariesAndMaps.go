package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
  
  var phones map[string]string
  phones = make(map[string]string)
  
  for i := 0; i < n*2; i++ {
    if(i < n){
      var key string
      var value string
      fmt.Scanln(&key, &value)
      phones[key] = value 
    }else{
      var name string
      fmt.Scanln(&name)
      if name == "" {
        break
      }
      if value, ok := phones[name]; ok {
        fmt.Printf("%s=%s\n", name, value)
      }else{
        fmt.Println("Not found")
      }
    }
  }
}
