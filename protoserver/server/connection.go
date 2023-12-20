package server

import (
	"fmt"
	"github.com/Metchain/MetblockD/protoserver/routerpkg"
	"net"
)

// Connection represents a server connection.
type Connection interface {
	fmt.Stringer
	Start(router *routerpkg.Router)
	Disconnect()
	IsConnected() bool
	IsOutbound() bool
	SetOnDisconnectedHandler(onDisconnectedHandler OnDisconnectedHandler)
	SetOnInvalidMessageHandler(onInvalidMessageHandler OnInvalidMessageHandler)
	Address() *net.TCPAddr
}

// OnDisconnectedHandler is a function that is to be
// called once a Connection has been disconnected.
type OnDisconnectedHandler func()

// OnInvalidMessageHandler is a function that is to be called when
// an invalid message (cannot be parsed/doesn't have a route)
// was received from a connection.
type OnInvalidMessageHandler func(err error)
