package main

import(
    "bufio"
    "io"
    "fmt"
    "os"
)

type charCount struct{
    chCount int
    num int
    space int
    other int
}

func main(){
    file,err := os.Open("nihao.txt")
    if err != nil{
	fmt.Println("读取文件失败")
	return
    }

    defer file.Close()
    read := bufio.NewReader(file)

    var count = charCount{}

    for {
	str,err :=read.ReadString('\n')
	if err == io.EOF{
	    break
	}

	for i:=0;i<len(str);i++{
	    switch{
	    case str[i] >= 'a' && str[i]<='z' :
		fallthrough
	    case str[i] >= 'A' && str[i]<='Z' :
		count.chCount++
	    case str[i] == ' ' && str[i]=='\t' :
		count.space++
	    case str[i] == '0' && str[i]=='9' :
		count.num++
	    default:
		count.other++
	    }
	}
    }

    fmt.Println(count)

}
