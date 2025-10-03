package main

import "fmt"

func fanIn(chs ...<-chan int) <-chan int {
    out := make(chan int)
    done := make(chan struct{}, len(chs))
    for _, ch := range chs {
        go func(c <-chan int){
            for v := range c { out <- v }
            done <- struct{}{}
        }(ch)
    }
    go func(){
        for i:=0;i<len(chs);i++{ <-done }
        close(out)
    }()
    return out
}

func main(){
    a := make(chan int); b := make(chan int)
    go func(){ for i:=1;i<=3;i++{ a<-i }; close(a) }()
    go func(){ for i:=100;i<=102;i++{ b<-i }; close(b) }()
    for v := range fanIn(a,b){ fmt.Println(v) }
}
