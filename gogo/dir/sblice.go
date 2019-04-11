package main

import "fmt"

type Sblice []interface{}
type Handle func(int,int)int

func(s *Sblice)Remove(value interface{})bool{
    for i,v := range *s {
	if isEqual(value,v){
	    *s = append((*s)[:i],(*s)[i+1:]...)
	    return true
	}
    }
    return false
}

func isEqual(value,v interface{})bool{
    if value == v {
	if res,ok :=  value.(int);ok{
	    fmt.Println("断言成功",res)
	}
	return true
    }else{
	return false
    }
}

func test(a,b int)int{
	fmt.Println("a+b",a+b)
	return a+b
}


func sum(a int,b int,f Handle)int{
    return f(a,b)
}

func main(){
    a  :=make(Sblice,5) 

    a[0] = 1
    a[2] = 2
    a[3] = 3
    a[1] = 4
    a[4] = 5

    if (&a).Remove(3){
	fmt.Println("success",a)
    }else{
	fmt.Println("fail")
    }

    fmt.Printf("value = %T",a)
    fmt.Printf("a[0] = %T",a[0])
    
    sum(1,3,test)
}
