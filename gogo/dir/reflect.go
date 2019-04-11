package main

import(
    "fmt"
    "reflect"
)

func reflectTest(b interface{}){
    rTyp := reflect.TypeOf(b)
    fmt.Println("rTyp=",rTyp)
    rVal := reflect.ValueOf(b)
    fmt.Println("rval=",rVal)
}

func reflect2(b interface{}){
    rval := reflect.ValueOf(b)

    rval.Elem().SetInt(30)

}

func main(){
    var num int
    num = 45
    reflect2(&num)
    fmt.Println("num=",num)

    var strr string
    strr = "string"
    
    rtyp := reflect.ValueOf(&strr)

    str := rtyp.Elem().Len()

    fmt.Println("str=",str)
}
