package main

import "fmt"

func main() {
    m := map[string]int{"a":1}
    m["b"] = 2
    v, ok := m["c"]
    fmt.Println("c:", v, ok, "size:", len(m))
}
