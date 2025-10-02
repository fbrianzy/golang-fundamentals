package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context){
    for {
        select{
        case <-ctx.Done():
            fmt.Println("cancelled:", ctx.Err())
            return
        case <-time.After(50*time.Millisecond):
            fmt.Println("working...")
        }
    }
}

func main(){
    ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
    defer cancel()
    go worker(ctx)
    time.Sleep(300*time.Millisecond)
}
