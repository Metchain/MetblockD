package addressexchange

import (
	"github.com/Metchain/MetblockD/appmessage"
	"github.com/Metchain/MetblockD/network/addressmanager"
	"github.com/Metchain/MetblockD/protoserver/routerpkg"
	"math/rand"
)

// SendAddressesContext is the interface for the context needed for the SendAddresses flow.
type SendAddressesContext interface {
	AddressManager() *addressmanager.AddressManager
}

func SendAddresses(context SendAddressesContext, incomingRoute *routerpkg.Route, outgoingRoute *routerpkg.Route) error {
	for {
		_, err := incomingRoute.Dequeue()
		if err != nil {
			return err
		}

		addresses := context.AddressManager().Addresses()
		msgAddresses := appmessage.NewMsgAddresses(shuffleAddresses(addresses))

		err = outgoingRoute.Enqueue(msgAddresses)
		if err != nil {
			return err
		}
	}
}

// shuffleAddresses randomizes the given addresses sent if there are more than the maximum allowed in one message.
func shuffleAddresses(addresses []*appmessage.NetAddress) []*appmessage.NetAddress {
	addressCount := len(addresses)

	if addressCount < appmessage.MaxAddressesPerMsg {
		return addresses
	}

	shuffleAddresses := make([]*appmessage.NetAddress, addressCount)
	copy(shuffleAddresses, addresses)

	rand.Shuffle(addressCount, func(i, j int) {
		shuffleAddresses[i], shuffleAddresses[j] = shuffleAddresses[j], shuffleAddresses[i]
	})

	// Truncate it to the maximum size.
	shuffleAddresses = shuffleAddresses[:appmessage.MaxAddressesPerMsg]
	return shuffleAddresses
}
