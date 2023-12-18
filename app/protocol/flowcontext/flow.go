package flowcontext

import (
	peerpkg "github.com/Metchain/Metblock/app/protocol/peer"
	"github.com/Metchain/Metblock/commanager"
	"github.com/Metchain/Metblock/domain"
	"github.com/Metchain/Metblock/external"
	"github.com/Metchain/Metblock/mconfig/infraconfig"
	"github.com/Metchain/Metblock/network/addressmanager"
	netadapter "github.com/Metchain/Metblock/protoserver"
	"github.com/Metchain/Metblock/protoserver/id"
	"github.com/Metchain/Metblock/utils/mstime"
	"sync"
	"time"
)

type FlowContext struct {
	cfg        *infraconfig.Config
	netAdapter *netadapter.NetAdapter

	addressManager    *addressmanager.AddressManager
	connectionManager *commanager.ConnectionManager

	timeStarted int64

	lastRebroadcastTime       time.Time
	onNewBlockTemplateHandler OnNewBlockTemplateHandler
	sharedRequestedBlocks     *SharedRequestedBlocks

	ibdPeer      *peerpkg.Peer
	ibdPeerMutex sync.RWMutex

	peers      map[id.ID]*peerpkg.Peer
	peersMutex sync.RWMutex

	orphans      map[external.DomainHash]*external.DomainBlock
	orphansMutex sync.RWMutex

	shutdownChan chan struct{}
}

func (f *FlowContext) Domain() domain.Domain {
	//TODO implement me
	panic("implement me")
}

func New(cfg *infraconfig.Config, addressManager *addressmanager.AddressManager,
	netAdapter *netadapter.NetAdapter, connectionManager *commanager.ConnectionManager) *FlowContext {

	return &FlowContext{
		cfg:        cfg,
		netAdapter: netAdapter,

		addressManager:        addressManager,
		connectionManager:     connectionManager,
		sharedRequestedBlocks: NewSharedRequestedBlocks(),
		peers:                 make(map[id.ID]*peerpkg.Peer),
		orphans:               make(map[external.DomainHash]*external.DomainBlock),
		timeStarted:           mstime.Now().UnixMilliseconds(),
		shutdownChan:          make(chan struct{}),
	}
}

// NewSharedRequestedBlocks returns a new instance of SharedRequestedBlocks.
func NewSharedRequestedBlocks() *SharedRequestedBlocks {
	return &SharedRequestedBlocks{
		blocks: make(map[external.DomainHash]struct{}),
	}
}

// OnNewBlockTemplateHandler is a handler function that's triggered when a new block template is available
type OnNewBlockTemplateHandler func() error

// SetOnNewBlockTemplateHandler sets the onNewBlockTemplateHandler handler
func (f *FlowContext) SetOnNewBlockTemplateHandler(onNewBlockTemplateHandler OnNewBlockTemplateHandler) {
	f.onNewBlockTemplateHandler = onNewBlockTemplateHandler
}
