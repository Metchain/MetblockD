package addressmanager

import "github.com/Metchain/MetblockD/appmessage"

func IsRoutable(na *appmessage.NetAddress, acceptUnroutable bool) bool {
	if acceptUnroutable {
		return !IsLocal(na)
	}

	return IsValid(na) && !(IsRFC1918(na) || IsRFC2544(na) ||
		IsRFC3927(na) || IsRFC4862(na) || IsRFC3849(na) ||
		IsRFC4843(na) || IsRFC5737(na) || IsRFC6598(na) ||
		IsLocal(na) || (IsRFC4193(na)))
}

// NetAddressKey returns a key of the ip address to use it in maps.
func netAddressKey(netAddress *appmessage.NetAddress) addressKey {
	key := addressKey{port: netAddress.Port}
	// all IPv4 can be represented as IPv6.
	copy(key.address[:], netAddress.IP.To16())
	return key
}

const (
	// GetAddressesMax is the most addresses that we will send in response
	// to a getAddress (in practise the most addresses we will return from a
	// call to AddressCache()).
	GetAddressesMax = 2500
)
