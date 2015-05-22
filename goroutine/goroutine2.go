package main
import (

    "sync"
    "runtime"
)

func main(){
    wg:=new(sync.WaitGroup)
    wg.Add(1)
    go func(){
         defer  wg.Done()
         defer  println("A.defer")

        func(){
            defer  println("B.defer")
            //立即终止当前goroutine执行，调度器确保所有已经注册 defer 延迟调用被执行
            runtime.Goexit()
            //下面语句不会执行
            println("B")
        }()
        //println("A") 不会执行
        println("A")

    }()

    wg.Wait()
}