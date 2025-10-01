package main

import "fmt"

type Number interface {
    ~int | ~int64 | ~float64
}

func Sum[T Number](xs ...T) T {
    var total T
    for _, v := range xs { total += v }
    return total
}

func main(){ fmt.Println(Sum(1,2,3), Sum(1.5,2.5)) }
