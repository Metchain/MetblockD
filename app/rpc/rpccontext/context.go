package rpccontext

import (
	"github.com/Metchain/Metblock/app/protocol"
	"github.com/Metchain/Metblock/commanager"
	"github.com/Metchain/Metblock/domain"
	"github.com/Metchain/Metblock/mconfig/infraconfig"
	"github.com/Metchain/Metblock/network/addressmanager"
	netadapter "github.com/Metchain/Metblock/protoserver"
)

// Context represents the RPC context
type Context struct {
	Config            *infraconfig.Config
	NetAdapter        *netadapter.NetAdapter
	Domain            domain.Domain
	ProtocolManager   *protocol.Manager
	ConnectionManager *commanager.ConnectionManager
	AddressManager    *addressmanager.AddressManager

	ShutDownChan chan<- struct{}
}

// NewContext creates a new RPC context
func NewContext(cfg *infraconfig.Config,
	domain domain.Domain,
	netAdapter *netadapter.NetAdapter,
	protocolManager *protocol.Manager,
	connectionManager *commanager.ConnectionManager,
	addressManager *addressmanager.AddressManager,

	shutDownChan chan<- struct{}) *Context {

	context := &Context{
		Config:            cfg,
		NetAdapter:        netAdapter,
		Domain:            domain,
		ProtocolManager:   protocolManager,
		ConnectionManager: connectionManager,
		AddressManager:    addressManager,

		ShutDownChan: shutDownChan,
	}

	return context
}
