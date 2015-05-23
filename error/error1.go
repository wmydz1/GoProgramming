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

func test2() {
    //延迟调用引发的错误，可以被后续延迟调用捕获，但是仅最后一个错误被捕获
    defer func() {
        fmt.Println(recover())
    }()

    defer func() {
        panic("defer panic")
    }()
    panic("test painc")
}
//捕获函数 recover 只有在延迟调用内，直接调用才会终止错误，否则总是返回 nil 。 任何未捕获的错误都会沿调用堆栈向外传递。
func test3() {
    defer recover()  //无效
    defer fmt.Println(recover()) //无效
    defer func() {
        func() {
            println("defer inner")
            recover() //无效
        }()
    }()
    //   panic("panic test3")


}
func except() {
    fmt.Println(" from Go")
    recover()
}
//使用延迟匿名函数或者下面的都是有效的。
func test4(){
    defer  except()
    panic("test4")
}
//如果需要保护代码片段，可以将代码块重构成匿名函数，如此可以确保后续代码被执行
func test5(x ,y int){
    var z int
    func(){
        defer func(){
          if recover() !=nil{
              z=0
          }
        }()
        z=x/y
        return

    }()
    fmt.Println("x / y = ",z)
}
func main() {
   // test()
  //  test2()
   // test3()
    test4()
    test5(100,2)
}