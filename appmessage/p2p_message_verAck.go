package appmessage

type MsgVerAck struct {
	baseMessage
}

// Message interface.
func NewMsgVerAck() *MsgVerAck {
	return &MsgVerAck{}
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgVerAck) Command() MessageCommand {
	return CmdVerAck
}
