package addressexchange

import (
	"github.com/Metchain/Metblock/app/protocol/common"
	peerpkg "github.com/Metchain/Metblock/app/protocol/peer"
	"github.com/Metchain/Metblock/app/protocol/protocolerrors"
	"github.com/Metchain/Metblock/appmessage"
	"github.com/Metchain/Metblock/network/addressmanager"
	"github.com/Metchain/Metblock/protoserver/routerpkg"
)

// ReceiveAddressesContext is the interface for the context needed for the ReceiveAddresses flow.
type ReceiveAddressesContext interface {
	AddressManager() *addressmanager.AddressManager
}

// ReceiveAddresses asks a peer for more addresses if needed.
func ReceiveAddresses(context ReceiveAddressesContext, incomingRoute *routerpkg.Route, outgoingRoute *routerpkg.Route,
	peer *peerpkg.Peer) error {

	subnetworkID := peer.SubnetworkID()
	msgGetAddresses := appmessage.NewMsgRequestAddresses(false, subnetworkID)
	err := outgoingRoute.Enqueue(msgGetAddresses)
	if err != nil {
		return err
	}

	message, err := incomingRoute.DequeueWithTimeout(common.DefaultTimeout)
	if err != nil {
		return err
	}

	msgAddresses := message.(*appmessage.MsgAddresses)
	if len(msgAddresses.AddressList) > addressmanager.GetAddressesMax {
		return protocolerrors.Errorf(true, "address count exceeded %d", addressmanager.GetAddressesMax)
	}

	return context.AddressManager().AddAddresses(msgAddresses.AddressList...)
}
