package main
import (

    "net/http"

    "html/template"
    "os"
    "log"
    "fmt"
    "strings"
)
type Friend struct {
    Fname string
}
type Person struct {
    UserName string
    Emails []string
    Friends []*Friend
}
func EmailDealWith(args ...interface{}) string {
    ok := false
    var s string
    if len(args)==1 {
        s, ok=args[0].(string)
    }
    if !ok {
        s=fmt.Sprint(args...)
    }
    // find the @ symbol
    substrs := strings.Split(s, "@")

    if len(substrs)!=2 {
        return s
    }
    //replace the @ by at
    return (substrs[0]+" at "+substrs[1])
}
func do4(w http.ResponseWriter, r *http.Request) {
    t := template.New("model4.html")
    workplace, err := os.Getwd()
    if err!=nil {
        log.Fatal(err.Error())
        os.Exit(1)
    }
    tplpath := workplace+"/netprogram/tpl/model4.html"

    t, err=t.ParseFiles(tplpath)

 //  template.Must(t.ParseFiles(tplpath))

    if err!=nil {
        log.Fatal(err)
        os.Exit(2)
    }

    t = t.Funcs(template.FuncMap{"emailDeal":EmailDealWith})

    p := Person{UserName:"samchen",
        Emails:[]string{"wmydz1@gmail.com", "golang@gmail.com"},
        Friends:[]*Friend{&Friend{Fname:"Google"}, &Friend{Fname:"Apple"}}}

    t.Execute(w, p)

}


func main() {
    http.HandleFunc("/", do4)
    http.ListenAndServe(":10000", nil)
}
