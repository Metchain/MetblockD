package protoserver

import (
	pb "github.com/Metchain/Metblock/proto"
)

func Server_GetNotifyNewBlockTemplateRequest(stream pb.RPC_MessageStreamServer) {
	//log.Println("GetNotifyNewBlockTemplateRequest")
	go func() {
		//log.Printf("Recieved request from user : %v \n", req.GetPayload())
		response := &pb.MetchainMessage{
			Payload: new(pb.MetchainMessage_NotifyNewBlockTemplateResponse),
		}

		err := stream.Send(response)

		if err != nil {
			//log.Println("Error Sending")

		}
		//log.Println("Sent result to user")

	}()

}
