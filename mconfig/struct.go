package mconfig

import (
	"encoding/json"
)

type Genesis struct {
	Timestamp    int64
	Nonce        int
	PreviousHash [32]byte //As per the Hash size
	Merkleroot   []byte
	Height       int64
	Transaction  Txtransaction
}

type Txtransaction struct {
	Txhex      []byte
	Txcoinbase []byte
}

func (tx *Txtransaction) MarshalJSON() ([]byte, error) {
	n, _ := json.Marshal(struct {
		Txhex      []byte `json:"txhex"`
		Txcoinbase []byte `json:"txcoinbase"`
	}{
		Txhex:      tx.Txhex,
		Txcoinbase: tx.Txcoinbase,
	})

	return n, nil
}
