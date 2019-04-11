package main

import (
        "fmt"
)

type People interface {
        Speak(string) (tak string,b int)
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string,a int) {
        if think == "bitch" {
                talk = "You are a good boy"
        } else {
                talk = "hi"
        }
	a = 5
        return 
}

func main() {
        var peo People = &Stduent{}
        think := "bitch"
        fmt.Println(peo.Speak(think))
}
