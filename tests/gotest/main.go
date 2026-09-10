package main

import (
	"github.com/iqtlabs/faucetconfrpc/faucetconfserver"
	"google.golang.org/grpc"
)

type faucetconfrpcer struct {
	conn   *grpc.ClientConn
	client faucetconfserver.FaucetConfServerClient
}

func main() {}
