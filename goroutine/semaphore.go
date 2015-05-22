package main
import (

    "sync"
    "fmt"
)
func main() {
    wg := new(sync.WaitGroup)
    wg.Add(3)

    sem := make(chan int, 1)
    for i := 0; i<3; i++ {
        go func(id int) {
            defer wg.Done()
            sem <- 1     //向 sem 发送数据，阻塞或者成功

            for x:=0; x<3;x++  {
                fmt.Println(id,x)
            }

            <-sem //接受数据，使得其他阻塞的 goroutine 可以发送数据
        }(i)
    }
    wg.Wait()
}
