package protoserver

import (
	"github.com/Metchain/Metblock/blockchain"
	"github.com/Metchain/Metblock/domain"
	pb "github.com/Metchain/Metblock/proto"
	"github.com/Metchain/Metblock/utils/logger"
	"google.golang.org/grpc/peer"
	"net"
)

var logp2p = logger.RegisterSubSystem("METD-P2P-Network")

func P2PServer(listen []string, mc *domain.Metchain, bc *blockchain.Blockchain) (*p2pServer, error) {
	//listen on the port

	lis, err := net.Listen("tcp", "localhost"+listen[0])
	if err != nil {
		logp2p.Criticalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	newGRPCServer := newGRPCServer()
	// register the greet service
	p2pServers := &p2pServer{gRPCServer: *newGRPCServer, Metchain: mc, Blockchain: bc}
	pb.RegisterP2PServer(newGRPCServer.server, p2pServers)

	logp2p.Info("Server started at %v", lis.Addr())
	//list is the port, the grpc server needs to start there

	go func() error {
		if err := newGRPCServer.server.Serve(lis); err != nil {

			return err
		}
		return nil
	}()
	if err != nil {
		logp2p.Criticalf("Failed to start: %v", err)
	}
	return p2pServers, nil
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
	logp2p.Infof("Request recieved")
	template := &pb.MetchainMessage{
		Payload: &pb.MetchainMessage_P2PBlockWithTrustedDataResponseMessage{
			P2PBlockWithTrustedDataResponseMessage: s.Blockchain.ConvertGroupToSecureDomainInfo(s.Blockchain.ConvertBlockToDomainGroupBlock(s.Blockchain.ConvertBlockToDomainBlock())),
		},
	}

	err = stream.Send(template)
	if err != nil {
		logp2p.Infof("Error Sending Block Info: ", err)
		return err
	}
	return nil
}
