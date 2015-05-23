package main
import (

    "os"
)
//关键字 defer 用于延迟调用，这些调用直到 return 执行前才被执行。通常用于释放资源和错误处理。
func work() error {
    f, err := os.Create("test.txt")
    if err!=nil {
        return err
    }
    defer f.Close()  //注册调用，而不是注册函数。必须提供参数，哪怕为空。
    f.WriteString("Hello, World!")
    return nil


}
func main() {
    work()
}