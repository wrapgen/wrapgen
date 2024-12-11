package main

import (
	"context"
	"fmt"
	"log"
	"net/rpc"

	"github.com/wrapgen/wrapgen/testdata/examples/rpc/common"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	arith := common.ArithClient{Client: client}

	reply, err := arith.Multiply(context.Background(), &common.Args{A: 10, B: 20})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Multiply => %+v\n", *reply)
}
