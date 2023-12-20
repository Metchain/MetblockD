package appmessage

import (
	"github.com/Metchain/MetblockD/external"
)

// MsgRequestAddresses implements the Message interface and represents a Metchain
// RequestAddresses message. It is used to request a list of known active peers on the
// network from a peer to help identify potential nodes. The list is returned
// via one or more addr messages (MsgAddresses).
//
// This message has no payload.
type MsgRequestAddresses struct {
	baseMessage
	IncludeAllSubnetworks bool
	SubnetworkID          *external.DomainSubnetworkID
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgRequestAddresses) Command() MessageCommand {
	return CmdRequestAddresses
}

// NewMsgRequestAddresses returns a new Metchain RequestAddresses message that conforms to the
// Message interface. See MsgRequestAddresses for details.
func NewMsgRequestAddresses(includeAllSubnetworks bool, subnetworkID *external.DomainSubnetworkID) *MsgRequestAddresses {
	return &MsgRequestAddresses{
		IncludeAllSubnetworks: includeAllSubnetworks,
		SubnetworkID:          subnetworkID,
	}
}
