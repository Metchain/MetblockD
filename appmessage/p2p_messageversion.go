package appmessage

import (
	"fmt"
	"github.com/Metchain/Metblock/external"
	"github.com/Metchain/Metblock/protoserver/id"
	"github.com/Metchain/Metblock/utils/mstime"
	"github.com/Metchain/Metblock/version"
	"strings"
)

var DefaultUserAgent = fmt.Sprintf("/Metchaind:%s/", version.Version())

type MsgVersion struct {
	baseMessage
	// Version of the protocol the node is using.
	ProtocolVersion uint32

	// The peer's network (mainnet, testnet, etc.)
	Network string

	// Bitfield which identifies the enabled services.
	Services ServiceFlag

	// Time the message was generated. This is encoded as an int64 on the appmessage.
	Timestamp mstime.Time

	// Address of the local peer.
	Address *NetAddress

	// The peer unique ID
	ID *id.ID

	// The user agent that generated messsage. This is a encoded as a varString
	// on the appmessage. This has a max length of MaxUserAgentLen.
	UserAgent string

	// Don't announce transactions to peer.
	DisableRelayTx bool

	// The subnetwork of the generator of the version message. Should be nil in full nodes
	SubnetworkID *external.DomainSubnetworkID
}

// HasService returns whether the specified service is supported by the peer
// that generated the message.
func (msg *MsgVersion) HasService(service ServiceFlag) bool {
	return msg.Services&service == service
}

// AddService adds service as a supported service by the peer generating the
// message.
func (msg *MsgVersion) AddService(service ServiceFlag) {
	msg.Services |= service
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgVersion) Command() MessageCommand {
	return CmdVersion
}

func NewMsgVersion(addr *NetAddress, id *id.ID, network string,
	subnetworkID *external.DomainSubnetworkID, protocolVersion uint32) *MsgVersion {

	// Limit the timestamp to one millisecond precision since the protocol
	// doesn't support better.
	return &MsgVersion{
		ProtocolVersion: protocolVersion,
		Network:         network,
		Services:        0,
		Timestamp:       mstime.Now(),
		Address:         addr,
		ID:              id,
		UserAgent:       DefaultUserAgent,
		DisableRelayTx:  false,
		SubnetworkID:    subnetworkID,
	}
}

func (msg *MsgVersion) AddUserAgent(name string, version string,
	comments ...string) {

	newUserAgent := fmt.Sprintf("%s:%s", name, version)
	if len(comments) != 0 {
		newUserAgent = fmt.Sprintf("%s(%s)", newUserAgent,
			strings.Join(comments, "; "))
	}
	newUserAgent = fmt.Sprintf("%s%s/", msg.UserAgent, newUserAgent)
	msg.UserAgent = newUserAgent
}
