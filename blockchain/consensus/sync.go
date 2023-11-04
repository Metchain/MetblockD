package consensus

import (
	"github.com/Metchain/Metblock/domain"
	"github.com/Metchain/Metblock/mconfig"
	"log"
	"os"
)

type status bool

func Sync(gc []byte) *domain.Metchain {
	metch := checkDB(gc)

	return metch
}
func checkDB(gc []byte) *domain.Metchain {
	s := checkDBDir()
	if !s {
		db := domain.GenesisConsensusDBCreate(gc)
		if !db {
			log.Println("Genesis Block Created")

		}
	}
	db, metch := domain.VerifyGenesisConsensusDB(gc)
	if !db {
		log.Println("Gensis Block Verified. Starting Node")
	}
	return metch

}

func checkDBDir() status {
	dir := mconfig.GetDatadir()

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println(dir, "does not exist")
		return false
	} else {
		log.Println("The provided directory named", dir, "exists")
		return true
	}
}
