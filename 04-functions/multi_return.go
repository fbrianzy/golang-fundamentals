package main

import "fmt"

func divmod(a, b int) (int, int, error) {
    if b == 0 {
        return 0, 0, fmt.Errorf("division by zero")
    }
    return a / b, a % b, nil
}

func main() {
    q, r, err := divmod(10, 3)
    fmt.Println(q, r, err)
}
