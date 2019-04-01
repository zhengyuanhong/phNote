package main

import "fmt"
import "errors"

const SIZE = 4
type A [4]int
func main(){
	p :=person{}
	pp := p.newPerson("fengzhuang")
	fmt.Println("person.name=",pp.name)	
	//M :=make(map[string]string) 
//	var M = make(map[string]string,2)
	var M = map[string]string{
	    "c":"c",
	    "d":"d",
	}
	M["a"] = "a"
	M["b"] = "b"
	//map删除
	delete(M,"b")
	fmt.Println("Map = ",M)
	//map查询	
	if val,ress := M["a"]; ress{
	    fmt.Println("查询结果 = ",val,ress)
	}
	//map遍历
	for k,v :=range M {
	    fmt.Println("遍历结果 = ",k,v)
	}
	//map切片
	var MapS  []map[string]string
	MapS = make([]map[string]string,4)
	for k,v := range MapS{
	    if v == nil {
		MapS[k]  =  make(map[string]string)
		MapS[k]["nihao"] = "niaho"
	    }
	    fmt.Println("map切片遍历结果 = ",k,v)
	}
	    fmt.Println("MapS切片结果 = ",MapS)
	//普通切片
	var Pu []int
	Pu = make([]int,4)
	Pu = append(Pu,3)
	for k,v := range Pu{
	    fmt.Println("普通切片遍历结果 = ",k,v,Pu[4])
	}
	//数组指针
	arr :=[SIZE]int{1,2,3,4}
	f :=bibao(&arr)
	fmt.Println("main",*f())
	aa := A{1,1,1,1}
	aa.Print()
	//工厂类
	A :=&Aphone{"Azheng"}
	s,_ :=newInterface(A)
	fmt.Println("A=",s);
	//切片操作数组
	slice := arr[:]
	fmt.Println("slice=",slice[1:]);
	//new和结构体结合使用
	B :=new(Bphone)
	B.name = "Bzheng" 
	s1,i1 :=newInterface(B)
	s2 := []byte(s1)
	fmt.Println("B=",s2,i1);

	test()
	fmt.Println("正常执行...")

	str :="helloworld"
	strrr := str[2:]
	fmt.Println("字符串切片strrr = ",strrr)
	//异常处理	
	sb := []rune(str)
	sb[3] = '你'
	str = string(sb)
	fmt.Println("替换字符串中的字母 sb = ",str)

	err := readFile("")
	fmt.Println("正常执行...",err)
	//结构体方法
	AA := Aphone{}
	(&AA).name = "jieougtifangfa"
	fmt.Println("结构体方法...",AA.Say())
}

func bibao(p *[SIZE]int) (res func() *[SIZE]int){
	res = func() (c *[SIZE]int){
		for i := 0;i < len(*p); i++ {
			(*p)[i] = i+20
		}

		a:=*p
	fmt.Printf("%T\n",a)	
	fmt.Println("1 = ",p)
	fmt.Printf("p = %T\n",p)	
	fmt.Println("2 = ",a)
	fmt.Printf("*p = %T\n",*p)	

	dd := new(int)	
	cc :=&a
	fmt.Printf("dd : %T\n",dd)	
	fmt.Println("dd = ",dd)
	fmt.Printf("*dd : %T\n",*dd)	
	fmt.Println("*dd = ",*dd)

	fmt.Printf("cc : %T\n",cc)	
	fmt.Println("cc = ",cc)
	fmt.Printf("*cc : %T\n",*cc)	
	fmt.Println("*cc = ",*cc)
		for i,j := range a{
			a[i]= j+j
		}
		c = &a
		return
	}
	return
}

func (f A) Print(){
	var nih = "aaa"
	fmt.Println("*cc ============ ",f,nih)
}

type Phone interface{
	Call() string
}

type Aphone struct{
	name string
}

type Bphone struct{
	name string
}

func (A * Aphone) Say() string{
    return A.name
}

func(A *Aphone) Call() string{
	return (*A).name
}

func(B *Bphone) Call() string{
	return B.name
}

func newInterface(p Phone) (string,int){
    return p.Call() ,12
}

func test(){

 defer func(){
     err := recover()
     if err != nil{
	fmt.Println("错误异常",err) 
     }
     
 }()

    var (
	num1  = 23
	num2 = 0
    )
    res := num1/num2
    fmt.Println("错误处理",res)
}

func readFile(name string) (err error){
    if name == "text"{
	return nil
    } else {
	return errors.New("读取文件错误") 
    }
}

type person struct{
    name string
}

func (p person) newPerson(name string) *person{
    return &person{name:name}
}
