package model

import "github.com/Metchain/MetblockD/external"

// BlockBuilder is responsible for creating blocks from the current state
type BlockBuilder interface {
	BuildBlockTemplate(coinbaseData *external.DomainCoinbaseData, transactions []*external.DomainTransaction) (block *external.DomainBlock, coinbaseHasRedReward bool, err error)
	BuildBlock(coinbaseData *external.TempBlock) (block *external.TempBlock, err error)
}
