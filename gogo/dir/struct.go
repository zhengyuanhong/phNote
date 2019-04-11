package main

import(
    "fmt"
)

type person struct{
    Name string
    Age int
}

func test(age int,name string) (p *person){
    p =new(person) 
    p.Name = name
    p.Age = age
    return 
}

func main(){
    
    p := test(18,"nihao")

    fmt.Println(*p)
}

