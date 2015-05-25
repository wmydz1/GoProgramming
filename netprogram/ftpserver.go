package main
import (

    "net"
    "log"
    "os"
)
const (
    DIR = "DIR"
    CD = "CD"
    PWD = "PWD"
)
func main() {
    service := ":12000"
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    if err!=nil {
        log.Fatal(err.Error())
        os.Exit(1)
    }

    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err!=nil {
        log.Fatal(err.Error())
    }
    for {
        conn, err := listener.Accept()
        if err!=nil {
            log.Fatal(err.Error())
            continue
        }
        go handleClient(conn)
    }

}
func handleClient(conn net.Conn) {
    defer conn.Close()
    var buf [1024]byte
    for {
        n, err := conn.Read(buf[0:])
        if err!=nil {
            log.Fatal(err)
            conn.Close()
            return
        }
        s := string(buf[0:n])
        //decode request
        if s[0:2] ==CD {
            chdir(conn, s[3:])
        }else if s[0:3]==DIR {
            dirList(conn)
        }else if s[0:3] ==PWD {
            pwd(conn)
        }

    }
}
func chdir(conn net.Conn, s string) {
    if os.Chdir(s)==nil {
        conn.Write([]byte("OK"))
    }else {
        conn.Write([]byte("ERROR"))
    }

}
func pwd(conn net.Conn) {
    s, err := os.Getwd()
    if err!=nil {
        conn.Write([]byte(""))
    }
    conn.Write([]byte(s))
}
func dirList(conn net.Conn) {
    defer conn.Write([]byte("\r\n"))

    dir, err := os.Open(".")
    if err!=nil {
        log.Fatal(err.Error())
        return
    }
    names, err := dir.Readdirnames(-1)
    if err!=nil {
        log.Fatal(err.Error())
        return
    }
    for _, nm := range names {
        conn.Write([]byte(nm+"\r\n"))
    }
}