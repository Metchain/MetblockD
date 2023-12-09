package domain

import (
	"github.com/Metchain/Metblock/mconfig/infraconfig"

	"os"
	"path/filepath"
)

func DBReset(cfg *infraconfig.Config) bool {
	err := os.RemoveAll(databasePath(cfg))
	if err == nil {
		log.Infof("Database Successfully Reset")
		return true

	}
	log.Infof("Error reseting db", err)

	// Add exit code with Manual

	return false
}

func databasePath(cfg *infraconfig.Config) string {
	return filepath.Join(cfg.AppDir, cfg.DataDir)
}
