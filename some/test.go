package main
import (

// "fmt"
    "fmt"
)
func main() {
    ch := make(chan int)

    for i := 1; i<100; i++ {
        go func(id int) {
            fmt.Println("Go",id)

        }(i)
        if i ==50{
            ch <- 1
        }
    }


    <-ch

}
