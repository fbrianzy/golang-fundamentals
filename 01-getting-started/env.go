package main

import (
    "fmt"
    "os"
)

func main() {
    key := "PATH"
    fmt.Printf("%s=%s\n", key, os.Getenv(key))
}
