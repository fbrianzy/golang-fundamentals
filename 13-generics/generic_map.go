package main

import "fmt"

func Map[T any, R any](xs []T, f func(T) R) []R {
    out := make([]R, len(xs))
    for i, v := range xs { out[i] = f(v) }
    return out
}

func main(){
    xs := []int{1,2,3}
    ys := Map(xs, func(v int) int { return v*v })
    fmt.Println(ys)
}
