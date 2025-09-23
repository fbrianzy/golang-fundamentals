package main

import "fmt"

func main() {
    // while-like
    i := 0
    for i < 3 {
        fmt.Println("i:", i)
        i++
    }
    // classic
    for j := 0; j < 3; j++ {
        fmt.Println("j:", j)
    }
    // range
    for idx, v := range []string{"a", "b", "c"} {
        fmt.Println(idx, v)
    }
}
