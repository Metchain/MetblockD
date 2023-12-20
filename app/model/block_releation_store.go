package model

import "github.com/Metchain/MetblockD/external"

type BlockRelationStore interface {
	StageBlockRelation(stagingArea *StagingArea, blockHash *external.DomainHash, blockRelations *BlockRelations)
	IsStaged(stagingArea *StagingArea) bool
	BlockRelation(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (*BlockRelations, error)
	Has(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (bool, error)
	UnstageAll(stagingArea *StagingArea)
}
