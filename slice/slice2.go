package main
import (

    "fmt"
)
func main() {
    //使⽤ make 动态创建 slice,避免了数组必须用常量做长度的麻烦。还可⽤指针直接访问
    //底层数组,退化成普通数组操作
    s:=[]int{0,1,2,3}
    p:=&s[2]
    *p+=100
    fmt.Println(s)

    //至于 [][]T,是指元素类型为 []T
     data:=[][]int{
         []int{1,2,3},
         []int{100,200},
         []int{11,22,33,44},
     }
    fmt.Println(data)
}