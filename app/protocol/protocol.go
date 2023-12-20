package protocol

import (
	"github.com/Metchain/MetblockD/app/protocol/flowcontext"
	"github.com/Metchain/MetblockD/commanager"
	"github.com/Metchain/MetblockD/mconfig/infraconfig"
	"github.com/Metchain/MetblockD/network/addressmanager"
	netadapter "github.com/Metchain/MetblockD/protoserver"
	"sync"
)

// Manager manages the p2p protocol
type Manager struct {
	context          *flowcontext.FlowContext
	routersWaitGroup sync.WaitGroup
	isClosed         uint32
}

// NewManager creates a new instance of the p2p protocol manager
func NewManager(cfg *infraconfig.Config, netAdapter *netadapter.NetAdapter, addressManager *addressmanager.AddressManager,
	connectionManager *commanager.ConnectionManager) (*Manager, error) {

	manager := Manager{
		context: flowcontext.New(cfg, addressManager, netAdapter, connectionManager),
	}

	netAdapter.SetP2PRouterInitializer(manager.routerInitializer)
	return &manager, nil
}

func (m *Manager) Context() *flowcontext.FlowContext {
	return m.context
}
