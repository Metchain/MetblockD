package model

import "github.com/Metchain/Metblock/external"

// BlockStore represents a store of blocks
type BlockStore interface {
	Stage(stagingArea *StagingArea, blockHash *external.DomainHash, block *external.DomainBlock)
	IsStaged(stagingArea *StagingArea) bool
	Block(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (*external.DomainBlock, error)
	HasBlock(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (bool, error)
	Blocks(dbContext DBReader, stagingArea *StagingArea, blockHashes []*external.DomainHash) ([]*external.DomainBlock, error)
	Delete(stagingArea *StagingArea, blockHash *external.DomainHash)
	Count(stagingArea *StagingArea) uint64
	AllBlockHashesIterator(dbContext DBReader) (BlockIterator, error)
}
