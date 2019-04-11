package main

import(
    "fmt"
)

type person struct{}

func (person) say(){
    fmt.Println("nihao")
}

func main(){
    var p person
    p.say()
}
