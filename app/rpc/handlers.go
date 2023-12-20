package rpc

import (
	"github.com/Metchain/MetblockD/app/rpc/rpccontext"
	"github.com/Metchain/MetblockD/app/rpc/rpchandlers"
	"github.com/Metchain/MetblockD/appmessage"
	"github.com/Metchain/MetblockD/protoserver/routerpkg"
)

type handler func(context *rpccontext.Context, router *routerpkg.Router, request appmessage.Message) (appmessage.Message, error)

var handlers = map[appmessage.MessageCommand]handler{
	appmessage.CmdNotifyNewBlockTemplateRequestMessage: rpchandlers.HandleNotifyNewBlockTemplate,
	appmessage.CmdGetBlockDAGInfoRequestMessage:        rpchandlers.HandleGetBlockDAGInfoRequestMessage,
	appmessage.CmdGetBlockTemplateRequestMessage:       rpchandlers.HandleGetBlockTemplateRequestMessage,
	appmessage.CmdSubmitBlockRequestMessage:            rpchandlers.HandleBlockSubmitMessage,
}
