package test
import (

    "testing"
    "time"


    "fmt"
)

func sum(n ...int) int {
    var c int
    for _, i := range n {
        c+=i
    }
    return c
}
func TestSum(t *testing.T) {
    time.Sleep(time.Second*2)
    if sum(1, 2, 3)!=6 {
        t.Fatal("error from sum")
    }

}
func TestTimeout(t *testing.T) {
    time.Sleep(time.Second*3)
}
// gotest 默认执行所有单元测试函数，支持 go build 参数
/*
-c 仅编译，不执行测试。
-v 显示所有测试函数的细节
-run regex 执行指定的测试函数（正则表达式）
-parallel n 并发执行测试函数（默认 ：GOMAXPROCS）
-timeout 单元测试的超时时间 -timeout 2m30s

*/
//$ go test -v -timeout 3s
//go test -v -run "(?i)sum"


//性能测试需要运行足够多的次数才能计算单次执行的平均时间。
func BenchmarkSum(b *testing.B) {
    for i := 0; i<b.N; i++ {
        if sum(1, 2, 3)!=6 {
            b.Fatal("sum error")
        }
    }
}

//go test -v -bench .
//go test -bench . -benchmem -cpu 1,2,4 -benchtime 30s

//默认情况下，go test 不会执行性能测试，需要使用 -beach 参数
/*
-beach regex 指定性能测试函数（正则表达式）
-beachmem 输出内存统计信息
-beachtime t 设置每个性能测试运行的时间 -beachtime 1m30s
-cpu 设置并发测试 默认 GOMAXPROCS -cpu 1,2,4
*/

//Example
//和 testing.T类似，区别通过捕获 stdout 输出来判断测试结果

func ExampleSum(){
    fmt.Println(sum(1,2,3))
    fmt.Println(sum(10,20,30))
}
//go test -v


//Cover 除了显示覆盖率百分比外，还输出详细分析记录文件

/*
-cover 允许覆盖分析
-covermode 代码分析模式（set:是否执行，count 执行次数，atomic 次数，并发支持）
-converprofile 输出结果文件
*/
//go test -cover -coverprofile=cover.out -covermode=count
//go tool cover -func=cover.out

//用浏览器输出结果，能查看更详细直观的信息，包括用不同颜色标记覆盖 、运行次数等等。
//go tool cover -html=cover.out
//说明将鼠标移到代码块可以看见具体的运行次数。


