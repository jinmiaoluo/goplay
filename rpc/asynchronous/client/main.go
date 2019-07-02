package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/jinmiaoluo/goplay/rpc/asynchronous/server/server"
)

func main() {
	// connect rpc client
	// return a  rpc client
	client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Asynchronous call
	args := &server.Args{7, 8}
	quotient := new(server.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	<-divCall.Done // will be equal to divCall

	// check errors, print, etc.
	fmt.Printf("Quotient.A %% Quotient.B Quo is %d Rem is %d\n", quotient.Quo, quotient.Rem)
}
