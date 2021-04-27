package main

import (
	"crawlerByGo/src/__HelloDemo/rpcdemo"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//{"method":"aaa"}
// {"method":"DemoService.Div","params":[{"A":1,"B":0}],"id":1}
func main() {
	rpc.Register(rpcdemo.DemoService{})
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
