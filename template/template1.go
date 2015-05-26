package main
import (

    "net/http"

    "fmt"
    "log"
    "strings"
    "html/template"
    "os"
)
func sayHello(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key", k)
        fmt.Println("val", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Go Cloud Power")
}
func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method =="GET" {
        workspace, _ := os.Getwd()
        tplpath := workspace+"/template/login.tpl"


        t, _ := template.ParseFiles(tplpath)
        t.Execute(w, nil)
    }else {
        r.ParseForm()
        fmt.Println("username", r.Form["username"])
        fmt.Println("password", r.Form["password"])
    }
}
func main() {
    http.HandleFunc("/", sayHello)
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":10000", nil)
    if err!=nil {
        log.Fatal("ListenAndServe: ", err)
    }
}