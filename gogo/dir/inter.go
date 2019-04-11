package main

import "fmt"

type A struct{
    name string
}

type B interface{
    say() interface{}
}

func (a *A) say() interface{}{
    return a.name
}

func main(){
    a:= &A{
	name:"nihao",
    }

    res := a.say()
    fmt.Println(res)
    res1,ok := res.(string)
    if ok {
	fmt.Println("断言成功")
    }
    fmt.Println(res1)
}
