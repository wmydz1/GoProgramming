package main
import (

    "net/http"
    "log"
    "os"
    "fmt"
)
func main() {
    path,_ :=os.Getwd()
    filepath:=path+"/netprogram/home/"
    fmt.Println(path)
    fileServer := http.FileServer(http.Dir(filepath))
    err := http.ListenAndServe(":10000", fileServer)
    if err!=nil {
        log.Fatal(err.Error())
        os.Exit(1)
    }
}
