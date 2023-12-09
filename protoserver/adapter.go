package protoserver

import (
	"errors"
	"github.com/Metchain/Metblock/appmessage"
	"github.com/Metchain/Metblock/blockchain"
	"github.com/Metchain/Metblock/domain"
	"github.com/Metchain/Metblock/mconfig/infraconfig"

	"github.com/Metchain/Metblock/protoserver/id"
	"github.com/Metchain/Metblock/protoserver/routerpkg"
	"sync"
	"sync/atomic"
)

type RouterInitializer func(*routerpkg.Router, *NetConnection)

// NetConnection is a wrapper to a server connection for use by services external to NetAdapter
type NetConnection struct {
	connection            Connection
	id                    *id.ID
	router                *routerpkg.Router
	onDisconnectedHandler OnDisconnectedHandler
	isRouterClosed        uint32
}

// NetAdapter is an abstraction layer over networking.
// This type expects a RouteInitializer function. This
// function weaves together the various "routes" (messages
// and message handlers) without exposing anything related
// to networking internals.
type NetAdapter struct {
	cfg                  *infraconfig.Config
	id                   *id.ID
	p2pServer            *p2pServer
	p2pRouterInitializer RouterInitializer
	rpcServer            *rpcServer
	rpcRouterInitializer RouterInitializer
	stop                 uint32

	p2pConnections     map[*NetConnection]struct{}
	p2pConnectionsLock sync.RWMutex
}

// NewNetAdapter creates and starts a new NetAdapter on the
// given listeningPort
func NewNetAdapter(cfg *infraconfig.Config, mc *domain.Metchain, bc *blockchain.Blockchain) (*NetAdapter, error) {
	netAdapterID, err := id.GenerateID()
	if err != nil {
		return nil, err
	}
	p2pServer, err := P2PServer(cfg.Listeners, mc, bc)
	if err != nil {
		return nil, err
	}
	rpcServer, err := RPCServer(cfg.RPCListeners, cfg.RPCMaxClients, mc, bc)
	if err != nil {
		return nil, err
	}
	adapter := NetAdapter{
		cfg:       cfg,
		id:        netAdapterID,
		p2pServer: p2pServer,
		rpcServer: rpcServer,

		p2pConnections: make(map[*NetConnection]struct{}),
	}

	adapter.p2pServer.SetOnConnectedHandler(adapter.onP2PConnectedHandler)
	adapter.rpcServer.SetOnConnectedHandler(adapter.onRPCConnectedHandler)

	return &adapter, nil
}

// Start begins the operation of the NetAdapter
func (na *NetAdapter) Start() error {
	if na.p2pRouterInitializer == nil {
		return errors.New("p2pRouterInitializer was not set")
	}
	if na.rpcRouterInitializer == nil {
		return errors.New("rpcRouterInitializer was not set")
	}

	err := na.p2pServer.Start()
	if err != nil {
		return err
	}
	err = na.rpcServer.Start()
	if err != nil {
		return err
	}

	return nil
}

// Stop safely closes the NetAdapter
func (na *NetAdapter) Stop() error {
	if atomic.AddUint32(&na.stop, 1) != 1 {
		return errors.New("net adapter stopped more than once")
	}
	err := na.p2pServer.Stop()
	if err != nil {
		return err
	}
	return na.rpcServer.Stop()
}

// P2PConnect tells the NetAdapter's underlying p2p server to initiate a connection
// to the given address
func (na *NetAdapter) P2PConnect(address string) error {
	_, err := na.p2pServer.Connect(address)
	return err
}

// P2PConnections returns a list of p2p connections currently connected and active
func (na *NetAdapter) P2PConnections() []*NetConnection {
	na.p2pConnectionsLock.RLock()
	defer na.p2pConnectionsLock.RUnlock()

	netConnections := make([]*NetConnection, 0, len(na.p2pConnections))

	for netConnection := range na.p2pConnections {
		netConnections = append(netConnections, netConnection)
	}

	return netConnections
}

// P2PConnectionCount returns the count of the connected p2p connections
func (na *NetAdapter) P2PConnectionCount() int {
	na.p2pConnectionsLock.RLock()
	defer na.p2pConnectionsLock.RUnlock()

	return len(na.p2pConnections)
}

func (na *NetAdapter) onP2PConnectedHandler(connection Connection) error {
	netConnection := newNetConnection(connection, na.p2pRouterInitializer, "on P2P connected")

	na.p2pConnectionsLock.Lock()
	defer na.p2pConnectionsLock.Unlock()

	netConnection.setOnDisconnectedHandler(func() {
		na.p2pConnectionsLock.Lock()
		defer na.p2pConnectionsLock.Unlock()

		delete(na.p2pConnections, netConnection)
	})

	na.p2pConnections[netConnection] = struct{}{}

	netConnection.start()

	return nil
}

func (na *NetAdapter) onRPCConnectedHandler(connection Connection) error {
	netConnection := newNetConnection(connection, na.rpcRouterInitializer, "on RPC connected")
	netConnection.setOnDisconnectedHandler(func() {})
	netConnection.start()

	return nil
}

// SetP2PRouterInitializer sets the p2pRouterInitializer function
// for the net adapter
func (na *NetAdapter) SetP2PRouterInitializer(routerInitializer RouterInitializer) {
	na.p2pRouterInitializer = routerInitializer
}

// SetRPCRouterInitializer sets the rpcRouterInitializer function
// for the net adapter
func (na *NetAdapter) SetRPCRouterInitializer(routerInitializer RouterInitializer) {
	na.rpcRouterInitializer = routerInitializer
}

// ID returns this netAdapter's ID in the network
func (na *NetAdapter) ID() *id.ID {
	return na.id
}

// P2PBroadcast sends the given `message` to every peer corresponding
// to each NetConnection in the given netConnections
func (na *NetAdapter) P2PBroadcast(netConnections []*NetConnection, message appmessage.Message) error {
	na.p2pConnectionsLock.RLock()
	defer na.p2pConnectionsLock.RUnlock()

	for _, netConnection := range netConnections {
		err := netConnection.router.OutgoingRoute().Enqueue(message)
		if err != nil {
			if errors.Is(err, routerpkg.ErrRouteClosed) {
				log.Debugf("Cannot enqueue message to %s: router is closed", netConnection)
				continue
			}
			return err
		}
	}
	return nil
}