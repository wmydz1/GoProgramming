package main
import (

    "fmt"

)
type User struct {
    id int
    name string
}
func (user *User) String() string {
    return fmt.Sprintf("%d,%s", user.id, user.name)
}
func main() {
    var o interface{} = &User{1, "samchen"}
    if i, ok := o.(fmt.Stringer); ok {
        fmt.Println(i)
    }
    u:=o.(*User)
    // u := o.(User) // panic: interface is *main.User, not main.User
    fmt.Println(u)

    //可以用 switch 做批量类型判断 不支持 fallthroungh

    switch v :=o.(type){
        case nil:
        fmt.Println("nil")

        case fmt.Stringer :
        fmt.Println(v)

        case func() string:
        fmt.Println(v)

        case *User:
        fmt.Printf("%d,%s\n",v.id,v.name)

        default:
        fmt.Println("unknown")
    }
}
