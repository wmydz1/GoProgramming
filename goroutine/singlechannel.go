package main
import (

)
func main(){
    c:=make(chan int ,3)
    var send chan<- int =c  //send-only
    var recv <-chan int =c  //read-only

    send<-1
    <-recv


    //不能将单向的 channel 转换为普通的 channel
//    _:=(chan int)(send)
}
