package app

import (
	"fmt"
	"github.com/Metchain/Metblock/app/protocol"
	"github.com/Metchain/Metblock/app/rpc"
	"github.com/Metchain/Metblock/commanager"
	"github.com/Metchain/Metblock/domain"
	"github.com/Metchain/Metblock/mconfig/infraconfig"
	"github.com/Metchain/Metblock/network/addressmanager"
	netadapter "github.com/Metchain/Metblock/protoserver"
	"github.com/Metchain/Metblock/utils/panics"
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
