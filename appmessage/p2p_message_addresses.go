package appmessage

const MaxAddressesPerMsg = 1000

// MsgAddresses implements the Message interface and represents a Metchain
// Addresses message.
type MsgAddresses struct {
	baseMessage
	AddressList []*NetAddress
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgAddresses) Command() MessageCommand {
	return CmdAddresses
}

func NewMsgAddresses(addressList []*NetAddress) *MsgAddresses {
	return &MsgAddresses{
		AddressList: addressList,
	}
}
