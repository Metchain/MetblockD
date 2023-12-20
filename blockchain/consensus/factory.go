package consensus

import (
	consensusdatabase "github.com/Metchain/MetblockD/blockchain/consensus/database"
	"github.com/Metchain/MetblockD/blockchain/consensus/internal/blockbuilder"
	"github.com/Metchain/MetblockD/db/database"
	"github.com/Metchain/MetblockD/external"
	"github.com/Metchain/MetblockD/mconfig/dagconfig"
	"sync"
)

type Config struct {
	dagconfig.Params
	// IsArchival tells the consensus if it should not prune old blocks
	IsArchival bool
	// EnableSanityCheckPruningUTXOSet checks the full pruning point utxo set against the commitment at every pruning movement
	EnableSanityCheckPruningUTXOSet bool

	SkipAddingGenesis bool
}

type Factory interface {
	NewConsensus(config *Config, db database.Database, consensusEventsChan chan external.ConsensusEvent) (
		external.Consensus, error)
}

const (
	defaultPreallocateCaches = true
)

type factory struct {
	dataDir string

	pastMedianTimeConsructor PastMedianTimeManagerConstructor
	cacheSizeMiB             *int
	preallocateCaches        *bool
}

func (f *factory) NewConsensus(config *Config, db database.Database, consensusEventsChan chan external.ConsensusEvent) (consensusInstance external.Consensus, err error) {
	dbManager := consensusdatabase.New(db)

	blockBuilder := blockbuilder.New(dbManager)

	c := &consensus{
		lock:            &sync.Mutex{},
		databaseContext: dbManager,

		consensusEventsChan: consensusEventsChan,
		virtualNotUpdated:   true,
		blockBuilder:        blockBuilder,
	}
	return c, nil
}
