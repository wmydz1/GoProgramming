package main
import (

    "sync"
    "runtime"
)
//和协程yield类似，Gosched 让出底层线程，将goroutine暂停，放回队列
//等待下次调用被执行
func main(){
    wg:=new(sync.WaitGroup)
    wg.Add(2)

    go func(){
       defer  wg.Done()

        for i:=0; i<6; i++ {
            println(i)
            if i==3{
                runtime.Gosched()
            }
        }

    }()

    go func(){
        defer  wg.Done()
        println("Hello Golang")
    }()
    wg.Wait()
}
