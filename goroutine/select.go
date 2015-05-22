package main
import (

    "fmt"
    "os"
)
//如果需要同时处理多个 channel ,可以使用 select 语句，它可以随机选择一个可用的 channel 做收发操作
func main() {
    a, b := make(chan int, 3), make(chan int)
    go func() {
        v, ok, s := 0, false, ""
        for {
            select {     //随机选择可用的 channel 接受数据
            case v, ok = <-a:
                s="a"
            case v, ok = <-b:
                s="b"
            }

            if ok {
                fmt.Println(v, s)
            }else {
                os.Exit(0)
            }
        }
    }()

    //向 channel 中发送数据
    for i := 0; i<5; i++ {
        select {    //随机选择可用的 channel 发送数据
        case a <- i:
        case b <- i:

        }
    }
    close(a)
    select {}

    // 在循环中使用 select default  case 要小心 ，避免形成洪水
}
