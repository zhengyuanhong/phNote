package main

import (
    "fmt"
    "time"
    "sync"
)

var (
    mymap = make(map[int]int,10)
    lock sync.Mutex
)

func test(n int){
    res :=1
    for i:=0;i<=n;i++{
	res+=1
    }
    lock.Lock()
    mymap[n] = res
    lock.Unlock()
}

func main(){

    for i:=0;i<=10;i++{
	go test(i)
    }

    time.Sleep(time.Second*5)

    lock.Lock()
    for i,v := range mymap{
	fmt.Println(i,"===",v)
    }
    lock.Unlock()
}
