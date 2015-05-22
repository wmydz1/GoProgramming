package main
import (

    "fmt"
    "unsafe"
)
//常量必须是编译期可确定的数字.字符串，布尔值
const  x,y int =1,2 //多变量初始化
const  s ="Hello Golang" //类型判断

const (
    a,b=10,100
     c bool =false
)

const (
    l ="Java"
    j   //在常量组中，如果不提供类型和初始化值，那么将视作与上一个常量相同 j="Java"
)

const (
    //常量还可以是 len cap unsafe.Sizeof 等编译期可以确定结果的函数的返回值
    lang="Go"
    lang_len=len(lang)
    lang_size=unsafe.Sizeof(lang_len)
)
const (
    //如果常量足以储存初始化的值，那么不会引发溢出错误
    in byte =100
    in2 int =1e20  //overflows
)
func main(){
   const  x ="samchen"  //未使用的局部常量不会引发编译错误
    fmt.Println(j)
    fmt.Println(lang_size)
//    fmt.Println(in2)
}
