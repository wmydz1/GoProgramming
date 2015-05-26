package main
import (


    "errors"
    "net/rpc"
    "net/http"
    "log"
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
func main(){
    arith:=new(Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()

    err :=http.ListenAndServe(":12000",nil)
    if err!=nil{
        log.Fatal(err.Error())
    }
}

/*
我们注册了一个Arith的RPC服务，然后通过rpc.HandleHTTP函数把该服务注册到了HTTP协议上，然后我们就可以利用http的方式来传递数据了
*/