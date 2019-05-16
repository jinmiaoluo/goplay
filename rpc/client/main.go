package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/jinmiaoluo/goplay/rpc/server/server"
)

func main() {
	// connect rpc client
	// return a  rpc client
	client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &server.Args{7, 8}
	var reply int
	//Call invokes the named function, waits for it to complete, and returns its errorstatus.
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	quotient := new(server.Quotient)
	err = client.Call("Arith.Divide", args, quotient)
	if err != nil {
		log.Fatal("quotient error:", err)
	}
	if err != nil {
		log.Fatal(err)
	}
	// check errors, print, etc.
	fmt.Printf("Quotient: %d divide %d quotient is %d remainder is %d\n", args.A, args.B, quotient.Quo, quotient.Rem)
}
