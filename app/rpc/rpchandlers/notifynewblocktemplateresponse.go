package rpchandlers

import (
	"github.com/Metchain/Metblock/app/rpc/rpccontext"
	"github.com/Metchain/Metblock/appmessage"
	"github.com/Metchain/Metblock/protoserver/routerpkg"
)

func HandleNotifyNewBlockTemplate(context *rpccontext.Context, router *routerpkg.Router, _ appmessage.Message) (appmessage.Message, error) {

	response := appmessage.NewNotifyNewBlockTemplateResponseMessage()
	return response, nil
}
