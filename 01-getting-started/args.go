package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Printf("Program: %s\n", os.Args[0])
    fmt.Printf("Args (%d): %v\n", len(os.Args)-1, os.Args[1:])
}
