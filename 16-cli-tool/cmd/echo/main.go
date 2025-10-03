package main

import (
    "flag"
    "fmt"
    "strings"
)

func main(){
    upper := flag.Bool("upper", false, "uppercase output")
    sep := flag.String("sep", " ", "separator")
    flag.Parse()

    out := strings.Join(flag.Args(), *sep)
    if *upper { out = strings.ToUpper(out) }
    fmt.Println(out)
}
