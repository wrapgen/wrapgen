package main

import (
	"errors"
	"log/slog"
	"net"
	"net/http"
	"net/rpc"

	"github.com/wrapgen/wrapgen/testdata/examples/rpc/common"
)

type Server struct{}

var _ common.Arith = &Server{}

func (t *Server) Multiply(args *common.Args, reply *int) error {
	slog.Info("Multiply", "a", args.A, "b", args.B)
	*reply = args.A * args.B
	return nil
}

func (t *Server) Divide(args *common.Args, quo *common.Quotient) error {
	slog.Info("Divide", "a", args.A, "b", args.B)
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	s := new(Server)
	err := rpc.RegisterName("Arith", s)
	if err != nil {
		panic(err)
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	http.Serve(l, nil)
}
