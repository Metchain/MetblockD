package rpchandlers

import (
	"github.com/Metchain/Metblock/app/rpc/rpccontext"
	"github.com/Metchain/Metblock/appmessage"
	"github.com/Metchain/Metblock/protoserver/routerpkg"
	"log"
)

func HandleGetBlockDAGInfoRequestMessage(context *rpccontext.Context, router *routerpkg.Router, _ appmessage.Message) (appmessage.Message, error) {

	response := appmessage.NewGetBlockDAGInfoResponseMessage()
	log.Println("Response:", response)
	return response, nil
}
