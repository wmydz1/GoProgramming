package main
import (

    "fmt"
)
//接⼝是一个或多个方法签名的集合,任何类型的方法集中只要拥有与之对应的全部方法, 就表示它 "实现" 了该接口,无须在该类型上显式添加接口声明。
/*
• 接口命名习惯以 er 结尾,结构体。只有⽅法签名,没有实现。
• 接口没有数据字段。
• 可在接口中嵌入其他接口。
• 类型可实现多个接口。
*/

type Stringer interface {
    String() string
}
type Printer interface {
    Stringer //接口嵌入
    Print()
}
type User struct {
    id int
    name string
}
func (user User) String() string{
    return  fmt.Sprintf("user %d,%s",user.id,user.name)
}
func (use User) Print() {
    fmt.Println(use.String())
}
func main(){
    // *User ⽅方法集包含 String、Print。
    var t Printer =&User{100,"samchen"}
    t.Print()
}
