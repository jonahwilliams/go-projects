package main

import (
    "fmt"
    "os"
)

// Exercies 1.1 - Modify the program to also primt os.Args[0]
// Exercise 1.2 - Modify the program to print the index and value of each arg on a line
// Exercies 1.3 - TODO
func main() {
    for i, val := range os.Args {
        fmt.Println(i, val)  
    }
}