package main
import (

    "fmt"
)

//异步方式通过缓冲区来决定是否阻塞。如果缓冲区已满，发送将被阻塞；缓冲区为空，接收将被阻塞
//通常情况下，异步channel可以减少排队阻塞，具备更高的效率。但应该考虑使用指针规避大对象拷贝
//将多个元素打包，减少缓存区的大小
func main() {
    data := make(chan int,3) //缓存区可以存储3个元素
    exit := make(chan bool)


    data<-1  //缓存取没有被塞满之前不会阻塞
    data<-2
    data<-3

    go func(){
       for d:=range data{ //缓存区不为空之前，不会阻塞
           fmt.Println(d)
       }
        exit<-true
    }()

    data<-4  //如果缓存区已经满了 ,阻塞
    data<-5

    close(data)
    <-exit
}
