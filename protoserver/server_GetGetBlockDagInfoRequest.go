package protoserver

import (
	pb "github.com/Metchain/Metblock/proto"
	"log"
)

func (s *rpcServer) Server_GetGetBlockDagInfoRequest(stream pb.RPC_MessageStreamServer) {

	log.Println("Get block Dag Info")

	response := new(pb.MetchainMessage)
	payload := new(pb.MetchainMessage_GetBlockDagInfoResponse)
	payload.GetBlockDagInfoResponse = s.GetGetBlockDagInfoResponse()
	response.Payload = payload
	err := stream.Send(response)
	if err != nil {
		log.Printf("Error Sending Block Info")
	}
	//log.Println("Returned data")

}

func (s *rpcServer) GetGetBlockDagInfoResponse() *pb.GetBlockDagInfoResponseMessage {

	return &pb.GetBlockDagInfoResponseMessage{
		Difficulty: float64(s.Cdiff),
		BlockCount: s.Blockheight,
	}
}
