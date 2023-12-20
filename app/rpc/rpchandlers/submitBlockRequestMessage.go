package rpchandlers

import (
	"github.com/Metchain/MetblockD/app/rpc/convertor"
	"github.com/Metchain/MetblockD/app/rpc/rpccontext"
	"github.com/Metchain/MetblockD/appmessage"
	"github.com/Metchain/MetblockD/protoserver/routerpkg"
	"github.com/Metchain/MetblockD/utils/logger"
)

func HandleBlockSubmitMessage(context *rpccontext.Context, _ *routerpkg.Router, request appmessage.Message) (appmessage.Message, error) {
	submitBlockRequest := request.(*appmessage.SubmitBlockRequestMessage)

	var err error
	isSynced := true
	// The node is considered synced if it has peers and consensus state is nearly synced

	//Uncomment for Minimum node listing
	/*if context.ProtocolManager.Context().HasPeers() {
		isSynced, err = context.ProtocolManager.Context().IsNearlySynced()
		if err != nil {
			return nil, err
		}
	}*/

	if !context.Config.AllowSubmitBlockWhenNotSynced && !isSynced {
		return &appmessage.SubmitBlockResponseMessage{
			Error:        appmessage.RPCErrorf("Block not submitted - node is not synced"),
			RejectReason: appmessage.RejectReasonIsInIBD,
		}, nil
	}

	domainBlock, err, reward := convertor.RPCBlockToDomainBlock(submitBlockRequest.Block, context)
	if err != nil {
		return &appmessage.SubmitBlockResponseMessage{
			Error:        appmessage.RPCErrorf("Could not parse block: %s", err),
			RejectReason: appmessage.RejectReasonBlockInvalid,
		}, nil
	}

	/*err = context.ProtocolManager.AddBlock(domainBlock)
	if err != nil {
		isProtocolOrRuleError := errors.As(err, &ruleerrors.RuleError{}) || errors.As(err, &protocolerrors.ProtocolError{})
		if !isProtocolOrRuleError {
			return nil, err
		}

		jsonBytes, _ := json.MarshalIndent(submitBlockRequest.Block.Header, "", "    ")
		if jsonBytes != nil {
			log.Warnf("The RPC submitted block triggered a rule/protocol error (%s), printing "+
				"the full header for debug purposes: \n%s", err, string(jsonBytes))
		}

		return &appmessage.SubmitBlockResponseMessage{
			Error:        appmessage.RPCErrorf("Block rejected. Reason: %s", err),
			RejectReason: appmessage.RejectReasonBlockInvalid,
		}, nil
	}*/

	log.Infof("Accepted block %s via submitBlock Reward %v", domainBlock, reward)

	response := appmessage.NewSubmitBlockResponseMessage()
	return response, nil
}

var log = logger.RegisterSubSystem("Met-BlockAcceptence")
