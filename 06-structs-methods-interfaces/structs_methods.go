package main

import "fmt"

type Point struct{ X, Y float64 }

func (p Point) Length() float64 { return p.X*p.X + p.Y*p.Y }

func (p *Point) Move(dx, dy float64) {
    p.X += dx; p.Y += dy
}

func main() {
    p := Point{3,4}
    fmt.Println("len2:", p.Length())
    p.Move(1,-2)
    fmt.Println("after:", p)
}
