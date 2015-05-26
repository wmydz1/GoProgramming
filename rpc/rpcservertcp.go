package main
import (

    "errors"
    "net/rpc"
    "net"
    "log"
    "os"
)

type Args struct {
    A, B int
}
type Quotient struct {
    Quo, Rem int
}
type Arith int

func (t *Arith) Multiply(arg *Args, reply *int) error {
    *reply=arg.A*arg.B
    return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
    if args.B ==0 {
        return errors.New("divider by  zero")
    }
    quo.Quo= args.A/args.B
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

    listener,err:=net.ListenTCP("tcp",tcpAddr)
    if err!=nil{
        log.Fatal(err)
    }
    for{
        conn,err:=listener.Accept()
        if err!=nil{
            continue
        }
        rpc.ServeConn(conn)
    }

}
