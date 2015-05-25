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
    Book []string

}
/*
Unix用户已经很熟悉什么是pipe了，ls | grep "beego"类似这样的语法你是不是经常使用，过滤当前目录下面的文件，显示含有"beego"的数据，
表达的意思就是前面的输出可以当做后面的输入，最后显示我们想要的数据，而Go语言模板最强大的一点就是支持pipe数据，
在Go语言里面任何{{}}里面的都是pipelines数据，例如我们上面输出的email里面如果还有一些可能引起XSS注入的，那么我们如何来进行转化呢？
{{. | html}}
*/

func do3(w http.ResponseWriter, r *http.Request) {
    t := template.New("model3.html")

    workplace, err := os.Getwd()
    if err!=nil {
        log.Fatal(err.Error())
        os.Exit(1)
    }
    tplpath := workplace+"/netprogram/tpl/model3.html"

    t, err=t.ParseFiles(tplpath)


    if err!=nil {
        log.Fatal(err)
        os.Exit(2)
    }



    //    p := Person{UserName:"samchen"}
    p := Person{Book:[]string{"Google","Android","Golang"}}
    t.Execute(w, p)


}





func main() {
    http.HandleFunc("/", do3)
    http.ListenAndServe(":10000", nil)
}

