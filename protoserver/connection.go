package protoserver

import (
	"context"
	"fmt"
	"github.com/Metchain/Metblock/protoserver/routerpkg"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
	"net"
	"time"
)

// Connection represents a server connection.
type Connection interface {
	fmt.Stringer
	Start(router *routerpkg.Router)
	Disconnect()
	IsConnected() bool
	IsOutbound() bool
	SetOnDisconnectedHandler(onDisconnectedHandler OnDisconnectedHandler)
	SetOnInvalidMessageHandler(onInvalidMessageHandler OnInvalidMessageHandler)
	Address() *net.TCPAddr
}

// OnDisconnectedHandler is a function that is to be
// called once a Connection has been disconnected.
type OnDisconnectedHandler func()

// OnInvalidMessageHandler is a function that is to be called when
// an invalid message (cannot be parsed/doesn't have a route)
// was received from a connection.
type OnInvalidMessageHandler func(err error)

type OnConnectedHandler func(connection Connection) error

// Connect connects to the RPC server with the given address
func Connect(address string) (*GRPCClient, error) {
	const dialTimeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()

	gRPCConnection, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrapf(err, "error connecting to %s", address)
	}

	grpcClient := NewRPCClient(gRPCConnection)
	stream, err := grpcClient.MessageStream(context.Background(), grpc.UseCompressor(gzip.Name),
		grpc.MaxCallRecvMsgSize(RPCMaxMessageSize), grpc.MaxCallSendMsgSize(RPCMaxMessageSize))
	if err != nil {
		return nil, errors.Wrapf(err, "error getting client stream for %s", address)
	}
	return &GRPCClient{stream: stream, connection: gRPCConnection}, nil
}
