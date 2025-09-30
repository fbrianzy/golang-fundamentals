package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Msg struct { Text string `json:"text"` }

func main(){
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        _ = json.NewEncoder(w).Encode(Msg{"Hello, API!"})
    })
    log.Println("listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
