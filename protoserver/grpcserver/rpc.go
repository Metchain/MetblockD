package grpcserver

import (
	"github.com/Metchain/MetblockD/blockchain"
	"github.com/Metchain/MetblockD/domain"
	"github.com/Metchain/MetblockD/protoserver/grpcserver/protowire"
	"github.com/Metchain/MetblockD/protoserver/server"
	"github.com/Metchain/MetblockD/utils/panics"
)

// RPCMaxMessageSize is the max message size for the RPC server to send and receive
const RPCMaxMessageSize = 1024 * 1024 * 1024 // 1 GB
type rpcServer struct {
	protowire.UnimplementedRPCServer
	gRPCServer
	*domain.Metchain
	*blockchain.Blockchain
}

// NewRPCServer creates a new RPCServer
func NewRPCServer(listeningAddresses []string, rpcMaxInboundConnections int, bc *blockchain.Blockchain) (server.Server, error) {
	gRPCServer := newGRPCServer(listeningAddresses, RPCMaxMessageSize, rpcMaxInboundConnections, "RPC")
	rpcServer := &rpcServer{gRPCServer: *gRPCServer, Blockchain: bc}
	protowire.RegisterRPCServer(gRPCServer.server, rpcServer)
	return rpcServer, nil
}

func (r *rpcServer) MessageStream(stream protowire.RPC_MessageStreamServer) error {
	defer panics.HandlePanic(log, "rpcServer.MessageStream", nil)

	return r.handleInboundConnection(stream.Context(), stream)
}
