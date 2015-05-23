package main
import (

    "sync"
    "testing"

)
//滥用 defer 可能会导致性能问题，尤其是在一个大循环里。
var lock sync.Mutex
func testm(){
    lock.Lock()
    lock.Unlock()
}
func testdefer(){
    lock.Lock()
    defer lock.Unlock()
}
func BenchmarkTest(b *testing.B){
    for i:=0;i<b.N ; i++ {
        testm()
    }
}
func BenchmarkTestDefer(b *testing.B){
    for i:=0; i<b.N;  i++{
        testdefer()
    }
}
func main(){

}
