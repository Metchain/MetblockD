package rpc

import (
	"github.com/Metchain/Metblock/app/rpc/rpccontext"
	"github.com/Metchain/Metblock/app/rpc/rpchandlers"
	"github.com/Metchain/Metblock/appmessage"
	"github.com/Metchain/Metblock/protoserver/routerpkg"
)

type handler func(context *rpccontext.Context, router *routerpkg.Router, request appmessage.Message) (appmessage.Message, error)

var handlers = map[appmessage.MessageCommand]handler{
	appmessage.CmdNotifyNewBlockTemplateRequestMessage: rpchandlers.HandleNotifyNewBlockTemplate,
	appmessage.CmdGetBlockDAGInfoRequestMessage:        rpchandlers.HandleGetBlockDAGInfoRequestMessage,
	appmessage.CmdGetBlockTemplateRequestMessage:       rpchandlers.HandleGetBlockTemplateRequestMessage,
}
