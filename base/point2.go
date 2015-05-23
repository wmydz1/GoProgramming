package main
import (


    "unsafe"
    "fmt"
)
//将 Pointer 转换 uintptr,可以变相实现指针运算。
func main() {
    d := struct {
        s string
        x int
    }{"samchen", 1000}

    p := uintptr(unsafe.Pointer(&d))  //*struct->Pointer->uintptr
    p+=unsafe.Offsetof(d.x)       //uintptr+offset

    p2 := unsafe.Pointer(p)  //unintptr->Pointer
    px := (*int)(p2)  //Pointer ->*int
    *px=200 //d.x =200
    fmt.Println(d)
    //GC 会把 uintptr 当做普通对象，它无法阻止 "关联的对象被回收"
}
