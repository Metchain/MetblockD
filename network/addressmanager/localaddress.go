package addressmanager

import (
	"github.com/Metchain/Metblock/appmessage"
	"github.com/Metchain/Metblock/utils/mstime"
	"github.com/pkg/errors"
	"net"
	"sync"
	"time"
)

type AddressPriority int

type localAddress struct {
	netAddress *appmessage.NetAddress
	score      AddressPriority
}

type localAddressManager struct {
	localAddresses map[addressKey]*localAddress
	lookupFunc     func(string) ([]net.IP, error)
	cfg            *Config
	mutex          sync.Mutex
}

// IsBanned returns true if the given address is marked as banned
func (am *AddressManager) IsBanned(address *appmessage.NetAddress) (bool, error) {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	key := netAddressKey(address)
	err := am.unbanIfOldEnough(key)
	if err != nil {
		return false, err
	}
	if !am.store.isBanned(key) {
		if !am.store.isNotBanned(key) {
			return false, errors.Wrapf(ErrAddressNotFound, "address %s "+
				"is not registered with the address manager", address.TCPAddress())
		}
		return false, nil
	}

	return true, nil
}

func (as *addressStore) isBanned(key addressKey) bool {
	_, ok := as.bannedAddresses[key.address]
	return ok
}

func (am *AddressManager) unbanIfOldEnough(key addressKey) error {
	address, ok := am.store.getBanned(key)
	if !ok {
		return nil
	}

	const maxBanTime = 24 * time.Hour
	if mstime.Since(address.netAddress.Timestamp) > maxBanTime {
		err := am.store.removeBanned(key)
		if err != nil {
			return err
		}
	}
	return nil
}

func (as *addressStore) getBanned(key addressKey) (*address, bool) {
	bannedAddress, ok := as.bannedAddresses[key.address]
	return bannedAddress, ok
}

func (as *addressStore) isNotBanned(key addressKey) bool {
	_, ok := as.notBannedAddresses[key]
	return ok
}

func (as *addressStore) removeBanned(key addressKey) error {
	delete(as.bannedAddresses, key.address)

	databaseKey := as.bannedDatabaseKey(key)
	return as.database.Delete(databaseKey)
}

// Ban marks the given address as banned
func (am *AddressManager) Ban(addressToBan *appmessage.NetAddress) error {
	am.mutex.Lock()
	defer am.mutex.Unlock()

	keyToBan := netAddressKey(addressToBan)
	keysToDelete := make([]addressKey, 0)
	for _, address := range am.store.getAllNotBannedNetAddresses() {
		key := netAddressKey(address)
		if key.address.equal(keyToBan.address) {
			keysToDelete = append(keysToDelete, key)
		}
	}
	for _, key := range keysToDelete {
		err := am.store.remove(key)
		if err != nil {
			return err
		}
	}

	address := &address{netAddress: addressToBan}
	return am.store.addBanned(keyToBan, address)
}

func (as *addressStore) getAllNotBannedNetAddresses() []*appmessage.NetAddress {
	addresses := make([]*appmessage.NetAddress, 0, len(as.notBannedAddresses))
	for _, address := range as.notBannedAddresses {
		addresses = append(addresses, address.netAddress)
	}
	return addresses
}

func (i ipv6) equal(other ipv6) bool {
	return i == other
}

func (as *addressStore) addBanned(key addressKey, address *address) error {
	if _, ok := as.bannedAddresses[key.address]; ok {
		return nil
	}

	as.bannedAddresses[key.address] = address

	databaseKey := as.bannedDatabaseKey(key)
	serializedAddress := as.serializeAddress(address)
	return as.database.Put(databaseKey, serializedAddress)
}

func (lam *localAddressManager) bestLocalAddress(remoteAddress *appmessage.NetAddress) *appmessage.NetAddress {
	lam.mutex.Lock()
	defer lam.mutex.Unlock()

	bestReach := 0
	var bestScore AddressPriority
	var bestAddress *appmessage.NetAddress
	for _, localAddress := range lam.localAddresses {
		reach := reachabilityFrom(localAddress.netAddress, remoteAddress, lam.cfg.AcceptUnroutable)
		if reach > bestReach ||
			(reach == bestReach && localAddress.score > bestScore) {
			bestReach = reach
			bestScore = localAddress.score
			bestAddress = localAddress.netAddress
		}
	}

	if bestAddress == nil {
		// Send something unroutable if nothing suitable.
		var ip net.IP
		if !IsIPv4(remoteAddress) {
			ip = net.IPv6zero
		} else {
			ip = net.IPv4zero
		}
		bestAddress = appmessage.NewNetAddressIPPort(ip, 0)
	}

	return bestAddress
}
