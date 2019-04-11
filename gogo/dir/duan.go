package main

import "fmt"

func main(){
    
    var b interface{}
    b = "string"
    r ,ok := b.(string)
    r = "ssss"
    if ok {
	fmt.Println("断言成功",r)
    }

}
