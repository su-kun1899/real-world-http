package main

import (
	"net/rpc/jsonrpc"
	"net"
	"net/http"
	"net/rpc"
	"log"
)

// メソッドが属す構造体
type Calculator int

// RPCで外部から呼ばれるメソッド
func (c *Calculator) Multiply(args Args, result *int) error {
	log.Printf("Multiply called: %d, %d\n", args.A, args.B)
	*result = args.A * args.B

	return nil
}

// 外部から呼ばれるときの引数
type Args struct {
	A, B int
}

func main() {
	calculator := new (Calculator)
	server := rpc.NewServer()
	server.Register(calculator)
	http.Handle(rpc.DefaultRPCPath, server)
	log.Println("start http listening :18888")
	listener, err := net.Listen("tcp", ":18888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
