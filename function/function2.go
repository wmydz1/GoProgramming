package main
import (

    "fmt"

)
//函数是第一类对象，可以作为参数传递。建议将复杂的签名定义为函数类型，以便于阅读。
func test(fn func() int) int {

    return fn()
}
type FormatFunc func(s string, x, y int) string//定义函数类型
func format(fn FormatFunc, s string, x, y int) string {
    return fn(s, x, y)
}
func main() {
    s1 := test(func() int {return 100})   //直接把匿名函数当做参数
    s2 := format(func(s string, x, y int) string {
        return fmt.Sprintf(s, x, y)
    }, "%d, %d", 10, 20)

    println(s1, s2)
    //有返回值的函数，必须有明确的终止语句，否则容易引发编译错误。
}
