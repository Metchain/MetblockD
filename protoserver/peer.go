package protoserver

import (
	"github.com/Metchain/Metblock/blockchain"
	"github.com/Metchain/Metblock/domain"
	pb "github.com/Metchain/Metblock/proto"
	"google.golang.org/grpc/peer"
	"log"
	"net"
)

func P2PServer(mc *domain.Metchain, bc *blockchain.Blockchain) *p2pServer {
	//listen on the port

	lis, err := net.Listen("tcp", "localhost:14031")
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	newGRPCServer := newGRPCServer()
	// register the greet service
	p2pServers := &p2pServer{gRPCServer: *newGRPCServer, Metchain: mc, Blockchain: bc}
	pb.RegisterP2PServer(newGRPCServer.server, p2pServers)

	log.Printf("Server started at %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	if err := newGRPCServer.server.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
	return p2pServers
}

func (s *p2pServer) MessageStream(stream pb.P2P_MessageStreamServer) error {
	d := stream.Context()

	d.Done()
	p, _ := peer.FromContext(d)
	req, err := stream.Recv()

	if err != nil {
		return err
	}
	if req.GetP2PBlockWithTrustedDataRequestMessage() != nil {

		template := &pb.MetchainMessage{}
		template.Payload = req.Payload

		s.Blockchain.MatchDomainBlockToP2PBlock(template.GetP2PBlockWithTrustedDataRequestMessage(), p.Addr)
	}
	log.Println("Request recieved")
	template := &pb.MetchainMessage{
		Payload: &pb.MetchainMessage_P2PBlockWithTrustedDataResponseMessage{
			P2PBlockWithTrustedDataResponseMessage: s.Blockchain.ConvertGroupToSecureDomainInfo(s.Blockchain.ConvertBlockToDomainGroupBlock(s.Blockchain.ConvertBlockToDomainBlock())),
		},
	}

	err = stream.Send(template)
	if err != nil {
		log.Printf("Error Sending Block Info: ", err)
		return err
	}
	return nil
}
