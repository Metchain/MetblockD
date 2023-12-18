package rpc

import (
	"github.com/Metchain/Metblock/app/protocol"
	"github.com/Metchain/Metblock/app/rpc/rpccontext"
	"github.com/Metchain/Metblock/appmessage"
	"github.com/Metchain/Metblock/commanager"
	"github.com/Metchain/Metblock/domain"
	"github.com/Metchain/Metblock/external"
	"github.com/Metchain/Metblock/mconfig/infraconfig"
	"github.com/Metchain/Metblock/network/addressmanager"
	netadapter "github.com/Metchain/Metblock/protoserver"
	"github.com/Metchain/Metblock/protoserver/routerpkg"
	"github.com/Metchain/Metblock/utils/logger"
	"github.com/pkg/errors"
)

// Manager is an RPC manager
type Manager struct {
	context *rpccontext.Context
}

func NewManager(
	cfg *infraconfig.Config,
	domain domain.Domain,
	netAdapter *netadapter.NetAdapter,
	protocolManager *protocol.Manager,
	connectionManager *commanager.ConnectionManager,
	addressManager *addressmanager.AddressManager,

	shutDownChan chan<- struct{}) *Manager {

	manager := Manager{
		context: rpccontext.NewContext(
			cfg,
			domain,
			netAdapter,
			protocolManager,
			connectionManager,
			addressManager,

			shutDownChan,
		),
	}
	netAdapter.SetRPCRouterInitializer(manager.routerInitializer)

	return &manager
}

func (m *Manager) routerInitializer(router *routerpkg.Router, netConnection *netadapter.NetConnection) {
	messageTypes := make([]appmessage.MessageCommand, 0, len(handlers))
	for messageType := range handlers {
		messageTypes = append(messageTypes, messageType)
	}
	incomingRoute, err := router.AddIncomingRoute("rpc router", messageTypes)
	if err != nil {
		panic(err)
	}

	spawn("routerInitializer-handleIncomingMessages", func() {

		err := m.handleIncomingMessages(router, incomingRoute)
		m.handleError(err, netConnection)
	})
}

func (m *Manager) handleError(err error, netConnection *netadapter.NetConnection) {
	if errors.Is(err, routerpkg.ErrTimeout) {
		log.Warnf("Got timeout from %s. Disconnecting...", netConnection)
		netConnection.Disconnect()
		return
	}
	if errors.Is(err, routerpkg.ErrRouteClosed) {
		return
	}
	panic(err)
}

func (m *Manager) initConsensusEventsHandler(consensusEventsChan chan external.ConsensusEvent) {
	spawn("consensusEventsHandler", func() {
		for {
			consensusEvent, ok := <-consensusEventsChan
			if !ok {
				return
			}
			switch event := consensusEvent.(type) {

			case *external.BlockAdded:
				err := m.notifyBlockAdded(event.Block)
				if err != nil {
					panic(err)
				}
			case *external.BlockSent:
				err := m.notifyBlockAdded(event.Block)
				if err != nil {
					panic(err)
				}
			default:
				panic(errors.Errorf("Got event of unsupported type %T", consensusEvent))
			}
		}
	})
}

// notifyBlockAddedToDAG notifies the manager that a block has been added to the DAG
func (m *Manager) notifyBlockAdded(block *external.DomainBlock) error {
	onEnd := logger.LogAndMeasureExecutionTime(log, "RPCManager.notifyBlockAddedToDAG")
	defer onEnd()

	rpcBlock := appmessage.DomainBlockToRPCBlock(block)
	/*err := m.context.PopulateBlockWithVerboseData(rpcBlock, block.Header, block, true)
	if err != nil {
		return err
	}*/
	blockAddedNotification := appmessage.NewBlockAddedNotificationMessage(rpcBlock)
	log.Infof("Block added:", blockAddedNotification)

	return nil
}

// NotifyNewBlockTemplate notifies the manager that a new
// block template is available for miners
func (m *Manager) NotifyNewBlockTemplate() error {
	appmessage.NewNewBlockTemplateNotificationMessage()
	return nil
}
