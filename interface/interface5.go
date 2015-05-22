package main
import (

    "fmt"
)
//超集接口可以转换成子集接口，反之出错
type Stringer interface {
    String() string
}
type Printer interface {
    String() string
    Print()
}
type User struct {
    id int
    name string
}

func (user *User) String() string{

    return fmt.Sprintf("%d,%v",user.id,user.name)
}
func (user *User) Print(){
    fmt.Println(user.String())
}
func main(){
    var o Printer =&User{10,"samchen"}
    var s Stringer =o
    fmt.Println(s.String())


}
