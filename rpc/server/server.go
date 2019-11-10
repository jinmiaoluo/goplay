// the method's type is exported.
// the method is exported.
// the method has two arguments, both exported (or builtin) types.
// the method's second argument is a pointer.
// the method has return type error.
package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
//,de

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

// multiply object's fields
// return result
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// divide object's fields
// Quotient object store quotient and remainder
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith) // create a type instance
	rpc.Register(arith) // register instance
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}
