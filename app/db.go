package app

import (
	"fmt"
	"github.com/Metchain/MetblockD/app/execenv"
	"github.com/Metchain/MetblockD/db/database"
	"github.com/Metchain/MetblockD/db/database/ldb"
	"github.com/Metchain/MetblockD/mconfig/infraconfig"
	"github.com/Metchain/MetblockD/utils/logger"
	"github.com/Metchain/MetblockD/utils/panics"
	"github.com/pkg/errors"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func StartApp() error {
	execenv.Initialize(desiredLimits)
	cfg, err := infraconfig.LoadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	defer logger.BackendLog.Close()
	defer panics.HandlePanic(log, "MAIN", nil)

	app := &metchainApp{cfg: cfg}

	return app.main(nil)

}

func openDB(cfg *infraconfig.Config) (database.Database, error) {
	dbPath := databasePath(cfg)

	err := checkDatabaseVersion(dbPath)
	if err != nil {
		return nil, err
	}

	log.Infof("Loading database from '%s'", dbPath)
	db, err := ldb.NewLevelDB(dbPath, leveldbCacheSizeMiB)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// dbPath returns the path to the block database given a database type.
func databasePath(cfg *infraconfig.Config) string {
	return filepath.Join(cfg.AppDir, cfg.DataDir)
}

func checkDatabaseVersion(dbPath string) (err error) {
	versionFileName := versionFilePath(dbPath)

	versionBytes, err := os.ReadFile(versionFileName)
	if err != nil {
		if os.IsNotExist(err) { // If version file doesn't exist, we assume that the database is new
			return createDatabaseVersionFile(dbPath, versionFileName)
		}
		return err
	}

	databaseVersion, err := strconv.Atoi(string(versionBytes))
	if err != nil {
		return err
	}

	if databaseVersion != currentDatabaseVersion {
		// TODO: Once there's more then one database version, it might make sense to add upgrade logic at this point
		return errors.Errorf("Invalid database version %d. Expected version: %d", databaseVersion, currentDatabaseVersion)
	}

	return nil
}

func versionFilePath(dbPath string) string {
	dbVersionFileName := path.Join(dbPath, "version")
	return dbVersionFileName
}

func createDatabaseVersionFile(dbPath string, versionFileName string) error {
	err := os.MkdirAll(dbPath, 0700)
	if err != nil {
		return err
	}

	versionFile, err := os.Create(versionFileName)
	if err != nil {
		return nil
	}
	defer versionFile.Close()

	versionString := strconv.Itoa(currentDatabaseVersion)
	_, err = versionFile.Write([]byte(versionString))
	return err
}

func removeDatabase(cfg *infraconfig.Config) error {
	dbPath := databasePath(cfg)
	return os.RemoveAll(dbPath)
}
