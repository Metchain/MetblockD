package rpccontext

import (
	"github.com/Metchain/MetblockD/app/protocol"
	"github.com/Metchain/MetblockD/commanager"
	"github.com/Metchain/MetblockD/domain"
	"github.com/Metchain/MetblockD/external"
	"github.com/Metchain/MetblockD/mconfig/infraconfig"
	"github.com/Metchain/MetblockD/network/addressmanager"
	netadapter "github.com/Metchain/MetblockD/protoserver"
)

// Context represents the RPC context
type Context struct {
	Config            *infraconfig.Config
	NetAdapter        *netadapter.NetAdapter
	Domain            domain.Domain
	ProtocolManager   *protocol.Manager
	ConnectionManager *commanager.ConnectionManager
	AddressManager    *addressmanager.AddressManager
	LastRPCBlock      *external.DomainBlock
	ShutDownChan      chan<- struct{}
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
