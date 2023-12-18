package rpchandlers

import (
	"github.com/Metchain/Metblock/app/rpc/rpccontext"
	"github.com/Metchain/Metblock/appmessage"
	"github.com/Metchain/Metblock/external"
	"github.com/Metchain/Metblock/mconfig"
	"github.com/Metchain/Metblock/protoserver/routerpkg"
	"github.com/Metchain/Metblock/utils/bech32"
	"log"
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

	// Add txverifcation here.

	//Build Coinbasedata for new metchain encryption
	//Updated on 12-18-2023
	Coinbasedata := external.AddressToConbaseData(payAddress, getBlockTemplateRequest.ExtraData)

	templateBlock, isNearlySynced, err := context.Domain.MiningManager().GetBlockTemplate(Coinbasedata)

	if isNearlySynced {
		return nil, err
	}
	rpcBlock := appmessage.DomainBlockToRPCBlock(templateBlock)
	isSynced := true
	response := appmessage.NewGetBlockTemplateResponseMessage(rpcBlock, isSynced)
	log.Println("Response:", response)

	return response, nil
}
