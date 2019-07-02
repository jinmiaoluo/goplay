// the method's type is exported.
// the method is exported.
// the method has two arguments, both exported (or builtin) types.
// the method's second argument is a pointer.
// the method has return type error.
package server

import (
	"errors"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo int //quotient(商)
	Rem int //remainder(余数)
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
