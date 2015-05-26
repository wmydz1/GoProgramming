package main
import (

    "errors"
    "net/rpc"
    "net"
    "log"
    "net/rpc/jsonrpc"
    "os"
)
type Arith int

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}
func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply =args.A*args.B
    return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
    if args.B==0 {
        return errors.New("divide by zero")
    }
    quo.Quo=args.A/args.B
    quo.Rem=args.A%args.B
    return nil
}
func main() {
    arith := new(Arith)
    rpc.Register(arith)
    tcpAddr, err := net.ResolveTCPAddr("tcp", ":12000")
    if err!=nil {
        log.Fatal(err)
        os.Exit(1)
    }
    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err!=nil {
        log.Fatal(err)
        os.Exit(1)
    }
    for {
        conn, err := listener.Accept()
        if err!=nil {
            log.Fatal(err)
            continue
        }
        jsonrpc.ServeConn(conn)
    }
}