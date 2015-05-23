package main
import (


    "fmt"

)
func main() {
    //类似迭代器的操作，返回索引或者键值
    s := "abc"
    for i := range s {  //忽略2nd value ,支持 string/array/slice/map
        fmt.Println(s[i])
    }
    for _, c := range s {
        fmt.Println(c) //忽略 index
    }
    m := map[string]int{"a":1, "b":100}
    for k, v := range m {  //返回 key value
        fmt.Println(k, v)
    }
    //注意 range 会复制对象
    a := [3]int{0, 1, 2}
    for i, v := range a {  //index 和 value 都是从复制品里取出的
        if i==0 {       //修改之前，先修改原数组
            a[1], a[2] =999, 99
            fmt.Println(a)  //确认修改有效输出 [0 999 99]
        }
        a[i]=v+100  //使用复制品取出的 value 修改原数组
    }
    fmt.Println(a)  // 输出 [100 101 102]

    //建议修改引用类型，其底层数据不会复制

    arr :=[]int{1,2,3,4,5}
    for i ,v :=range arr{  //复制 struct slice (pointer,len ,cap)
        if i==0{
            arr=arr[:3]
            arr[2]=100
        }
        fmt.Println(i,v)
    }

    //另外两种引用类型 map channel 是指针包装，而不像 slice 是 struct 。


}
