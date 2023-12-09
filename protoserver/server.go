package protoserver

type Server interface {
	Start() error
	Stop() error
	SetOnConnectedHandler(onConnectedHandler OnConnectedHandler)
	Connect(address string) (Connection, error)
}
