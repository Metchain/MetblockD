package addressmanager

import "github.com/Metchain/Metblock/db/database"

var notBannedAddressBucket = database.MakeBucket([]byte("not-banned-addresses"))
var bannedAddressBucket = database.MakeBucket([]byte("banned-addresses"))

type addressStore struct {
	database           database.Database
	notBannedAddresses map[addressKey]*address
	bannedAddresses    map[ipv6]*address
}
