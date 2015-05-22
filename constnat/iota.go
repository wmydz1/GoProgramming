package main
import (

    "fmt"

)
//关键字 iota 定义常量组中从0 开始按行计数的自增枚举值
const (
    Sunday = iota  //0
    Monday // 1
    Tuesday  // 2
    Wednesday // 3
    Thursday // 4
    Friday  // 5
    Saturday // 6

)
const (
    _ = iota                   //iota = 0
    KB int64 = 1<<(10 * iota)  //iota = 1
    MB
    GB
    TB
)
//在同一个常量组中，可以提供多个 iota ，他们各自增长
const (
    A, B = iota, iota<<10  //0 ,0 <<10
    C, D // 1,1 <<10
)
//如果 iota 自增被打断，需要显式恢复
const (

    a = iota  // 0
    b  // 1
    c = "c" // c
    d  // c  //与上一行相同
    e = iota // 4  显式恢复，计数包含了 c d 两行
    f       // 5

)
type  Color int
const (
    Black Color = iota
    Red
    Blue
)
func test(c Color) {

}
//可以通过自定义类型来实现枚举类型的限制
func main() {
    fmt.Println(Saturday)
    fmt.Println(KB)
    fmt.Println(D)
    fmt.Println(a, d, f)

    black := Black
    test(black)

    //    x:=1
    //    test(x) //cannot use x (type int) as type Color in argument to test
    test(1)  //常量会被编译器自动转换
}
