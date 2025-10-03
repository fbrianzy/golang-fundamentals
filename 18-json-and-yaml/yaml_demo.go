package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
)

type Cfg struct {
    Port int    `yaml:"port"`
    Env  string `yaml:"env"`
}

func main(){
    data := []byte("port: 8080\nenv: dev\n")
    var c Cfg
    _ = yaml.Unmarshal(data, &c)
    fmt.Printf("%+v\n", c)
}
