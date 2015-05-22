package main
import (

    "time"
    "fmt"
)
func main(){
    now :=time.Now().Format(time.Stamp)
    fmt.Println(now)
    fmt.Println(Substr(now,7,5))
    fmt.Println(time.Now().Format("15:04:05"))
    fmt.Println(time.Now().String()[11:16])

}
func Substr(str string, start, length int) string {
    rs := []rune(str)
    rl := len(rs)
    end := 0

    if start < 0 {
        start = rl - 1 + start
    }
    end = start + length

    if start > end {
        start, end = end, start
    }

    if start < 0 {
        start = 0
    }
    if start > rl {
        start = rl
    }
    if end < 0 {
        end = 0
    }
    if end > rl {
        end = rl
    }

    return string(rs[start:end])
}
