package main
import (


    "net"
    "log"
    "os"
    "time"
)
func main() {
    sevice := ":12000"
    tcpAddr, err := net.ResolveTCPAddr("tcp", sevice)
    if err!=nil {
        log.Fatal(err.Error())
        os.Exit(1)
    }
    listen, err := net.ListenTCP("tcp", tcpAddr)
    if err!=nil {
        log.Fatal(err.Error())
        os.Exit(1)
    }
    for {

        conn, err := listen.Accept()
        if err!=nil {
            log.Fatal(err.Error())
            continue
        }
        daytime := time.Now().Format(time.ANSIC)
        conn.Write([]byte(daytime))
        conn.Close()

    }
}
