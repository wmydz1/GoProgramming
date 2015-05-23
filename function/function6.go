package main
import (


    "fmt"
)
// 闭包复制的原对象的指针，这样很容易解释延迟引用的现象。
func delay() func(){
 x:=100
    fmt.Printf("x (%p) == %d\n",&x,x)
    return func(){fmt.Printf("x (%p) == %d\n",&x,x)}
}
func main(){
      f:=delay()
      f()
}
//在汇编层面，test 实际返回的是 FuncVal 对象，其中包含匿名函数的地址和闭包对象的指针。
//只需将返回地址对象保存到某个寄存器中，就可以让匿名函数获取相关闭包对象的指针。