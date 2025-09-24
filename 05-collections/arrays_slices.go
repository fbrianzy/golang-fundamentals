package main

import "fmt"

func main() {
    arr := [3]int{1,2,3}
    sl := arr[:]
    sl = append(sl, 4)
    fmt.Println("arr:", arr, "slice:", sl, "len/cap:", len(sl), cap(sl))
}
