package main
import (

    "fmt"
)
// 字符串是不可变值类型，内部用指针指向 UTF-8 数组 。
/*
  1 默认是空字符串
  2 用索引号访问某个字节，如s[i]
  3 不能用序号获得字节元素的指针，&s[i] 非法
  4 不可变类型，无法修改字节数组。
  5 字节数组尾部不包含 NULL
*/
func main() {
    //使用 "`" 定义不做转义处理的原始字符串，支持跨行
    s := `Hello
           Golang
        Google Inc
    `
    fmt.Println(s)

    //连接跨行字符串时，"+" 号必须在上一行的末尾，否则容易导致编译错误
    add := "Google " +
    "Android"
    fmt.Println(add)

    //支持用两个索引返回字串。子串依然指向原字节数组，仅修改了指针和长度属性

    str := "Golang Google Inc"
    lang  :=str[0:6]
    name :=str[7:13]
    com :=str[14:]

    fmt.Println(lang)
    fmt.Println(name)
    fmt.Println(com)

    //单引号字符常量 Unicode Code Point , 支持 \uFFFF \U7FFFFFFF \xFF
    //对应 rune UCS-4
    fmt.Printf("%T\n",'a')

      var c1,c2 rune = '\u6211','们'
//    var c1, c2 rune = '\u6211', '的'
    fmt.Println(c1 =='我',string(c2)=="\xe4\xbb\xac")


    //要修改字符串，可以先将其转换成[]rune 或者[]byte,完成后再转化为 string 。无论那种转换
    //都会重新分配内存，并且复制字节数组

    change:="Hello GO"
    change_byte:=[]byte(change)
    change_byte[6]='C'
    change_byte[7]='#'
    fmt.Println(string(change_byte))

    tool:="电脑"
    tool_byte:=[]rune(tool)
    tool_byte[1]='话'
    fmt.Println(string(tool_byte))

    // for 循环遍历字符串时，也有 rune 和 byte 两种方式

    loop :="Android Google "
    for i:=0;i<len(loop) ;i++  {  //byte
        fmt.Printf("%c",loop[i])
    }
    fmt.Println()

    for _,r:=range loop{ //rune
         fmt.Printf("%c",r)
    }
}
