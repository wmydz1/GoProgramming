package main
import (

    "fmt"
)
func main(){
    d1:=make(chan int)
    d2:=make(chan int,3)

    d2<-1
    //len()返回未被读取缓冲元素的数量，cap()返回缓冲区的大小
    fmt.Println(len(d1),cap(d1))
    fmt.Println(len(d2),cap(d2))
}
// 向clsoed channel 发送数据会引发 panic 错误,nil channel 无论收发都会被阻塞
