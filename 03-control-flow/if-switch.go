package main

import "fmt"

func grade(score int) string {
    switch {
    case score >= 90:
        return "A"
    case score >= 80:
        return "B"
    case score >= 70:
        return "C"
    default:
        return "D"
    }
}

func main() {
    for _, s := range []int{95, 83, 72, 60} {
        fmt.Println(s, "=>", grade(s))
    }
}
