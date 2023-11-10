package main

import "encoding/json"

type MBlock struct {
	Height       uint64
	Timestamp    int64
	Nonce        uint64
	PreviousHash [32]byte //As per the Hash size
	Megablock    [32]byte
	Metblock     [32]byte
	Transactions []*Transaction
	CurrentHash  [32]byte
	Bits         uint64
}

func (b *MBlock) UnmarshalJSON(data []byte) error {

	v := &struct {
		Height       uint64         `json:"height"`
		Timestamp    int64          `json:"timestamp"`
		Nonce        uint64         `json:"nonce"`
		PreviousHash [32]byte       `json:"previousHash"` //As per the Hash size
		Megablock    [32]byte       `json:"megablock"`
		Metblock     [32]byte       `json:"metblock"`
		Transactions []*Transaction `json:"transactions"`
		CurrentHash  [32]byte       `json:"currentHash"`
		Bits         uint64         `json:"bits"`
	}{
		Height:       b.Height,
		Timestamp:    b.Timestamp,
		Nonce:        b.Height,
		PreviousHash: b.PreviousHash,
		Megablock:    b.Megablock,
		Metblock:     b.Metblock,
		Transactions: b.Transactions,
		CurrentHash:  b.CurrentHash,
		Bits:         b.Bits,
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	b.Height = v.Height
	b.Timestamp = v.Timestamp
	b.Nonce = v.Height
	b.PreviousHash = v.PreviousHash
	b.Megablock = v.Megablock
	b.Metblock = v.Metblock
	b.Transactions = v.Transactions
	b.CurrentHash = v.CurrentHash
	b.Bits = v.Bits

	return nil
}
