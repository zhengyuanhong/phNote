package main
import "fmt"
import "sort"
import "reflect"

type Hero struct{
    Age int 
    Name string
}

type HeroSort []Hero

func (h HeroSort) Len()int{
    return len(h)
}

func (h HeroSort) Less(i,j int) bool{
    return h[i].Age>h[j].Age 
} 

func (h HeroSort) Swap(i,j int) {
    h[i] ,h[j] = h[j] ,h[i]
}


type H struct{
    Name string
}

type M []map[string]*H

func main(){
    var herosort HeroSort

    for i :=0 ;i<10;i++{
	hero := Hero{
	    Name:"zheng"+string(i), 
	    Age: 23+i,	
	}
	herosort = append(herosort,hero)
    }

    for k,v := range herosort{
	fmt.Println(v.Name,"++++",v.Age,"-----",k) 
    }
    sort.Sort(herosort)    

    for k,v := range herosort{
	fmt.Println(v.Name,"++++",v.Age,"-----",k) 
    }

   qie := make([]int,3)

    Ma := make(M,10)
    for i:=0;i<len(Ma);i++{
	Ma[i] = make(map[string]*H)
	Ma[i]["gogo"] = &H{
	    Name:"go"+string(i),	
	}
	fmt.Println(Ma[i]) 
    }
    fmt.Printf("Ma的类型%v",reflect.TypeOf(Ma))    
    fmt.Printf("qie的类型%v",reflect.TypeOf(qie))    
}
