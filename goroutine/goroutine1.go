package main
import (
    "math"
    "fmt"
    "sync"
    "runtime"
)
func sum(id int) {
    var x int64
    for i := 0; i<math.MaxUint32; i++ {
        x+=int64(i)
    }
    fmt.Println(id, x)
}

func main() {
    //默认情况下，进程启动后仅允许一个系统线程服务于goroutine
    //runtime.GOMAXPROCS(runtime.NumCPU()) 可以让调度器用多线程实现多核并行，而不仅仅是并发
    runtime.GOMAXPROCS(runtime.NumCPU())
    wg := new(sync.WaitGroup)
    wg.Add(2)
    for i := 0; i<2; i++ {
        go func(id int) {
            defer wg.Done()
            sum(id)
        }(i)
    }
    wg.Wait()
}
