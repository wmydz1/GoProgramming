package main
import (

    "fmt"
)
//变参的本质就是 slice 。只能有一个，且必须是最后一个。
func change(s  string, n ...int) string {
    var x int
    for _, i := range n {
        x+=i
    }
    return fmt.Sprintf(s, x)
}
func main() {
    println("sum is %d", 1, 2, 3)
}
