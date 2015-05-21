package main
import (

    "fmt"
)
func main() {
    //预先给 make 函数一个合理元素数量参数,有助于提升性能。因为事先申请一⼤块内存, 可避免后续操作时频繁扩张。
    //m := make(map[string]int, 1000)
    m := map[string]int{
        "a":10000,
    }
    // 判断 key 是否存在。
    if v, ok := m["a"]; ok {
        fmt.Println(v)
    }

    // 对于不存在的 key,直接返回 \0,不会出错。
    fmt.Println(m["c"])

    // 新增或修改。
    m["b"]=9999
    // 删除。如果 key 不存在,不会出错。
    delete(m, "c")
    // 获取键值对数量。cap 无效
    fmt.Println(len(m))
    // 迭代,可仅返回 key。随机顺序返回,每次都不相同。
    for k, v := range m {
        fmt.Println(k, v)
    }
    //不能保证迭代返回次序,通常是随机结果,具体和版本实现有关
    //从 map 中取回的是⼀一个 value 临时复制品,对其成员的修改是没有任何意义的
    /*
    type user struct{ name string }
    m := map[int]user{
    1: {"user1"},
    }
    m[1].name = "Tom"         // Error: cannot assign to m[1].name
    */

    //正确做法是完整替换或使用指针。
    type user struct {
        name string
    }
    usermap :=map[int]user{
        1:{"Chrome"},
    }
    u:=usermap[1]
    u.name="FireFox"
    // 替换 value。
    usermap[1]=u
     fmt.Println(usermap)
   //--------------------------------
    usermap2 :=map[int]*user{
        1:&user{"Chrome"},
    }
    // 返回的是指针复制品。透过指针修改原对象是允许的。
    fmt.Println(usermap2[1].name)
    usermap2[1].name="Safari"
    fmt.Println(usermap2[1].name)

    //可以在迭代时安全删除键值。但如果期间有新增操作,那么就不知道会有什么意外了
    for i:=0; i<5;i++  {
        maps :=map[int]string{
            0:  "a", 1:  "a", 2:  "a", 3:  "a", 4:  "a",
            5:  "a", 6:  "a", 7:  "a", 8:  "a", 9:  "a",
        }
        for k:=range maps{
            maps[k+k]="b"
            delete(maps,k)
        }
        fmt.Println(maps)
    }
}
