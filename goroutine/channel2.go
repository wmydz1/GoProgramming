package main

import (

    "fmt"
)
//引用channel是CSP模式的具体实现，用于多个 goroutine 通信。内部实现了同步确保并发安全
//默认为同步模式，需要发送和接受配对。否则会被阻塞，直到另一方准备好后被唤醒
func main(){
    data :=make(chan int)
    exit:=make(chan bool)

    go func(){
      for d :=range data{
          fmt.Println(d)
      }
        fmt.Println("recv over")
        exit<-true
    }()
    data<-1
    data<-2
    data<-3
    close(data)
    fmt.Println("send over")

    <-exit
}
