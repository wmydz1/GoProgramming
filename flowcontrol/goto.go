package main
import (

)
//支持函数在 goto 内进行跳转，标签名区分大小写，未使用标签引发错误。
func main() {
    var i int
    for {
        println(i)
        i++
        if i>2 {
            goto BREAK
        }
    }

    BREAK:
    println("break")
    // EXIT:    //label EXIT defined and not used

    //配合标签，break和 continue可在多级嵌套循环中跳出

    L1:
    for x := 0; x<3; x++ {
        L2: for y := 0; y<5; y++ {
            if y>2 {
                continue L2
            }
            if x>1 {
                break L1
            }
            print(x, ":", y, " ")
        }
        println()
    }

    //break 可用于 for switch select,而 continue 只能用于 for 循环

}
