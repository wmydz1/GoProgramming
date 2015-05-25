package main
import (

    "net/http"
    "fmt"
    "os"
    "net/http/httputil"
    "log"
    "strings"
)
func main() {
    url := "http://www.baidu.com/"
    response, err := http.Get(url)
    if err!=nil {
        fmt.Println(err.Error())
        os.Exit(2)
    }
    if response.Status!="200 OK" {
        fmt.Println(response.Status)
        os.Exit(0)
    }
    b, err := httputil.DumpResponse(response, false)
    if err!=nil {
        log.Fatal(err.Error())
    }
    fmt.Println(string(b))

    fmt.Println("--------------------")
    contentTypes := response.Header["Content-Type"]
    if !acceptableCharset(contentTypes) {
        fmt.Println("Cannot handle", contentTypes)
        os.Exit(4)
    }
    var buf [512]byte
    reader := response.Body
    for {
        n, err := reader.Read(buf[0:])
        if err!=nil {
            //            log.Fatal(err.Error())
            //            os.Exit(1)
            //            return
            continue
        }
        fmt.Println(string(buf[0:n]))
    }
    os.Exit(0)


}
func acceptableCharset(contentTypes []string) bool {
    //we want uft8 only
    for _, cType := range contentTypes {
        if strings.Index(cType, "utf-8")!=-1 {
            return true
        }
    }
    return false
}