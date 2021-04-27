package rpcSupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)
	listen, err := net.Listen("tcp", ":"+host)
	if err != nil {
		return err
	}
	log.Printf("Listening on #%s", host)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", ":"+host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), err
}
