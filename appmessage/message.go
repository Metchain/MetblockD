package appmessage

import "time"

type MetchainNet uint32

const (
	Mainnet      MetchainNet = 0x3ddcf71d
	MaxInvPerMsg             = 1 << 17
)

// MessageCommand is a number in the header of a message that represents its type.
type MessageCommand uint32

type Message interface {
	Command() MessageCommand
	MessageNumber() uint64
	SetMessageNumber(index uint64)
	ReceivedAt() time.Time
	SetReceivedAt(receivedAt time.Time)
}
