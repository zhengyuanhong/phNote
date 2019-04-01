package main

import (
    "fmt"
    "bufio"
    "os"
    "io"
    "io/ioutil"
)

func main(){
    /**
  if createFile("aa.txt","aa") {
	fmt.Println("写入成功")
    }
    line := openFile("nihao.txt")
   importFile("aa.txt","nihao.txt") 
    fmt.Print(line)
*/
    res , _ :=pathExit("aa.txt")
    if res {
	fmt.Println("文件存在")
    }
}

func createFile(str string,write string) bool{

    file,err := os.OpenFile(str,os.O_WRONLY|os.O_CREATE,0666)
    if err != nil{
	fmt.Println("文件创建失败")
    }

    defer file.Close()
    
    writer := bufio.NewWriter(file)

    for i:=0;i<8;i++{
	writer.WriteString(write + "\n")
    }
    writer.Flush()
    return true 
}

func openFile(str string) string{
    file , err := os.Open(str)
    if err != nil{
	fmt.Println("文件创建失败")
    }

    defer file.Close()

    read := bufio.NewReader(file)
    linestr := ""
    for {
	line,err1 := read.ReadString('\n')
	if err1 == io.EOF{
	    break
	}
	linestr += line
    }
    return linestr
}

func importFile(start,end string) bool{
    content,err := ioutil.ReadFile(start)    
    if err != nil{
	fmt.Println("文件读取失败")
	return false
    }

    err = ioutil.WriteFile(end,content,0666)
    if err != nil{
	fmt.Println("写入失败")
	return false
    }
    fmt.Println(start,"导入",end,"文件成功")
    return true
}

func pathExit(name string) (bool,error){
    _,err :=os.Stat(name)
    if err == nil{
	return true,nil
    }
    if os.IsNotExist(err){
	return false,nil
    }
    return false,nil
}
