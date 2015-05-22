package main
import (

    "fmt"
)
//channel 是第一类对象，可以传参(内部实现为指针)或者作为结构成员
type Request struct {
    data []int
    ret chan int
}
func NewRequset(data ...int) *Request {
    return &Request{data, make(chan int, 1)}
}
func Process(req *Request) {
    x := 0
    for _, i := range req.data {
        x+=i
    }
    req.ret <- x
}
func main() {
    req := NewRequset(10, 20, 30)
    fmt.Println(req.data)
    Process(req)

    fmt.Println(<-req.ret)
}
