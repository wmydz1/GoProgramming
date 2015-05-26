package main
import (

    "net/http"
    "fmt"
)
type MyMux struct {

}
func (p *MyMux) ServeHTTP(w http.ResponseWriter,r *http.Request){
    if r.URL.Path=="/"{
        sayHelloName(w,r)
        return
    }
    http.NotFound(w,r)
    return
}
func sayHelloName(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w,"This is myrouter")
}

func main(){
    mux :=&MyMux{}
    http.ListenAndServe(":10000",mux)
}