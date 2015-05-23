package main
import (
// 不支持嵌套(nested)重载(overload) 和默认参数(default parameter)
    "fmt"
)
/*
1 无需声明原型
2 支持不定长变参
3 支持多返回值
4 支持命名返回参数
5 支持匿名函数和闭包
*/
//使用关键字 func 定义函数，左大括号依旧不能另起一行
func test1(x, y int, s string) (int, string) {  //类型相同的相邻参数可以合并
    n := x+y
    return n, fmt.Sprintf(s, "lang")  //多返回值必须使用括号

}



func main() {
    _, s := test1(1, 2, "Go")
    fmt.Println(s)
}