package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "math"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    N := int32(NTemp)
    if math.Mod(float64(N), 2) == 1 {
        fmt.Println("Weird")
    } else {
        if N < 6 && N > 1 {
            fmt.Println("Not Weird")
        }
        if N < 21 && N > 5 {
            fmt.Println("Weird")
        }
        if N > 20 {
            fmt.Println("Not Weird")
        }
    }
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
