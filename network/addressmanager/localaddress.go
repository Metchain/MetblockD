package addressmanager

import (
	"github.com/Metchain/Metblock/appmessage"
	"net"
	"sync"
)

type AddressPriority int

type localAddress struct {
	netAddress *appmessage.NetAddress
	score      AddressPriority
}

type localAddressManager struct {
	localAddresses map[addressKey]*localAddress
	lookupFunc     func(string) ([]net.IP, error)
	cfg            *Config
	mutex          sync.Mutex
}
