package metflow

import (
	"github.com/Metchain/MetblockD/app/protocol/common"
	"github.com/Metchain/MetblockD/app/protocol/flowcontext"
	"github.com/Metchain/MetblockD/appmessage"
	"github.com/Metchain/MetblockD/protoserver/routerpkg"
)

type protocolManager interface {
	RegisterFlow(name string, router *routerpkg.Router, messageTypes []appmessage.MessageCommand, isStopping *uint32,
		errChan chan error, initializeFunc common.FlowInitializeFunc) *common.Flow
	RegisterOneTimeFlow(name string, router *routerpkg.Router, messageTypes []appmessage.MessageCommand,
		isStopping *uint32, stopChan chan error, initializeFunc common.FlowInitializeFunc) *common.Flow
	RegisterFlowWithCapacity(name string, capacity int, router *routerpkg.Router,
		messageTypes []appmessage.MessageCommand, isStopping *uint32,
		errChan chan error, initializeFunc common.FlowInitializeFunc) *common.Flow
	Context() *flowcontext.FlowContext
}

func Register(m protocolManager, router *routerpkg.Router, errChan chan error, isStopping *uint32) (flows []*common.Flow) {
	flows = registerAddressFlows(m, router, isStopping, errChan)
	/*flows = append(flows, registerBlockRelayFlows(m, router, isStopping, errChan)...)
	flows = append(flows, registerPingFlows(m, router, isStopping, errChan)...)
	flows = append(flows, registerTransactionRelayFlow(m, router, isStopping, errChan)...)
	flows = append(flows, registerRejectsFlow(m, router, isStopping, errChan)...)*/

	return flows
}
