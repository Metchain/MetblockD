package addressmanager

import "github.com/Metchain/Metblock/appmessage"

// BestLocalAddress returns the most appropriate local address to use
// for the given remote address.
func (am *AddressManager) BestLocalAddress(remoteAddress *appmessage.NetAddress) *appmessage.NetAddress {
	return am.localAddresses.bestLocalAddress(remoteAddress)
}

// Addresses returns all addresses
func (am *AddressManager) Addresses() []*appmessage.NetAddress {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	return am.store.getAllNotBannedNetAddresses()
}
