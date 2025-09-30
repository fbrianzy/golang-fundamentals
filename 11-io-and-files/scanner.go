package main

import (
    "bufio"
    "fmt"
    "strings"
)

func main(){
    s := "a\nb\nc\n"
    sc := bufio.NewScanner(strings.NewReader(s))
    for sc.Scan(){ fmt.Println("line:", sc.Text()) }
    if err := sc.Err(); err != nil { panic(err) }
}
