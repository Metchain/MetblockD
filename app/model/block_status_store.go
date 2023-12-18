package model

import (
	"github.com/Metchain/Metblock/db/database"
	"github.com/Metchain/Metblock/external"
)

// BlockStatusStore represents a store of BlockStatuses
type BlockStatusStore interface {
	Stage(stagingArea *StagingArea, blockHash *external.DomainHash, blockStatus external.BlockStatus)
	IsStaged(stagingArea *StagingArea) bool
	Get(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (external.BlockStatus, error)
	Exists(dbContext database.Database, stagingArea *StagingArea, blockHash *external.DomainHash) (bool, error)
}
