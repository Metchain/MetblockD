package appmessage

import (
	"github.com/Metchain/Metblock/utils/mstime"
	"net"
)

type NetAddress struct {
	// Last time the address was seen.
	Timestamp mstime.Time

	// IP address of the peer.
	IP net.IP

	// Port the peer is using. This is encoded in big endian on the appmessage
	// which differs from most everything else.
	Port uint16
}
