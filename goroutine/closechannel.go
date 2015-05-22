package main
import (

    "sync"
    "fmt"
    "time"
)
func main(){
    var wg sync.WaitGroup
    quit:=make(chan  bool)
    for i:=1;i<3 ;i++  {
      wg.Add(1)
        go func(id int){
          defer  wg.Done()
          task:=func(){
              fmt.Println(id,time.Now().Nanosecond())
              time.Sleep(time.Second)
          }
            for{
                select {
                       case <-quit:     //closed channel 不会阻塞，因此可以用作退出通知

                       return
                   default:     //正常执行任务
                        task()

                }
            }
        }(i)
    }
    time.Sleep(time.Second*5)  //让测试的 goroutine 运行一会儿
    close(quit)
    wg.Wait()

}
