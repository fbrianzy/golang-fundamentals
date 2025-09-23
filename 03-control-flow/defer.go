package main

import "fmt"

func main() {
    fmt.Println("start")
    defer fmt.Println("deferred 1")
    defer fmt.Println("deferred 2")
    fmt.Println("end")
}
