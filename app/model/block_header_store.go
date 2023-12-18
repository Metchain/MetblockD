package model

import "github.com/Metchain/Metblock/external"

// BlockHeaderStore represents a store of block headers
type BlockHeaderStore interface {
	Stage(stagingArea *StagingArea, blockHash *external.DomainHash, blockHeader external.BlockHeader)
	IsStaged(stagingArea *StagingArea) bool
	BlockHeader(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (external.BlockHeader, error)
	HasBlockHeader(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (bool, error)
	BlockHeaders(dbContext DBReader, stagingArea *StagingArea, blockHashes []*external.DomainHash) ([]external.BlockHeader, error)
	Delete(stagingArea *StagingArea, blockHash *external.DomainHash)
	Count(stagingArea *StagingArea) uint64
}
