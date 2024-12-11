package common

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

//wrapgen:generate -template ./rpc -destination client_gen.go
type Arith interface {
	Multiply(args *Args, reply *int) error
	Divide(args *Args, quo *Quotient) error
}
