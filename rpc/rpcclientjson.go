package main
import (

    "net/rpc/jsonrpc"
    "log"
    "fmt"
)
type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}
func main(){
    service:=":12000"
    client,err:=jsonrpc.Dial("tcp",service)
    if err!=nil{
        log.Fatal("dialing",err)
    }
    //Synchronous call
    args:=Args{17,9}
    var reply int
    err=client.Call("Arith.Multiply",args,&reply)
    if err!=nil{
        log.Fatal("arith error",err)
    }
    fmt.Printf("Arith %d * %d = %d\n",args.A,args.B,reply)

    var quo Quotient

    err =client.Call("Arith.Divide",args,&quo)
    if err !=nil{
        log.Fatal("arith error",err)
    }
    fmt.Printf("Arith %d/%d = %d remainder %d\n",args.A,args.B,quo.Quo,quo.Rem)
}
//json-rpc是基于TCP协议实现的，目前它还不支持HTTP方式。