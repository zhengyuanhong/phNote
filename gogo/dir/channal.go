package main

import(
    "fmt"
)
type pp interface{
    say()string
}

type person struct{
    Name string
}

func (p *person) say()string{
    return p.Name
}

var (
    ch = make(chan []map[string]interface{},3)
)

func main(){
    str := make([]map[string]interface{},3)
    str[1] =make(map[string]interface{}) 
    var sttr string 
    sttr = "aaaaaad"
   str[1]["stt"] = person{
	Name:sttr, 
    }
    ch<-str
    res := <-ch
    res1 := res[1]["stt"]
    var p pp 
    res11,ok := res1.(person)
    if ok {
	fmt.Println("杜昂眼陈工")
    }

    p = &res11
    
    fmt.Println(p.say())
}
