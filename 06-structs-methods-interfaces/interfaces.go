package main

import "fmt"

type Stringer interface{ String() string }

type User struct{ Name string }

func (u User) String() string { return "User(" + u.Name + ")" }

func printAll(xs []Stringer) {
    for _, x := range xs {
        fmt.Println(x.String())
    }
}

func main() {
    printAll([]Stringer{User{"A"}, User{"B"}})
}
