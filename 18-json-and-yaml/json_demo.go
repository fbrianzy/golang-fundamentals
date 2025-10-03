package main

import (
    "encoding/json"
    "fmt"
)

type User struct{ Name string; Age int }

func main(){
    b, _ := json.Marshal(User{"A", 21})
    fmt.Println(string(b))

    var u User
    _ = json.Unmarshal([]byte(`{"Name":"B","Age":20}`), &u)
    fmt.Printf("%+v\n", u)
}
