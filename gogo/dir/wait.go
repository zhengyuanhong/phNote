package main

import "fmt"
import "sync"

var wg = sync.WaitGroup{}

func main(){
    wg.Add(100)
    go func(){
	for i:=0;i<100;i++{
	    defer wg.Done()
	    fmt.Println("go func",i)
	}
    }()

    wg.Add(100)
    go func(){
	for i:=0;i<100;i++{
	    defer wg.Done()
	    fmt.Println("-----",i)
	}

    }()

    wg.Wait()
}
