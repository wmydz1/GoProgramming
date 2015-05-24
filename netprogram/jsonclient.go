package main
import (

    "encoding/json"
    "net"
    "log"
    "os"
    "fmt"
)
//注意 首字母大写才可以被Go导出
type Person struct {
    Name  Name
    Email []Email
}
type Name struct {
    Family string
    Personal string
}
type Email struct {
    Kind string
    Address string
}
func (p Person) String() string {
    s := p.Name.Personal+" "+p.Name.Family
    for _, v := range p.Email {
        s+="\n"+v.Kind+":"+v.Address
    }
    return s
}
func main() {
    person := Person{
        Name:Name{Family:"Coogle", Personal:"Go"},
        Email:[]Email{Email{Kind:"home", Address:"golang@gmail.com"},
            Email{Kind:"work", Address:"google@gmail.com"},
        },
    }
    service := ":12000"
    conn, err := net.Dial("tcp", service)
    if err!=nil {
        log.Fatal(err.Error())
        os.Exit(1)
    }
    encoder := json.NewEncoder(conn)
    decoder := json.NewDecoder(conn)
    for n := 0; n<10; n++ {
        encoder.Encode(person)
        var newPerson Person
        decoder.Decode(&newPerson)
        fmt.Println(newPerson.String())
    }
    os.Exit(0)
}