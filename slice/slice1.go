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
    s:=data[2:4]
    fmt.Println(s)
    s[0]+=100
    s[1]+=100
    fmt.Println(data)
    //可直接创建 slice 对象,自动分配底层数组

    // 通过初始化表达式构造,可使用索引号。
    s1:=[]int{0,1,2,3,8:100}
    fmt.Println(s1,len(s1),cap(s1))

    /// 使⽤ make 创建,指定 len 和 cap 值。
    s2:=make([]int,6,8)
    fmt.Println(s2,len(s2),cap(s2))
    // 省略 cap,相当于 cap = len。
    s3:=make([]int,6)
    fmt.Println(s3,len(s3),cap(s3))


}