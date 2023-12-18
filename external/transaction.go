package external

import "bytes"

// DomainOutpoint represents a Metchain transaction outpoint
type DomainOutpoint struct {
	TransactionID DomainTransactionID
	Index         uint32
}

// ByteSlice returns the bytes in this transactionID represented as a byte slice.
// The transactionID bytes are cloned, therefore it is safe to modify the resulting slice.
func (id *DomainTransactionID) ByteSlice() []byte {
	return (*DomainHash)(id).ByteSlice()
}

// Equal returns whether spk equals to other
func (spk *ScriptPublicKey) Equal(other *ScriptPublicKey) bool {
	if spk == nil || other == nil {
		return spk == other
	}

	if string(spk.Version) != string(other.Version) {
		return false
	}

	return bytes.Equal(spk.Script, other.Script)
}
