package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/jinmiaoluo/goplay/rpc/synchronous/server/server"
)

func main() {
	arith := new(server.Arith)         // create a type instance
	rpc.Register(arith)                // register instance
	rpc.HandleHTTP()                   // register http handler to DefaultServer
	l, e := net.Listen("tcp", ":1234") // announce network address
	if e != nil {
		log.Fatal("listen error:", e)
	}
	//Serve accepts incoming HTTP connections on the listener l, creating a new
	//service goroutine for each. The service goroutines read requests and then call
	//handler to reply to them.
	http.Serve(l, nil)
}
