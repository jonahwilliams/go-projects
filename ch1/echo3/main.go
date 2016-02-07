package main

import (
    "os"
    "strings"
    "fmt"
)

func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}