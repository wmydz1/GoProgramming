package main
import (

    "net"
    "log"
    "encoding/json"
    "fmt"
)
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
    service := ":12000"
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    if err!=nil {
        log.Fatal(err.Error())
    }
    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err!=nil {
        log.Fatal(err.Error())
        log.Fatal(err.Error())
    }
    for {
        conn, err := listener.Accept()
        if err!=nil {
            log.Panicln(err.Error())
            continue
        }
        encoder := json.NewEncoder(conn)
        decoder := json.NewDecoder(conn)
        for n := 0; n<10; n++ {
            var person Person
            decoder.Decode(&person)
            fmt.Println(person.String())
            encoder.Encode(person)
        }
        conn.Close()
    }
}
