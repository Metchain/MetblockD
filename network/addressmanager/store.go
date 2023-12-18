package addressmanager

import (
	"encoding/binary"
	"github.com/Metchain/Metblock/appmessage"
	"github.com/Metchain/Metblock/db/database"
)

var notBannedAddressBucket = database.MakeBucket([]byte("not-banned-addresses"))
var bannedAddressBucket = database.MakeBucket([]byte("banned-addresses"))

type addressStore struct {
	database           database.Database
	notBannedAddresses map[addressKey]*address
	bannedAddresses    map[ipv6]*address
}

func (as *addressStore) add(key addressKey, address *address) error {
	if _, ok := as.notBannedAddresses[key]; ok {
		return nil
	}

	as.notBannedAddresses[key] = address

	databaseKey := as.notBannedDatabaseKey(key)
	serializedAddress := as.serializeAddress(address)
	return as.database.Put(databaseKey, serializedAddress)
}

func (as *addressStore) notBannedDatabaseKey(key addressKey) *database.Key {
	serializedKey := as.serializeAddressKey(key)
	return notBannedAddressBucket.Key(serializedKey)
}

func (as *addressStore) bannedDatabaseKey(key addressKey) *database.Key {
	return bannedAddressBucket.Key(key.address[:])
}

func (as *addressStore) serializeAddressKey(key addressKey) []byte {
	serializedSize := 16 + 2 // ipv6 + port
	serializedKey := make([]byte, serializedSize)

	copy(serializedKey[:], key.address[:])
	binary.LittleEndian.PutUint16(serializedKey[16:], key.port)

	return serializedKey
}

func (as *addressStore) serializeAddress(address *address) []byte {
	serializedSize := 16 + 2 + 8 + 8 // ipv6 + port + timestamp + connectionFailedCount
	serializedNetAddress := make([]byte, serializedSize)

	copy(serializedNetAddress[:], address.netAddress.IP.To16()[:])
	binary.LittleEndian.PutUint16(serializedNetAddress[16:], address.netAddress.Port)
	binary.LittleEndian.PutUint64(serializedNetAddress[18:], uint64(address.netAddress.Timestamp.UnixMilliseconds()))
	binary.LittleEndian.PutUint64(serializedNetAddress[26:], uint64(address.connectionFailedCount))

	return serializedNetAddress
}

func (as *addressStore) notBannedCount() int {
	return len(as.notBannedAddresses)
}

func (as *addressStore) getAllNotBanned() []*address {
	addresses := make([]*address, 0, len(as.notBannedAddresses))
	for _, address := range as.notBannedAddresses {
		addresses = append(addresses, address)
	}
	return addresses
}

func (am *AddressManager) removeAddressNoLock(address *appmessage.NetAddress) error {
	key := netAddressKey(address)
	return am.store.remove(key)
}

func (as *addressStore) remove(key addressKey) error {
	delete(as.notBannedAddresses, key)

	databaseKey := as.notBannedDatabaseKey(key)
	return as.database.Delete(databaseKey)
}
