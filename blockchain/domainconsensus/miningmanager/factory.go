package miningmanager

import (
	"github.com/Metchain/Metblock/blockchain/consensus/consensusreference"
	"github.com/Metchain/Metblock/blockchain/mempool"
	"github.com/Metchain/Metblock/mconfig/dagconfig"
	"github.com/Metchain/Metblock/miningblockbuilder"
	"sync"
	"time"
)

// Factory instantiates new mining managers
type Factory interface {
	NewMiningManager(consensus consensusreference.ConsensusReference, params *dagconfig.Params, mempoolConfig *mempool.Config) MiningManager
}

type factory struct{}

// NewMiningManager instantiate a new mining manager
func (f *factory) NewMiningManager(consensusReference consensusreference.ConsensusReference, params *dagconfig.Params,
	mempoolConfig *mempool.Config) MiningManager {

	mempool := mempool.New(mempoolConfig, consensusReference)
	blockTemplateBuilder := miningblockbuilder.New(consensusReference, mempool)

	return &miningManager{
		consensusReference:   consensusReference,
		mempool:              mempool,
		blockTemplateBuilder: blockTemplateBuilder,
		cachingTime:          time.Time{},
		cacheLock:            &sync.Mutex{},
	}
}

// NewFactory creates a new mining manager factory
func NewFactory() Factory {
	return &factory{}
}
