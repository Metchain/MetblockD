package addressmanager

import (
	"github.com/Metchain/Metblock/appmessage"
	"net"
	"sync"
)

type AddressManager struct {
	store          *addressStore
	localAddresses *localAddressManager
	mutex          sync.Mutex
	cfg            *Config
}

// addressKey represents a pair of IP and port, the IP is always in V6 representation
type addressKey struct {
	port    uint16
	address ipv6
}

type ipv6 [net.IPv6len]byte

type address struct {
	netAddress            *appmessage.NetAddress
	connectionFailedCount uint64
}
