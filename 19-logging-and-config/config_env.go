package main

import (
    "fmt"
    "os"
)

func main(){
    port := os.Getenv("APP_PORT")
    if port == "" { port = "8080" }
    fmt.Println("port:", port)
}
