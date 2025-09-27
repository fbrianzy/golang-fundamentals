package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func lookup(id int) (string, error) {
    if id != 1 { return "", fmt.Errorf("lookup(%d): %w", id, ErrNotFound) }
    return "OK", nil
}

func main() {
    _, err := lookup(2)
    if err != nil {
        if errors.Is(err, ErrNotFound) {
            fmt.Println("handle not found:", err)
        }
    }
}
