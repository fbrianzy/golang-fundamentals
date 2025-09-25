package main

import "fmt"

type Logger struct{ Prefix string }
func (l Logger) Log(msg string){ fmt.Println(l.Prefix + msg) }

type Service struct {
    Logger // embedded
    Name string
}

func main() {
    svc := Service{Logger: Logger{"[SVC] "}, Name: "Auth"}
    svc.Log("starting " + svc.Name)
}
