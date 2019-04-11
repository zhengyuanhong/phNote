package main

import (
    "fmt"
    "io"
    "os"
    "bufio"
)

func main(){

    _,err := copyFile("nn.png","l.png")
    if err == nil{
	fmt.Println("复制成功")
    }

}

func copyFile(distName,srcName string) (written int64, err error){
	file,err := os.Open(srcName)
	if err != nil{
	    fmt.Println("读取失败")
	}
	
	defer file.Close()
	read := bufio.NewReader(file)
	
	file2,err := os.OpenFile(distName,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil{
	    fmt.Println("读取失败")
	}

	write := bufio.NewWriter(file2)
	defer file2.Close()
	
	return io.Copy(write,read)
}
