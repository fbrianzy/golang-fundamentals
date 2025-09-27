package main

import (
    "fmt"
    "github.com/fbrianzy/golang-fundamentals/07-packages-and-modules/libmath"
)

func main() {
    fmt.Println("Sum:", libmath.Sum(2,3))
}
