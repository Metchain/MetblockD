package mempool

import (
	"github.com/Metchain/Metblock/blockchain/consensus/consensusreference"
	"github.com/Metchain/Metblock/blockchain/mempool/controller"
	"github.com/Metchain/Metblock/external"
	"github.com/Metchain/Metblock/miningblockbuilder/miningmempoolmodel"
	"sync"
	"time"
)

type mempool struct {
	mtx sync.RWMutex

	config             *Config
	consensusReference consensusreference.ConsensusReference

	mempoolUTXOSet   *mempoolUTXOSet
	transactionsPool *transactionsPool
	rejectedPool     *rejectedPool
}

func New(config *Config, consensusReference consensusreference.ConsensusReference) miningmempoolmodel.Mempool {
	mp := &mempool{
		config:             config,
		consensusReference: consensusReference,
	}

	mp.mempoolUTXOSet = newMempoolUTXOSet(mp)
	mp.transactionsPool = newTransactionsPool(mp)
	mp.rejectedPool = newRejectedPool(mp)

	return mp
}

type mempoolUTXOSet struct {
	mempool                       *mempool
	poolUnspentOutputs            controller.OutpointToUTXOEntryMap
	transactionByPreviousOutpoint controller.OutpointToTransactionMap
}

func newMempoolUTXOSet(mp *mempool) *mempoolUTXOSet {
	return &mempoolUTXOSet{
		mempool:                       mp,
		poolUnspentOutputs:            controller.OutpointToUTXOEntryMap{},
		transactionByPreviousOutpoint: controller.OutpointToTransactionMap{},
	}
}

type transactionsPool struct {
	mempool         *mempool
	allTransactions controller.IDToTransactionMap

	chainedTransactionsByParentID controller.IDToTransactionsSliceMap
	highPriorityTransactions      controller.IDToTransactionMap

	lastExpireScanTime time.Time
}

func newTransactionsPool(mp *mempool) *transactionsPool {
	return &transactionsPool{
		mempool:                       mp,
		allTransactions:               controller.IDToTransactionMap{},
		highPriorityTransactions:      controller.IDToTransactionMap{},
		chainedTransactionsByParentID: controller.IDToTransactionsSliceMap{},

		lastExpireScanTime: time.Now(),
	}
}

type idToRejectedMap map[external.DomainTransactionID]*controller.RejectedTransaction
type previousOutpointToRejectedMap map[external.DomainOutpoint]*controller.RejectedTransaction
type previousInpointToRejectedMap map[external.DomainOutpoint]*controller.RejectedTransaction

type rejectedPool struct {
	mempool                    *mempool
	allRejected                idToRejectedMap
	rejectedByPreviousOutpoint previousOutpointToRejectedMap
	rejectedByPreviousInpoint  previousInpointToRejectedMap
	lastExpireScan             uint64
}

func newRejectedPool(mp *mempool) *rejectedPool {
	return &rejectedPool{
		mempool:                    mp,
		allRejected:                idToRejectedMap{},
		rejectedByPreviousOutpoint: previousOutpointToRejectedMap{},
		rejectedByPreviousInpoint:  previousInpointToRejectedMap{},
		lastExpireScan:             0,
	}
}
