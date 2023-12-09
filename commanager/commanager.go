package commanager

import (
	"github.com/Metchain/Metblock/mconfig/infraconfig"
	"github.com/Metchain/Metblock/network/addressmanager"
	"github.com/Metchain/Metblock/protoserver"
	"sync"
	"time"
)

// connectionRequest represents a user request (either through CLI or RPC) to connect to a certain node
type connectionRequest struct {
	address       string
	isPermanent   bool
	nextAttempt   time.Time
	retryDuration time.Duration
}

// ConnectionManager monitors that the current active connections satisfy the requirements of
// outgoing, requested and incoming connections
type ConnectionManager struct {
	cfg            *infraconfig.Config
	netAdapter     *protoserver.NetAdapter
	addressManager *addressmanager.AddressManager

	activeRequested  map[string]*connectionRequest
	pendingRequested map[string]*connectionRequest
	activeOutgoing   map[string]struct{}
	targetOutgoing   int
	activeIncoming   map[string]struct{}
	maxIncoming      int

	stop                   uint32
	connectionRequestsLock sync.RWMutex

	resetLoopChan chan struct{}
	loopTicker    *time.Ticker
}

const connectionsLoopInterval = 30 * time.Second

// New instantiates a new instance of a ConnectionManager
func New(cfg *infraconfig.Config, netAdapter *protoserver.NetAdapter, addressManager *addressmanager.AddressManager) (*ConnectionManager, error) {
	c := &ConnectionManager{
		cfg:              cfg,
		netAdapter:       netAdapter,
		addressManager:   addressManager,
		activeRequested:  map[string]*connectionRequest{},
		pendingRequested: map[string]*connectionRequest{},
		activeOutgoing:   map[string]struct{}{},
		activeIncoming:   map[string]struct{}{},
		resetLoopChan:    make(chan struct{}),
		loopTicker:       time.NewTicker(connectionsLoopInterval),
	}

	connectPeers := cfg.AddPeers
	if len(cfg.ConnectPeers) > 0 {
		connectPeers = cfg.ConnectPeers
	}

	c.maxIncoming = cfg.MaxInboundPeers
	c.targetOutgoing = cfg.TargetOutboundPeers

	for _, connectPeer := range connectPeers {
		c.pendingRequested[connectPeer] = &connectionRequest{
			address:     connectPeer,
			isPermanent: true,
		}
	}

	return c, nil
}
