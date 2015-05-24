package main
import (


    "log"
    "net"
    "bytes"
    "io"
    "fmt"
)
func main() {

    service := ":12000"
    conn, err := net.Dial("tcp", service)
    if err!=nil {
        log.Fatal(err.Error())
    }
    result, err := ReadData(conn)
    if err!=nil {
        log.Fatal(err.Error())
    }
    fmt.Println(string(result))

}

func ReadData(conn net.Conn) ([]byte, error) {
    defer conn.Close()
    result := bytes.NewBuffer(nil)
    var buf [1024]byte

    for {
        n, err := conn.Read(buf[0:])
        result.Write(buf[0:n])
        if err!=nil {
            if err ==io.EOF {
                break
            }
            return nil, err
        }

    }
    return result.Bytes(), nil
}
