package main
import (

    "fmt"
    "time"
)
//使用 select 设置超时
func main() {
    w := make(chan bool)
    c := make(chan int, 2)

    go func() {
        select {
        case v := <-c:
            fmt.Println(v)
        case <-time.After(time.Second *3):
            fmt.Println("time out")

        }
        w<-true
    }()
     <-w
}