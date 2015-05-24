package main
import (

    "os"
    "runtime/pprof"
)
//PProf
//监控程序执行，找出性能瓶颈
//除了调用 runtime/pprof 相关函数外，还可以直接用测试命令输出所需要的记录文件。
func main(){
    //CPU
    cpu,_:=os.Create("cpu.out")
    defer cpu.Close()
   pprof.StartCPUProfile(cpu)
    defer pprof.StopCPUProfile()

    //Memory
    mem,_:=os.Create("mem.out")
    defer mem.Close()
    defer pprof.WriteHeapProfile(mem)

    //code...
}

//go test
/*
-blockprofile block.out  //goroutine 阻塞
-blockprofilerate n    //超出该参数设置时间的阻塞才会被记录，单位 纳秒
-cpuprofile cpu.out
-memprofile  mem.out //内存分配
-memfilerate n //超出该参数设置的内存分配才会被记录 单位 字节 默认 512KB



*/
//go tool pprof -text test cpu.out
//浏览器查看
//go tool pprof -web test cpu.out