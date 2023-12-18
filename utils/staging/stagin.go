package staging

import (
	"github.com/Metchain/Metblock/app/model"
	"github.com/Metchain/Metblock/db/database"
	"github.com/Metchain/Metblock/utils/logger"
	"sync/atomic"
)

var lastShardingID uint64

// CommitAllChanges creates a transaction in `databaseContext`, and commits all changes in `stagingArea` through it.
func CommitAllChanges(databaseContext database.Database, stagingArea *model.StagingArea) error {
	onEnd := logger.LogAndMeasureExecutionTime(utilLog, "commitAllChanges")
	defer onEnd()

	dbTx, err := databaseContext.Begin()
	if err != nil {
		return err
	}

	err = stagingArea.Commit(dbTx)
	if err != nil {
		return err
	}

	return dbTx.Commit()
}

// GenerateShardingID generates a unique staging sharding ID.
func GenerateShardingID() model.StagingShardID {
	return model.StagingShardID(atomic.AddUint64(&lastShardingID, 1))
}
