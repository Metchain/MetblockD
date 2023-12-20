package model

import "github.com/Metchain/MetblockD/external"

type HeadersSelectedChainStore interface {
	Stage(dbContext DBReader, stagingArea *StagingArea, chainChanges *external.SelectedChainPath) error
	IsStaged(stagingArea *StagingArea) bool
	GetIndexByHash(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (uint64, error)
	GetHashByIndex(dbContext DBReader, stagingArea *StagingArea, index uint64) (*external.DomainHash, error)
}
