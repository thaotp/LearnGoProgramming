package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)
    var bits []int32
    var max int32 = 0
    for n != 0 { 
      if(n%2 == 0){
        bits = nil
      }else{
        bits = append(bits, 1)
      }
      if(max < int32(len(bits))){
        max = int32(len(bits))
      }
      n = n/2
    }
  
  fmt.Printf("%d", max)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
