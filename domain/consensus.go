package domain

import (
	"fmt"
	"github.com/Metchain/Metblock/db/database"
	"github.com/Metchain/Metblock/mconfig/infraconfig"
)

type Metchain struct {
	Db          database.Database
	Blockheight uint64
	Cdiff       uint64
}

func GenesisConsensusDBCreate(gc []byte, db database.Database) bool {

	log.Infof("Generating Genesis Block")

	err := db.Put(genkey, gc)
	if err != nil {
		log.Infof("Error: ", err)
	}

	return true
}

func VerifyGenesisConsensusDB(gc []byte, db database.Database, cfg *infraconfig.Config) *Metchain {

	log.Infof("Confirming Genesis Block")
	metch := new(Metchain)

	metch.Db = db

	data, err := db.Get(genkey)
	verr := 0

	if err != nil {
		log.Infof("Error Reading DB: ", err)

	}

	gn := ReadTx(gc)
	b := ReadTx(data)

	if gn.Height == b.Height {
		log.Infof("DB: Block Height Verified")
	} else {
		verr = 1
	}
	if gn.Timestamp == b.Timestamp {
		log.Infof("DB: Block Timestamp Verified")
	} else {
		verr = 1
	}
	if gn.Nonce == b.Nonce {
		log.Infof("DB: Block Nonce Verified")
	} else {
		verr = 1
	}

	if gn.PreviousHash == b.PreviousHash {
		log.Infof("DB: Block PreviousHash Verified")
	} else {
		verr = 1
	}

	if fmt.Sprintf("%x", gn.Merkleroot) == fmt.Sprintf("%x", b.Merkleroot) {
		log.Infof("DB: Block Merkleroot Verified")
	} else {
		verr = 1
	}

	if fmt.Sprintf("%x", gn.Transaction) == fmt.Sprintf("%x", b.Transaction) {
		log.Infof("DB: Block Transaction Verified")
	} else {
		verr = 1
	}
	if verr == 1 {
		DBReset(cfg)
		GenesisConsensusDBCreate(gc, db)
		verr = 0
	}
	if verr == 0 {
		return metch
	} else {

		log.Infof("ConsensusError: There is an unexpected error. Make sure you have the latest version")
		defer metch.Db.Close()
		//Add exit code with Manual
		//os.Exit(0)
		return nil
	}

}
