package rpchandlers

import (
	"github.com/Metchain/MetblockD/app/rpc/rpccontext"
	"github.com/Metchain/MetblockD/appmessage"
	"github.com/Metchain/MetblockD/external"
	"github.com/Metchain/MetblockD/mconfig"
	"github.com/Metchain/MetblockD/protoserver/routerpkg"
	"github.com/Metchain/MetblockD/utils/bech32"
	"strconv"
)

func HandleGetBlockTemplateRequestMessage(context *rpccontext.Context, router *routerpkg.Router, request appmessage.Message) (appmessage.Message, error) {
	getBlockTemplateRequest := request.(*appmessage.GetBlockTemplateRequestMessage)

	npayAddress := bech32.Encode(mconfig.Bech32PrefixMet.String(), []byte(getBlockTemplateRequest.PayAddress), []byte(strconv.Itoa(0x96e85745)))

	payAddress, err := mconfig.DecodeAddress(npayAddress, context.Config.ActiveNetParams.Prefix)
	if err != nil {
		errorMessage := &appmessage.GetBlockTemplateResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Could not decode address: %s", err)
		return errorMessage, nil
	}

	//Build Coinbasedata for new metchain encryption
	//Updated on 12-18-2023
	Coinbasedata := external.AddressToConbaseData(payAddress, getBlockTemplateRequest.ExtraData)

	templateBlock, _, err := context.Domain.MiningManager().GetBlockTemplate(Coinbasedata)

	rpcBlock, domainblock := appmessage.DomainBlockToRPCBlock(templateBlock)

	//Update this with interface
	context.LastRPCBlock = domainblock

	isSynced := true
	response := appmessage.NewGetBlockTemplateResponseMessage(rpcBlock, isSynced)

	return response, nil
}
