package main

import (
    "fmt"
)


var ch = make(chan int,1) 
var ch1 = make(chan int) 
var ch3 = make(chan int) 

func test(){
    for i:=0;i<10;i++{
	<-ch
	fmt.Println("我被挂起",i)
	ch1<-1
    }
}

func test1(){
    for i:=10;i<20;i++{
	<-ch1
	fmt.Println("2我被挂起",i)
	ch<-1
    }
    ch3<-1
}

func main(){
    go test()
    go test1()
    fmt.Println("程序结束",<-ch3)

}
