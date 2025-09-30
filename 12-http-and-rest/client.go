package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Msg struct { Text string `json:"text"` }

func main(){
    resp, err := http.Get("http://localhost:8080/hello")
    if err != nil { panic(err) }
    defer resp.Body.Close()
    var m Msg
    _ = json.NewDecoder(resp.Body).Decode(&m)
    fmt.Println(m.Text)
}
