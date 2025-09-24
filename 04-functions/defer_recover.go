package main

import "fmt"

func safeRun(f func()) (recovered any) {
    defer func() {
        if r := recover(); r != nil {
            recovered = r
        }
    }()
    f()
    return nil
}

func main() {
    r := safeRun(func() { panic("boom") })
    fmt.Println("recovered:", r)
}
