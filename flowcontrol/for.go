package main
import (


    "fmt"
)

//支持三种循环方式
func  main(){
    s :="Golang"
    for i,n:=0,len(s); i<n; i++ {
        fmt.Printf("%s ",s[i:i+1])
    }

    l :=5
    for l>0{
        fmt.Println(l)
        l--
    }
//    for{
//        fmt.Println("loop forever")
//    }

  // 不要期望编译器能理解你的想法，在初始化语句中计算全部结果是个好主意。

    data :="abcd"
    for i,n:=0,len(data) ;i<n ;i++  {
         fmt.Println(i,data[i])
    }

}