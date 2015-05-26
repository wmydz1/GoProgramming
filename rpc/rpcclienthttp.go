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

func main() {
    serverAddress := "127.0.0.1"
    client,err:=rpc.DialHTTP("tcp",serverAddress+":12000")
    if err!=nil{
        log.Fatal("dialing",err)
    }
    //Synchronous call
    args:=Args{1000,8}
    var reply int
    err=client.Call("Arith.Multiply",args,&reply)
    if err !=nil{
        log.Fatal("arith error",err)
    }
    fmt.Printf("Arith: %d*%d=%d\n",args.A,args.B,reply)

    var quo Quotient
    err=client.Call("Arith.Divide",args,&quo)
    if err!=nil{
        log.Fatal("arith error",err)
    }
    fmt.Printf("Arith: %d/%d=%d remainder %d",args.A,args.B,quo.Quo,quo.Rem)
}

/*
Go标准包中已经提供了对RPC的支持，而且支持三个级别的RPC：TCP、HTTP、JSONRPC。但Go的RPC包是独一无二的RPC，它和传统的RPC系统不同
，它只支持Go开发的服务器与客户端之间的交互，因为在内部，它们采用了Gob来编码。

Go RPC的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细的要求如下：

函数必须是导出的(首字母大写)

必须有两个导出类型的参数，

第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的

函数还要有一个返回值error
*/