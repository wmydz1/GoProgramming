package main
import (

    "fmt"
)
func main() {
    //引⽤类型,哈希表。键必须是支持相等运算符 (==、!=) 类型,比如 number、string、 pointer、array、struct,以及对应的 interface。
    //值可以是任意类型,没有限制
    m := map[int]struct {
        name string
        age int
    }{
        1:{"samchen", 100},
        2:{"Google Inc", 15},
    }
    fmt.Println(m[1].name)




}
