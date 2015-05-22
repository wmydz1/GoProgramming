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
func main() {
    n, s := 0x1234, "Hello Go"
    fmt.Println(x, n, s)
    //多变量赋值时，先计算所有相关值，然后从左往右依次赋值
    data,i :=[3]int{1,2,3},0
    i,data[i]=2,100
    fmt.Println(i,data) // (i = 0) -> (i = 2), (data[0] = 100)

    
}
