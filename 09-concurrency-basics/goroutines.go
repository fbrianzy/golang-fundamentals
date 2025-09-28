package main

import (
    "fmt"
    "time"
)

func work(id int){
    fmt.Println("start", id)
    time.Sleep(50 * time.Millisecond)
    fmt.Println("done", id)
}

func main() {
    for i:=0;i<5;i++ { go work(i) }
    time.Sleep(200 * time.Millisecond)
}
