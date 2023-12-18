package server

type OnConnectedHandler func(connection Connection) error

// P2PServer represents a p2p server.
type P2PServer interface {
	Server
	Connect(address string) (Connection, error)
}

// Server represents a server.
type Server interface {
	Start() error
	Stop() error
	SetOnConnectedHandler(onConnectedHandler OnConnectedHandler)
}
