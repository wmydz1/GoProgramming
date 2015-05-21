package main
import (

    "fmt"
)
type Resource struct {
    id int

}
type User struct {
    //不能同时嵌入某一类型和其指针类型
    *Resource
    title string
}
func main() {
    // 匿名字段不过是一种语法糖,从根本上说,就是一个与成员类型同名 (不含包名) 的字段。 被匿名嵌入的可以是任何类型,当然也包括指针。
    //可以像普通字段那样访问匿名字段成员,编译器从外向内逐级查找所有层次的匿名字段, 直到发现目标或出错
    u := User{&Resource{10000}, "samchen"}
    fmt.Println(u)
    fmt.Println(u.Resource.id,u.title)

}
