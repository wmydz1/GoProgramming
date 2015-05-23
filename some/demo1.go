package main
import (

    "fmt"
    "reflect"
)
type Apple struct {
      name string
}
type Fruit struct {
     Apple
     name string
}
func (a Apple) eat(){
    fmt.Println("eat ",a.name)
    fmt.Println(reflect.ValueOf(a))
}
func main(){
    f :=Fruit{Apple{"apple"},"Fruits"}
    f.eat()

}
