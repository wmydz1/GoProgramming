package main
import (

    "fmt"
)
type Tester interface {
    Do()
}
type FuncDo func()
func (self FuncDo) Do(){
    self()
}
//让函数直接实现接口
func main(){
   var t Tester =FuncDo(func(){
       fmt.Println("Hello Goalng")
   })
    t.Do()
}
