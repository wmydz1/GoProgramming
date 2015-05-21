package main
import (

    "fmt"
)
type User struct {
    id int
    name string
}
type Manager struct {
    User
    title string
}

func main(){
    m:=Manager{User{100,"samchen"},"Google Inc"}
    fmt.Println(m)
    // 同类型拷贝。
    var u User =m.User
    fmt.Println(u)
}