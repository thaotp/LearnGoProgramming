package main
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Print("Hello, World. \n")
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
}
