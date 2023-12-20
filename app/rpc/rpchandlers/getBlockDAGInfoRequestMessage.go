package rpchandlers

import (
	"github.com/Metchain/MetblockD/app/rpc/rpccontext"
	"github.com/Metchain/MetblockD/appmessage"
	"github.com/Metchain/MetblockD/protoserver/routerpkg"
	nlog "log"
)

func HandleGetBlockDAGInfoRequestMessage(context *rpccontext.Context, router *routerpkg.Router, _ appmessage.Message) (appmessage.Message, error) {

	response := appmessage.NewGetBlockDAGInfoResponseMessage()
	nlog.Println("Response:", response)
	return response, nil
}
