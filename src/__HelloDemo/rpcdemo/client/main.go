package main

import (
	"crawlerByGo/src/__HelloDemo/rpcdemo"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	error := client.Call("DemoService.Div", rpcdemo.Args{1, 2}, &result)
	fmt.Printf("result is #%v, error is #%v\n", result, error)
	error = client.Call("DemoService.Div", rpcdemo.Args{10, 3}, &result)
	fmt.Printf("result is #%v, error is #%v\n", result, error)
	error = client.Call("DemoService.Div", rpcdemo.Args{1, 0}, &result)
	fmt.Printf("result is #%v, error is #%v\n", result, error)
}
