package main

import (
    "fmt"
    "os"
)

func main(){
    fname := "sample.txt"
    if err := os.WriteFile(fname, []byte("hello\n"), 0644); err != nil { panic(err) }
    b, err := os.ReadFile(fname); if err != nil { panic(err) }
    fmt.Println(string(b))
}
