package main
import (

    "fmt"
)
//一次可以定义多个变量
var x, y, z int
var s, n = "abc ", 123

var (
    a int
    b float32
)
func test() (int, string) {
    return 1, "Golang"
}
func main() {
    n, s := 0x1234, "Hello Go"
    fmt.Println(x, n, s)
    //多变量赋值时，先计算所有相关值，然后从左往右依次赋值
    data, i := [3]int{1, 2, 3}, 0
    i, data[i]=2, 100
    fmt.Println(i, data) // (i = 0) -> (i = 2), (data[0] = 100)

    //特殊只写 "_" ,用于忽略值占位

    _, result := test()
    fmt.Println(result)
    //注意编译器会将未使用的局部变量当做错误处理，全局变量没问题
    y, result := 100, "Hello Golang" //重新赋值，与前 result 在同一层次的代码块中，且有新的变量被定义

    fmt.Println(y, result)
    //通常多返回值的 err 会被重复使用

    for i := 0; i<5; i++ {
        result, z := 100, 10 //定义新的同名变量：不在同一层次的代码块
        fmt.Println(result, z)
    }
}
