package main
import (

    "fmt"
)
//值类型,赋值和传参会复制全部内容。可⽤ "_" 定义补位字段,⽀持指向⾃身类型的指针成员
type Node struct {
    _ int
    id int
    data *byte
    next *Node
}
//顺序初始化必须包含全部字段,否则会出错。
type User struct {
    name string
    age int
}
//顺序初始化必须包含全部字段,否则会出错。
//too few values in struct initializer
//x  u2 := User{"tom"}
// fmt.Println(u2)
func main() {
    u1 := User{"samchen", 100}

    fmt.Println(u1)
    //支持匿名结构,可用作结构成员或定义变量。
}
