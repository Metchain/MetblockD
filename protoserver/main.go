package protoserver

import (
	"fmt"
	"github.com/Metchain/Metblock/blockchain"
	"github.com/Metchain/Metblock/domain"
	pb "github.com/Metchain/Metblock/proto"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"log"
	"net"
	"sync"
)

// define the port
const (
	port = ":14041"
)

type rpcServer struct {
	pb.UnimplementedRPCServer
	gRPCServer
	*domain.Metchain
	*blockchain.Blockchain
}

type p2pServer struct {
	pb.UnimplementedP2PServer
	gRPCServer
	*domain.Metchain
	*blockchain.Blockchain
}

const RPCMaxMessageSize = 1024 * 1024 * 1024

type gRPCServer struct {
	listeningAddresses []string
	server             *grpc.Server
	name               string

	maxInboundConnections      int
	inboundConnectionCount     int
	inboundConnectionCountLock *sync.Mutex
}

func newGRPCServer() *gRPCServer {
	return &gRPCServer{
		server:                     grpc.NewServer(grpc.MaxRecvMsgSize(RPCMaxMessageSize), grpc.MaxSendMsgSize(RPCMaxMessageSize)),
		inboundConnectionCount:     0,
		inboundConnectionCountLock: &sync.Mutex{},
	}
}

func RPCServer(mc *domain.Metchain, bc *blockchain.Blockchain) *rpcServer {
	//listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	newGRPCServer := newGRPCServer()
	// register the greet service
	rpcServers := &rpcServer{gRPCServer: *newGRPCServer, Metchain: mc, Blockchain: bc}
	pb.RegisterRPCServer(newGRPCServer.server, rpcServers)

	log.Printf("Server started ats %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	if err := newGRPCServer.server.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
	return rpcServers
}

type grpcStream interface {
	Send(message *pb.MetchainMessage) error
	Recv() (*pb.MetchainMessage, error)
}

func (s *rpcServer) MessageStream(stream pb.RPC_MessageStreamServer) error {

	//da := stream.Context()
	d := stream.Context()
	d.Done()
	req, err := stream.Recv()
	//req.GetPayload()
	if err != nil {
		return err
	}

	if req.GetGetBlockTemplateRequest() != nil {
		template := &pb.MetchainMessage{}
		template.Payload = req.Payload

		reqs := new(pb.MetchainMessage)
		payload := new(pb.MetchainMessage_GetBlockTemplateResponse)
		payload.GetBlockTemplateResponse = s.NewGetBlockTemplateResponse(template.GetGetBlockTemplateRequest())
		reqs.Payload = payload
		err := stream.Send(reqs)
		if err != nil {
			log.Printf("Error Sending Block Info")
		}

	}
	if req.GetSubmitBlockRequest() != nil {
		template := &pb.MetchainMessage{}
		template.Payload = req.Payload

		reqs := new(pb.MetchainMessage)
		payloads := new(pb.MetchainMessage_SubmitBlockResponse)
		payloads.SubmitBlockResponse = s.GetBlockSubmitResponses(template.GetSubmitBlockRequest())
		reqs.Payload = payloads
		err := stream.Send(reqs)
		if err != nil {
			log.Println("Error Sending Block Submit Response")
		}

	}
	if req.GetNotifyNewBlockTemplateRequest() != nil {
		Server_GetNotifyNewBlockTemplateRequest(stream)
	}
	if req.GetGetBlockDagInfoRequest() != nil {
		s.Server_GetGetBlockDagInfoRequest(stream)
	}

	//log.Printf("checking the recivied payload : %v", req.String())
	return nil
}

func (s *rpcServer) NewGetBlockTemplateResponse(response *pb.GetBlockTemplateRequestMessage) *pb.GetBlockTemplateResponseMessage {
	addr := response.PayAddress
	lb := blockchain.GetBlockTemplateBC(s.Metchain, addr)

	return &pb.GetBlockTemplateResponseMessage{
		Block: NewRPCBlock(lb),
	}
}
func NewRPCBlock(lb *blockchain.DomainBlock) *pb.RpcBlock {
	return &pb.RpcBlock{
		Header: NewRPCBlockHeader(lb),
	}
}

func NewRPCBlockHeader(lb *blockchain.DomainBlock) *pb.RpcBlockHeader {
	return &pb.RpcBlockHeader{
		Version:              1,
		HashMerkleRoot:       fmt.Sprintf("%x", lb.PreviousHash),
		Bits:                 uint32(lb.Bits),
		Parents:              NewRPCBlockLevelParents(lb),
		AcceptedIdMerkleRoot: fmt.Sprintf("%x", lb.PreviousHash),
		UtxoCommitment:       lb.UtxoCommitment,
		Timestamp:            lb.Timestamp,
		Nonce:                lb.Nonce,
	}
}
func NewRPCBlockLevelParents(lb *blockchain.DomainBlock) []*pb.RpcBlockLevelParents {
	parents := make([]*pb.RpcBlockLevelParents, 2)
	parents[0] = &pb.RpcBlockLevelParents{
		ParentHashes: []string{fmt.Sprintf("%x", lb.Megablock), fmt.Sprintf("%x", lb.Metblock)},
	}

	return parents
}

func (s *rpcServer) GetBlockSubmitResponses(block *pb.SubmitBlockRequestMessage) *pb.SubmitBlockResponseMessage {

	hash, err := blockchain.CreateMiniBlock(block.Block, s.Dbcon, s.Blockchain)
	ConesusHash := fmt.Sprintf("%x", hash)
	if err != nil {
		return &pb.SubmitBlockResponseMessage{
			Consensushash: ConesusHash,
			RejectReason:  1,
		}
	}
	return &pb.SubmitBlockResponseMessage{
		Consensushash: ConesusHash,
		RejectReason:  0,
	}
}
