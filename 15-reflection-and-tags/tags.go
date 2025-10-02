package main

import (
    "fmt"
    "reflect"
)

type User struct {
    ID   int    `json:"id" db:"id"`
    Name string `json:"name" db:"name"`
}

func main(){
    t := reflect.TypeOf(User{})
    for i:=0;i<t.NumField();i++{
        f := t.Field(i)
        fmt.Println(f.Name, "json=", f.Tag.Get("json"), "db=", f.Tag.Get("db"))
    }
}
