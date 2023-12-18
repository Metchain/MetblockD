package blockbuilder

import (
	"github.com/Metchain/Metblock/app/model"
	"github.com/Metchain/Metblock/external"
)

func (bb *blockBuilder) newBlockCoinbaseTransaction(stagingArea *model.StagingArea,
	coinbaseData *external.DomainCoinbaseData) (expectedTransaction *external.DomainTransaction, hasRedReward bool, err error) {
	return nil, true, nil
	//return bb.coinbaseManager.ExpectedCoinbaseTransaction(stagingArea, model.VirtualBlockHash, coinbaseData)
}
