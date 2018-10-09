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

    var arr [][]int32
    for i := 0; i < 6; i++ {
       arrRowTemp := strings.Split(readLine(reader), " ")
       var arrRow []int32
       for _, arrRowItem := range arrRowTemp {
         arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
         checkError(err)
         arrItem := int32(arrItemTemp)
         arrRow = append(arrRow, arrItem)
       }

      if len(arrRow) != int(6) {
         panic("Bad input")
      }

      arr = append(arr, arrRow)
    }
    var max int32 = -9*7;
    for z:=0; z < 4; z++{
      for q:=0; q < 4; q++{
        var sum int32 = 0
        for i:=z; i < z + 3; i++{
          for j:=q; j < q + 3; j++{
            if (i - z != 1) || (j - q == 1){
              sum += int32(arr[i][j])
            }
          }
        }
        if sum > max{
          max = sum
        }
      }
    }
    fmt.Println(max)
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
