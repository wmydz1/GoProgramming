package main
import (

    "errors"
    "fmt"
)
//除了用 panic 引发中断性错误外，还可以返回 error 类型的错误对象来表示函数的调用状态。
//标准库 error.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象。通过判断错误对象实例来确定错误的类型

var ErrdivByZero = errors.New("divion by zero")
func div(x, y int) (int, error) {
    if y==0 {
        return 0, ErrdivByZero
    }
    return x/y, nil
}
func main() {
    switch z, err := div(100, 0); err{
        case nil:
        fmt.Println(z)
        case ErrdivByZero:
        panic(err)
    }
    //如何区别使用 panic 和error 两种方式？ 惯例是： 在包内部使用 panic ,对外 API 使用 error 的返回值。

}
