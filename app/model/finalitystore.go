package model

import "github.com/Metchain/Metblock/external"

// UTXODiffStore represents a store of UTXODiffs
type UTXODiffStore interface {
	Stage(stagingArea *StagingArea, blockHash *external.DomainHash, utxoDiff external.UTXODiff, utxoDiffChild *external.DomainHash)
	IsStaged(stagingArea *StagingArea) bool
	UTXODiff(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (external.UTXODiff, error)
	UTXODiffChild(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (*external.DomainHash, error)
	HasUTXODiffChild(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (bool, error)
	Delete(stagingArea *StagingArea, blockHash *external.DomainHash)
}

// FinalityStore represents a store for finality data
type FinalityStore interface {
	IsStaged(stagingArea *StagingArea) bool
	StageFinalityPoint(stagingArea *StagingArea, blockHash *external.DomainHash, finalityPointHash *external.DomainHash)
	FinalityPoint(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (*external.DomainHash, error)
}
