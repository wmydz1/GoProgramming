package main
import (

    "fmt"
)
//可以省略表达式和括号
//支持初始化语句,可定义代码局部变量
// 代码块左括号必须在条件表达式尾部
func main() {
    x := 0
    if x<10 {
        fmt.Println(x)
    }
    if n := "Golang"; x <10 {
        fmt.Println(n[:2])  //初始化语句未必就是定义变量，比如 println("init") 是可以的。
    }else{
        fmt.Println("unkonwn")
    }
    // if 不支持 三元操作符 a>b ?a=bs
}
