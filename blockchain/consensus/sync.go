package consensus

import (
	"github.com/Metchain/Metblock/db/database"
	"github.com/Metchain/Metblock/domain"
	"github.com/Metchain/Metblock/mconfig/infraconfig"

	"os"
	"path/filepath"
)

type status bool

func Sync(gc []byte, db database.Database, cfg *infraconfig.Config) *domain.Metchain {
	metch := checkDB(gc, db, cfg)

	return metch
}
func checkDB(gc []byte, db database.Database, cfg *infraconfig.Config) *domain.Metchain {
	s := checkDBDir(cfg)
	if !s {
		db := domain.GenesisConsensusDBCreate(gc, db)
		if !db {
			log.Infof("Genesis Block Created")

		}
	}
	metch := domain.VerifyGenesisConsensusDB(gc, db, cfg)

	return metch

}

func checkDBDir(cfg *infraconfig.Config) status {
	dir := databasePath(cfg)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Infof(dir, "does not exist")
		return false
	} else {
		log.Infof("The provided directory named", dir, "exists")
		return true
	}
}

func databasePath(cfg *infraconfig.Config) string {
	return filepath.Join(cfg.AppDir, cfg.DataDir)
}
