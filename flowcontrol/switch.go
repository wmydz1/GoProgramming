package main
import (

    "fmt"
)
//分支表达式可以是任意类型，不限于常量。可以省略 break ，默认自动终止。
func main() {

    x := []int{1, 2, 3}
    i := 2

    switch i{
        case x[1]:
        fmt.Println("a")
        case 1, 3:
        fmt.Println("b")
        default:
        fmt.Println("c")

    }

    d := 10
    switch d{
        case 10:
        fmt.Println("a")
        //如果要继续下一分支，可以使用 fallthrough
        fallthrough
        case 0:
        fmt.Println("b")
    }
    //省略表达式，可当 if ..else if if ..else
    switch {
        case x[1]>0:
        println("a")
        case x[2]<0:
        println("b")
        default:
        println("c")
    }

    //带初始化语句
    switch i := x[2]; {
        case i>0:
        println("ok")
        case i<0:
        println("no")
        default:
        println("unkown")
    }


}
