package main
import (

    "fmt"
)
// 使用关键字 var 定义变量，自动初始化为零值。如果提供初始化值，可省略变量类型，由编译器自动判断。

var  x int
var  f float32 =1.8
var  s ="Golang"
//在函数内部，可使用 ":=" 方式定义变量
func main(){
    //注意检查，是定义新的局部变量，还是修改全局变量。该方式容易造成错误。
    x:="hello go"
    fmt.Println(x)
}
