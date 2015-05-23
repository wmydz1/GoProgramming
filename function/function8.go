package main
import (

    "fmt"
)
//多个 defer 注册，按FIOL次序执行。哪怕函数或某个延迟调用发生错误。这些调用旧会被执行。

func test8(x int) {
    defer fmt.Println("a")
    defer fmt.Println("b")
    defer func() {
        fmt.Println(100/x)  //异常未被捕获，逐步向外传递，最终终止进程。
    }()
    defer fmt.Println("c")
}
// 延迟调用参数在注册时求值或者复制，可用指针和闭包延迟读取
func test9(){
 x,y:=10,20
    defer func(i int){
      fmt.Println("defer:",i,y)  //y 闭包引用
    }(x)    //x 被复制

    x+=10
    y+=100
    fmt.Println("x =",x,"y = ",y)
}
func main() {
//    test8(0)
    test9()
}
