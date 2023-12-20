package model

import "github.com/Metchain/MetblockD/external"

// CoinbaseManager exposes methods for handling blocks'
// coinbase transactions
type CoinbaseManager interface {
	ExpectedCoinbaseTransaction(stagingArea *StagingArea, blockHash *external.DomainHash,
		coinbaseData *external.DomainCoinbaseData) (expectedTransaction *external.DomainTransaction, hasRedReward bool, err error)
	CalcBlockSubsidy(stagingArea *StagingArea, blockHash *external.DomainHash) (uint64, error)
	ExtractCoinbaseDataBlueScoreAndSubsidy(coinbaseTx *external.DomainTransaction) (blueScore uint64, coinbaseData *external.DomainCoinbaseData, subsidy uint64, err error)
}
