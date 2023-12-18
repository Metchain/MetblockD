package external

import (
	"bytes"
	"github.com/Metchain/Metblock/mconfig"
)

// DomainCoinbaseData contains data by which a coinbase transaction
// is built
type DomainCoinbaseData struct {
	ScriptPublicKey *ScriptPublicKey
	ExtraData       []byte
}

func (dcd *DomainCoinbaseData) Equal(other *DomainCoinbaseData) bool {
	if dcd == nil || other == nil {
		return dcd == other
	}

	if !bytes.Equal(dcd.ExtraData, other.ExtraData) {
		return false
	}

	return dcd.ScriptPublicKey.Equal(other.ScriptPublicKey)
}

func AddressToConbaseData(address mconfig.Address, data string) *DomainCoinbaseData {
	return &DomainCoinbaseData{
		ScriptPublicKey: &ScriptPublicKey{
			Script:  address.ScriptAddress(),
			Version: address.Version(),
		},
		ExtraData: []byte(data),
	}
}
