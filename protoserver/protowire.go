package protoserver

import (
	"context"
	"github.com/Metchain/Metblock/proto"
	"google.golang.org/grpc"
)

// RPCMaxMessageSize is the max message size for the RPC server to send and receive
const RPCMaxMessageSize = 1024 * 1024 * 1024 // 1 GB

type rPCMessageStreamClient struct {
	grpc.ClientStream
}

type RPC_MessageStreamClient interface {
	Send(message *proto.MetchainMessage) error
	Recv() (*proto.MetchainMessage, error)
	grpc.ClientStream
}

func (x *rPCMessageStreamClient) Send(m *proto.MetchainMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rPCMessageStreamClient) Recv() (*proto.MetchainMessage, error) {
	m := new(proto.MetchainMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	MessageStream(ctx context.Context, opts ...grpc.CallOption) (RPC_MessageStreamClient, error)
}
type rPCClient struct {
	cc grpc.ClientConnInterface
}

func (r rPCClient) MessageStream(ctx context.Context, opts ...grpc.CallOption) (RPC_MessageStreamClient, error) {
	//TODO implement me
	panic("implement me")
}
