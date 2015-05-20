package main
import (

    "fmt"
)
/*
• 数组是值类型,赋值和传参会复制整个数组,而不是指针。
• 数组⻓长度必须是常量,且是类型的组成部分。[2]int 和 [3]int 是不同类型。
• 支持 "=="、"!=" 操作符,因为内存总是被初始化过的。
• 指针数组 [n]*T,数组指针 *[n]T。
*/
func main(){
    // 未初始化元素值为 0。
    a:=[3]int{1,2}
    // 通过初始化值确定数组⻓长度。
    b:=[...]int{1,2,3,4}
    // 使⽤用索引号初始化元素
    c:=[5]int{2:100,3:1000}

    d:=[...]struct{
        name string
        age uint8
    }{
        {"samchen",100},
        {"google",10},
    }
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)

    //支持多维数组
    e:=[2][3]int{{1,2,3},{4,5,6}}
    fmt.Println(e)

    // 第 2 纬度不能⽤用 "..."。
    f:=[...][2]int{{1,1},{2,2},{3,3}}
    fmt.Println(f)

    //值拷⻉贝⾏行为会造成性能问题,通常会建议使⽤用 slice,或数组指针。
    g:=[2]int{}
    fmt.Printf("x:%p\n",&g)
    test(g)
    fmt.Println(g)

    //内置函数 len 和 cap 都返回数组⻓长度 (元素数量)
    fmt.Println(len(g),cap(g))
}

func test(x [2]int){
  fmt.Printf("x:%p\n",&x)
    x[1]=10000
}

