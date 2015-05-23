package main
import (

    "fmt"
)

//不能用容器对象接收多返回值。只能用多个变量，或者用 _ 忽略。
func do() (int, int) {
    return 1, 100
}
//多返回值可直接作为其他函数调用的实参
func add(x, y int) int {
    return x+y
}
func sum(n ...int) int {
    var x int
    for _, i := range n {
        x+=i
    }
    return x
}

//参数返回参数可看做与其他形参类似的局部变量，最后由return 隐式返回。
func add2(x, y int) (z int) {
    z=  x+y
    return
}
//命名的参数可被同名的局部变量遮蔽，此时需要显式返回。
func add3(x, y int) (z int) {
    {
        var z = x+y  //不在同一级别，引发 "z redeclared in this block" 错误。
        return z      //必须显式返回。
    }
}
//命名返回参数允许 defer 延迟调用，通过闭包读取和修改。
func add4(x, y int) (z int) {
    defer func() {
        z+=100
    }()
    z=x+y
    return
}

//显式 return 返回以前，会修改命名返回参数。
func add5(x, y int) (z int) {
    defer func() {
        fmt.Println(z)
    }()

    z=x+y
    return z+200

}
func main() {
    //s :=make([]int,2)
    // s =do()  // multiple-value do() in single-value context
    x, _ := do()
    fmt.Println(x)

    println(add(do()))
    println(sum(do()))
    fmt.Println(add2(do()))
    fmt.Println(add3(100, 200))
    fmt.Println(add4(1, 2))
    add5(1, 1000)
}
