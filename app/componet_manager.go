package app

import (
	"fmt"
	"github.com/Metchain/MetblockD/app/protocol"
	"github.com/Metchain/MetblockD/app/rpc"
	"github.com/Metchain/MetblockD/commanager"
	"github.com/Metchain/MetblockD/domain"
	"github.com/Metchain/MetblockD/mconfig/infraconfig"
	"github.com/Metchain/MetblockD/network/addressmanager"
	netadapter "github.com/Metchain/MetblockD/protoserver"
	"github.com/Metchain/MetblockD/utils/panics"
	"sync/atomic"
)

type ComponentManager struct {
	cfg               *infraconfig.Config
	addressManager    *addressmanager.AddressManager
	protocolManager   *protocol.Manager
	rpcManager        *rpc.Manager
	connectionManager *commanager.ConnectionManager
	netAdapter        *netadapter.NetAdapter

	started, shutdown int32
}

func (a *ComponentManager) Start() {
	// Already started?
	if atomic.AddInt32(&a.started, 1) != 1 {
		return
	}

	log.Trace("Starting MetchainD")

	err := a.netAdapter.Start()
	if err != nil {
		panics.Exit(log, fmt.Sprintf("Error starting the net adapter: %+v", err))
	}

	a.connectionManager.Start()
}

func setupRPC(
	cfg *infraconfig.Config,
	domain domain.Domain,
	netAdapter *netadapter.NetAdapter,
	protocolManager *protocol.Manager,
	connectionManager *commanager.ConnectionManager,
	addressManager *addressmanager.AddressManager,

	shutDownChan chan<- struct{},
) *rpc.Manager {

	rpcManager := rpc.NewManager(
		cfg,
		domain,
		netAdapter,
		protocolManager,
		connectionManager,
		addressManager,
		shutDownChan,
	)
	protocolManager.SetOnNewBlockTemplateHandler(rpcManager.NotifyNewBlockTemplate)

	return rpcManager
}
