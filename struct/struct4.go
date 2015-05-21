package main
import (

    "fmt"
)
type Resource struct {
    id int
    name string
}
type Classify struct {
    id int
}
type User struct {
    Resource
    Classify
    name string
}
func main(){
     u:=User{Resource{1,"samchen"},Classify{10000},"Google Glass"}
    fmt.Println(u)

}
