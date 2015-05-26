package main
import (

    "net/rpc"
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
    client,err:=rpc.Dial("tcp",service)
    if err!=nil{
        log.Fatal("dialng",err)
    }
    //synchronous call
    args :=Args{17,8}
    var reply int
    err=client.Call("Arith.Multiply",args,&reply)
    if err!=nil{
        log.Fatal("arith error",err)
    }
    fmt.Printf("Arith: %d * %d = %d\n",args.A,args.B,reply)
    var quo Quotient
    err=client.Call("Arith.Divide",args,&quo)
    if err!=nil{
        log.Fatal("arith error",err)
    }
    fmt.Printf("Arith: %d / %d = %d remainder %d",args.A,args.B,quo.Quo,quo.Rem)
}
