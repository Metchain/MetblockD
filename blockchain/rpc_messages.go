package blockchain

import (
	"github.com/Metchain/Metblock/domain"
)

type DomainBlock struct {
	Height         uint64
	Timestamp      int64
	Nonce          uint64
	PreviousHash   [32]byte //As per the Hash size
	Megablock      [32]byte
	Metblock       [32]byte
	Transactions   []*Transaction
	CurrentHash    [32]byte
	UtxoCommitment string
	Bits           uint64
}

func GetBlockTemplateBC(mc *domain.Metchain, addr string) *DomainBlock {
	lb := LastMiniBlock(mc)

	return &DomainBlock{
		Height:         lb.Height,
		Timestamp:      lb.Timestamp,
		Nonce:          lb.Nonce,
		PreviousHash:   lb.PreviousHash,
		Metblock:       lb.Metblock,
		Megablock:      lb.Megablock,
		CurrentHash:    lb.CurrentHash,
		UtxoCommitment: addr,
		Bits:           lb.Bits,
	}
}
