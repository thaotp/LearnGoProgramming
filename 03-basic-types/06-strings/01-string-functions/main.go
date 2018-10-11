package main

import "strings"
import "fmt"

func main() {

    // Here's a sample of the functions available in
    // `strings`. Since these are functions from the
    // package, not methods on the string object itself,
    // we need pass the string in question as the first
    // argument to the function. You can find more
    // functions in the [`strings`](http://golang.org/pkg/strings/)
    // package docs.
    fmt.Println("Contains:  ", strings.Contains("test", "es")) // Contains:   true
    fmt.Println("Count:     ", strings.Count("test", "t")) // Count:      2
    fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te")) // HasPrefix:  true
    fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st")) // HasSuffix:  true
    fmt.Println("Index:     ", strings.Index("test", "e")) // Index:      1
    fmt.Println("Join:      ", strings.Join([]string{"a", "b"}, "-")) // Join:       a-b
    fmt.Println("Repeat:    ", strings.Repeat("a", 5)) // Repeat:     aaaaa
    fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", -1)) // Replace:    f00
    fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", 1)) // Replace:    f0o
    fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-")) // Split:      [a b c d e]
    fmt.Println("ToLower:   ", strings.ToLower("TEST")) // ToLower:    test
    fmt.Println("ToUpper:   ", strings.ToUpper("test")) // ToUpper:    TEST
    fmt.Println()

    // Not part of `strings`, but worth mentioning here, are
    // the mechanisms for getting the length of a string in
    // bytes and getting a byte by index.
    fmt.Println("Len: ", len("hello")) // Len:  5
    fmt.Println("Char:", "hello"[1]) // Char: 101
}
