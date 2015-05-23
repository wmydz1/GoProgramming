package main
import (

    "fmt"
    "unsafe"
)
//支持指针类型 *T,指针的指针 **T,以及包含包名的前缀 *<package> .T
/*
  默认值是 nill ，没有默认 NULL
  操作符 & 取变量的地址， * 透过指针访问目标
  不支持指针运算，不支持 ->运算，直接用 . 访问目标成员
*/
type data struct {
    a int
}
//返回的局部变量指针式安全的，编译器会根据需要将其分配到 GC Head 上
func test(){
    x :=100
    return &x  //使用 runtime.new 分配 x 内存，但是关联时，也可能直接分配到目标栈上。
}
func main() {
    var d = data{1234}
    var p *data
    p=&d
    //直接用指针访问目标成员对象，无需转换
    fmt.Println(p, p.a)
    //不能对指针做加减法

    //可以在 unsafe.Pointer 和任意类型的指针间进行转换
    x :="0x12345678"
    point :=unsafe.Pointer(&x) //*int -> Pointer
    n:=(*[4]byte)(point) //Pointer ->*[4]byte

    for i:=0;i<len(n) ;i++  {
        fmt.Println(n[i])
    }








}
