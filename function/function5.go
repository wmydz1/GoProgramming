package main
import (

    "fmt"
)
//匿名函数可以赋值给变量，作为结构字段，或者在channel 里传送。
func main() {
    //function variable
    fn := func() {
        fmt.Println("Hello Go")
    }
    fn()
    //function collection
    fns := []func(x int) int{
        func(x int) int {return x +1},
        func(x int) int {return x+2},
    }
    fmt.Println(fns[0](100))
    //function as field
    d := struct {
        fn func() string
    }{
        fn:func() string {return "Hello,GO"},
    }
    fmt.Println(d.fn())
    //channel function
    fc := make(chan func() string, 2)
    fc <- func() string {return "Google Inc" }
    f := <-fc
    fmt.Println(f())

}
