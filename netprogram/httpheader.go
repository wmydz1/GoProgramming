package main
import (

    "net/http"
    "fmt"
    "os"
)
func main() {
    url := "http://www.baidu.com/"
    response, err := http.Head(url)
    if err!=nil {
        fmt.Println(err.Error())
        os.Exit(2)
    }
    fmt.Println(response.Status)
    fmt.Println("-----------------------")
    for k, v := range response.Header {
        fmt.Println(k, ":", v)
    }
    os.Exit(0)
}
