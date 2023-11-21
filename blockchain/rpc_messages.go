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
	block := new(DomainBlock)
	block.Height = lb.Height
	block.Timestamp = lb.Timestamp
	block.Nonce = lb.Nonce
	block.PreviousHash = lb.PreviousHash
	block.Metblock = lb.Metblock
	block.Megablock = lb.Megablock
	block.CurrentHash = lb.CurrentHash
	block.UtxoCommitment = addr
	block.Bits = lb.Bits

	return block
}
