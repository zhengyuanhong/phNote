package main

import "fmt"

type brid interface{
    flying()
    say()
}

type monkey struct{
    Name string
}

func (this *monkey) climbing(){
    fmt.Println("会跳的 = ",this.Name)
}

type litterMonkey struct{
    monkey
}

func (this * litterMonkey) flying(){
    fmt.Println("会飞 = ",this.Name)
}

func main(){
    var litter = litterMonkey{
	monkey{
	    Name:"go go",	
	},
    }
    litter.climbing()
    litter.flying()
}

