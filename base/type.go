package main
import (

    "fmt"
)
func main() {
    //引用类型包括 slice map 和 channel 。他们有着复杂的内部结构，除了申请内存外，还需要初始化相关属性。
    //内置函数 new 计算类型的大小，为其分配零值内存返回指针。而 make 会被编译器翻译成具体的创建函数，由其分配内存和
    //初始化成员结构，返回对象而非指针
    a := []int{0, 0, 0}  //提供初始化表达式
    a[1]=1

    b := make([]int, 3) //slice
    b[1]=10

   // c := new([]int)
   // c[1]=100  //(type *[]int does not support indexing)

    fmt.Println(a, b)

    // 不支持隐式类型转换，即便由窄向宽也不行
    var d byte =100

   // var n int =b //cannot use b (type byte) as type int in assignment

    fmt.Println(int(d))  //int(b) 显式转换

    //使用括号避免优先级错误
   /*
   *Point(p) // 相当于 *(Point(p)) (*Point)(p)
   <-chan int(c) // 相当于 <-(chan int(c)) (<-chan int)(c)
   */

    //同样不能把其他类型当做 bool 类型使用

}
