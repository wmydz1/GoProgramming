package main
import (

    "fmt"
)
func test() {
    defer func() {
        if err := recover(); err!=nil {
            fmt.Println(err.(string), "from Go")  //将 interface{}转换为具体的类型
        }
    }()
    panic("painc error")
}

// 由于 panic 和 recover 参数类型为 interface{} ,因此可以抛出任何类型的对象
// func panic (v interface)
//func recover(v interface)
func main() {
    test()
}