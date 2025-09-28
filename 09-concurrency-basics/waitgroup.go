package main

import (
    "fmt"
    "sync"
)

func main(){
    var wg sync.WaitGroup
    for i:=0;i<3;i++{
        wg.Add(1)
        go func(id int){
            defer wg.Done()
            fmt.Println("work", id)
        }(i)
    }
    wg.Wait()
    fmt.Println("all done")
}
