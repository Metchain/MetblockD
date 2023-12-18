package appmessage

// NewBlockAddedNotificationMessage returns a instance of the message
func NewBlockAddedNotificationMessage(block *RPCBlock) *BlockAddedNotificationMessage {
	return &BlockAddedNotificationMessage{
		Block: block,
	}
}

// BlockAddedNotificationMessage is an appmessage corresponding to
// its respective RPC message
type BlockAddedNotificationMessage struct {
	baseMessage
	Block *RPCBlock
}
