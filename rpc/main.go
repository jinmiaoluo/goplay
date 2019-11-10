package main

import (
	"net"
	"net/http"
	"net/rpc"
	"github.com/jinmiaoluo/goplay/rpc/server"
)

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}
