package main
import (

    "net/http"
    "os"
    "html/template"
)
func do5(w http.ResponseWriter, r *http.Request) {

    path, _ := os.Getwd()
    tplheader := path+"/netprogram/tpl/header.tpl"
    tplcontent := path+"/netprogram/tpl/content.tpl"
    tplfooter := path+"/netprogram/tpl/footer.tpl"
    s, _ := template.ParseFiles(tplheader, tplcontent, tplfooter)
    //    s,_:=template.ParseFiles(tplcontent,tplfooter,tplheader)
    //      s.ExecuteTemplate(w,"header",nil)
    s.ExecuteTemplate(w, "content", nil)
    //      s.ExecuteTemplate(w,"footer",nil)
}
func main() {
    http.HandleFunc("/", do5)
    http.ListenAndServe(":10000", nil)
}
