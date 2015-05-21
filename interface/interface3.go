package main
import(

    "fmt"
)
//匿名接口可用作变量类型,或结构成员
type Tester struct{
    s interface{
        String() string
    }

}
type User struct {
    id int
    name string
}
func (user *User) String() string{
    return fmt.Sprintf("user %d,%s",user.id,user.name)
}

func main(){
      t:=Tester{&User{10000,"samchen"}}
    fmt.Println(t.s.String())
}

