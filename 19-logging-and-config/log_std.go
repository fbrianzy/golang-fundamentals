package main

import (
    "log"
    "os"
)

func main(){
    log.SetOutput(os.Stdout)
    log.SetPrefix("[APP] ")
    log.Println("hello")
}
