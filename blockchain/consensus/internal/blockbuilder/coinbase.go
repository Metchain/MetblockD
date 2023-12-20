package blockbuilder

import (
	"github.com/Metchain/MetblockD/app/model"
	"github.com/Metchain/MetblockD/external"
)

func (bb *blockBuilder) newBlockCoinbaseTransaction(stagingArea *model.StagingArea,
	coinbaseData *external.DomainCoinbaseData) (expectedTransaction *external.DomainTransaction, hasRedReward bool, err error) {
	return nil, true, nil
	//return bb.coinbaseManager.ExpectedCoinbaseTransaction(stagingArea, model.VirtualBlockHash, coinbaseData)
}
