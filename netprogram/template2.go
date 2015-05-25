package main
import (

    "net/http"

    "html/template"
    "os"
    "log"

)
//未导出的字段，首字母是小写的
//如果要导出，首字母一定要大写
type Person struct {
    UserName string

}

func do2(w http.ResponseWriter, r *http.Request) {
    t := template.New("model2.html")

    workplace, err := os.Getwd()
    if err!=nil {
        log.Fatal(err.Error())
        os.Exit(1)
    }
    tplpath := workplace+"/netprogram/tpl/model2.html"

    t, err=t.ParseFiles(tplpath)


    if err!=nil {
        log.Fatal(err)
        os.Exit(2)
    }



//    p := Person{UserName:"samchen"}
    p := Person{}
    t.Execute(w, p)


}
func main() {
    http.HandleFunc("/", do2)
    http.ListenAndServe(":10000", nil)
}
