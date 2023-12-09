package protoserver

import (
	"google.golang.org/grpc"
)

// GRPCClient is a gRPC-based RPC client
type GRPCClient struct {
	stream     RPC_MessageStreamClient
	connection *grpc.ClientConn

	onDisconnectedHandler OnDisconnectedHandler
}
