package main
import (

    "fmt"
)
/*
slice 并不是数组或数组指针。它通过内部指针和相关属性引⽤用数组⽚片段,以 实现变⻓长⽅方案

• 引⽤类型。但⾃身是结构体,值拷⻉传递。
• 属性 len 表⽰示可⽤用元素数量,读写操作不能超过该限制。 • 属性 cap 表⽰示最大扩张容量,不能超出数组限制。
• 如果 slice == nil,那么 len、cap 结果都等于 0
*/

func main() {
    data := [...]int{0, 1, 2, 3, 4, 5, 6,7,8,9}
    // [low : high : max]
    //len=high-low
    //cap=max-low
    slice := data[1:4:5]

    fmt.Println(slice)
    // 0-5 从0开始取6位 最大扩充容量8
    fmt.Println(data[:6:8])
    //从第5位以后取到末尾
    fmt.Println(data[5:])
    //取前面3位
    fmt.Println(data[:3])
    //取出全部元素
    fmt.Println(data[:])

    //读写操作实际⺫⽬目标是底层数组,只需注意索引号的差别
}