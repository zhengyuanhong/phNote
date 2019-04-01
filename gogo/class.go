package main

import "fmt"

type person struct{
    Name string
    age int    
}

func(p *person) Setx(age int){
    p.age = age
}



type xiao struct{
    Name string
    person
}

type big struct{
    Name string
    person
}

func (b *big) Setb(name string){
    b.Name = name
}

type usb interface{
    start()
}

type computer struct{
    
}

func (c computer) aa(u usb){
    u.start()
}

func(b *big) start(){
    fmt.Println("big 实现了接口")
}

type A interface{
    say()
}

type B interface{
    hello()
}


type C struct{
    
}

func (c *C) say(){
    fmt.Println("C实现了A接口的方法")
}

func (c *C) hello(){
    fmt.Println("C实现了B接口的方法")
}

func main(){
    b :=&big{}
    b.Setb("big") 
    fmt.Println("b.Name=",b.Name)
    
    b.Setx(4)
    fmt.Println("b.age=",b.age)
    x := new(xiao)
    x.Setx(12)
    fmt.Println("x.age=",x.age)
    
    CC := computer{}
    CC.aa(b)

    c := &C{}
    var AA A = c
    AA.say()
    var BB B = c
    BB.hello()
    
    c.say()
    c.hello()
}
