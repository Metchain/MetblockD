package miningmanager

import (
	"github.com/Metchain/Metblock/blockchain/consensus/consensusreference"
	"github.com/Metchain/Metblock/blockchain/domainconsensus/miningmanager/miningmodel"
	"github.com/Metchain/Metblock/external"
	"github.com/Metchain/Metblock/miningblockbuilder/miningmempoolmodel"
	"sync"
	"time"
)

type miningManager struct {
	consensusReference   consensusreference.ConsensusReference
	mempool              miningmempoolmodel.Mempool
	blockTemplateBuilder miningmodel.BlockTemplateBuilder
	cachedBlockTemplate  *external.DomainBlockTemplate
	cachingTime          time.Time
	cacheLock            *sync.Mutex
}

func (mm *miningManager) ClearBlockTemplate() {
	//TODO implement me
	panic("implement me")
}

func (mm *miningManager) HandleNewBlockTransactions(txs []*external.DomainTransaction) ([]*external.DomainTransaction, error) {
	//TODO implement me
	panic("implement me")
}

func (mm *miningManager) BlockCandidateTransactions() []*external.DomainTransaction {
	//TODO implement me
	panic("implement me")
}

func (mm *miningManager) ValidateAndInsertTransaction(transaction *external.DomainTransaction, isHighPriority bool, allowOrphan bool) (acceptedTransactions []*external.DomainTransaction, err error) {
	//TODO implement me
	panic("implement me")
}

func (mm *miningManager) RemoveTransactions(txs []*external.DomainTransaction, removeRedeemers bool) error {
	//TODO implement me
	panic("implement me")
}

func (mm *miningManager) GetTransaction(transactionID *external.DomainTransactionID, includeTransactionPool bool, includeOrphanPool bool) (transactionPoolTransaction *external.DomainTransaction, isOrphan bool, found bool) {
	//TODO implement me
	panic("implement me")
}

func (mm *miningManager) GetTransactionsByAddresses(includeTransactionPool bool, includeOrphanPool bool) (sendingInTransactionPool map[string]*external.DomainTransaction, receivingInTransactionPool map[string]*external.DomainTransaction, sendingInOrphanPool map[string]*external.DomainTransaction, receivingInOrphanPool map[string]*external.DomainTransaction, err error) {
	//TODO implement me
	panic("implement me")
}

func (mm *miningManager) AllTransactions(includeTransactionPool bool, includeOrphanPool bool) (transactionPoolTransactions []*external.DomainTransaction, orphanPoolTransactions []*external.DomainTransaction) {
	//TODO implement me
	panic("implement me")
}

func (mm *miningManager) TransactionCount(includeTransactionPool bool, includeOrphanPool bool) int {
	//TODO implement me
	panic("implement me")
}

type MiningManager interface {
	GetBlockTemplate(coinbaseData *external.DomainCoinbaseData) (block *external.DomainBlock, isNearlySynced bool, err error)
	ClearBlockTemplate()
	GetBlockTemplateBuilder() miningmodel.BlockTemplateBuilder
	miningmempoolmodel.Mempool
}

func (mm *miningManager) GetBlockTemplate(coinbaseData *external.DomainCoinbaseData) (block *external.DomainBlock, isNearlySynced bool, err error) {

	mm.cacheLock.Lock()
	immutableCachedTemplate := mm.getImmutableCachedTemplate()
	// We first try and use a cached template
	if immutableCachedTemplate != nil {
		mm.cacheLock.Unlock()

		if immutableCachedTemplate.CoinbaseData.Equal(coinbaseData) {
			return immutableCachedTemplate.Block, immutableCachedTemplate.IsNearlySynced, nil
		}

		// Coinbase data is new -- make the minimum changes required
		// Note we first clone the block template since it is modified by the call
		modifiedBlockTemplate, err := mm.blockTemplateBuilder.ModifyBlockTemplate(coinbaseData)
		if err != nil {
			return nil, false, err
		}

		// No point in updating cache since we have no reason to believe this coinbase will be used more
		// than the previous one, and we want to maintain the original template caching time
		return modifiedBlockTemplate.Block, modifiedBlockTemplate.IsNearlySynced, nil
	}
	defer mm.cacheLock.Unlock()
	// No relevant cache, build a template
	blockTemplate, err := mm.blockTemplateBuilder.BuildBlockTemplate(coinbaseData)
	if err != nil {
		return nil, false, err
	}
	// Cache the built template
	mm.setImmutableCachedTemplate(blockTemplate)

	return blockTemplate.Block, blockTemplate.IsNearlySynced, nil
}

func (mm *miningManager) setImmutableCachedTemplate(blockTemplate *external.DomainBlockTemplate) {
	mm.cachingTime = time.Now()
	mm.cachedBlockTemplate = blockTemplate
}

func (mm *miningManager) getImmutableCachedTemplate() *external.DomainBlockTemplate {
	if time.Since(mm.cachingTime) > time.Second {
		// No point in cache optimizations if queries are more than a second apart -- we prefer rechecking the mempool.
		// Full explanation: On the one hand this is a sub-millisecond optimization, so there is no harm in doing the full block building
		// every ~1 second. Additionally, we would like to refresh the mempool access even if virtual info was
		// unmodified for a while. All in all, caching for max 1 second is a good compromise.
		mm.cachedBlockTemplate = nil
	}
	return mm.cachedBlockTemplate
}

func (mm *miningManager) GetBlockTemplateBuilder() miningmodel.BlockTemplateBuilder {
	return mm.blockTemplateBuilder
}
