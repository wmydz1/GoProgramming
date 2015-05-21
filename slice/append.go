package main
import (

    "fmt"
)
func main(){
    //向 slice 尾部添加数据,返回新的 slice 对象
    s:=make([]int,0,5)
    fmt.Printf("%p\n",&s)
    s2:=append(s,1)
    fmt.Printf("%p\n",&s2)
    fmt.Println(s,s2)

    //简单点说,就是在 array[slice.high] 写数据
    data:=[...]int{0,1,2,3,4,5,6,7,8,9}
    slice:=data[:3]
    fmt.Println(slice)
    slice2:=append(slice,100,200)
    fmt.Println(data)
    fmt.Println(slice2)

    //⼀旦超出原 slice.cap 限制,就会重新分配底层数组,即便原数组并未填满
    data2:=[...]int{0,1,2,3,4,5,6,7,8,9}
    out:=data2[:3:4]
    fmt.Println(out)
    out2:=append(out,100,200)
    fmt.Println(out2)
    fmt.Println(data2)
    fmt.Println(&data2[0],&out[0],&out2[0])
    //从输出结果可以看出,s2 重新分配了底层数组,并复制数据。如果只追加⼀个值,正好 没超过 s.cap 限制,那么就不会重新分配数组。
    //大批量添加数据时,建议一次性分配 len 足够⼤的 slice,然后⽤索引号进⾏操作。还有,及时释放不再使⽤的 slice 对象,
    //避免持有过期数组,造成 GC ⽆法回收
}
