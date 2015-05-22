package main
import (

    "fmt"
    "unicode/utf8"
)
func main(){
    data:="❤"
    fmt.Println(utf8.RuneCountInString(data))
}
