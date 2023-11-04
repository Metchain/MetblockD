package domain

import (
	"fmt"
	"github.com/Metchain/Metblock/mconfig"
	"github.com/btcsuite/goleveldb/leveldb"
	"log"
	"strings"
)

type Metchain struct {
	Dbcon       *leveldb.DB
	Blockheight uint64
	Cdiff       uint64
}

func GenesisConsensusDBCreate(gc []byte) bool {
	d := mconfig.GetDBDir()

	log.Println("Generating Genesis Block")
	db, err := leveldb.OpenFile(d, nil)
	genkey := "bk-" + strings.Repeat("0", 64)
	err = db.Put([]byte(genkey), gc, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	db.Close()
	return true
}

func VerifyGenesisConsensusDB(gc []byte) (bool, *Metchain) {
	d := mconfig.GetDBDir()
	log.Println(d)

	genkey := "bk-" + strings.Repeat("0", 64)
	log.Println("Confirming Genesis Block")
	metch := new(Metchain)
	db, err := leveldb.OpenFile(d, nil)

	if err != nil {
		log.Println("Error Opening DB: ", err)

	}
	metch.Dbcon = db
	data, err := db.Get([]byte(genkey), nil)
	if err != nil {
		log.Println("Error Reading DB: ", err)
	}
	verr := 0
	gn := ReadTx(gc)
	b := ReadTx(data)

	if gn.Height == b.Height {
		log.Println("DB: Block Height Verified")
	} else {
		verr = 1
	}
	if gn.Timestamp == b.Timestamp {
		log.Println("DB: Block Timestamp Verified")
	} else {
		verr = 1
	}
	if gn.Nonce == b.Nonce {
		log.Println("DB: Block Nonce Verified")
	} else {
		verr = 1
	}

	if gn.PreviousHash == b.PreviousHash {
		log.Println("DB: Block PreviousHash Verified")
	} else {
		verr = 1
	}

	if fmt.Sprintf("%x", gn.Merkleroot) == fmt.Sprintf("%x", b.Merkleroot) {
		log.Println("DB: Block Merkleroot Verified")
	} else {
		verr = 1
	}

	if fmt.Sprintf("%x", gn.Transaction) == fmt.Sprintf("%x", b.Transaction) {
		log.Println("DB: Block Transaction Verified")
	} else {
		verr = 1
	}
	if verr == 1 {
		DBReset()
		GenesisConsensusDBCreate(gc)
		verr = 0
	}
	if verr == 0 {
		return true, metch
	} else {

		log.Println("ConsensusError: There is an unexpected error. Make sure you have the latest version")
		defer metch.Dbcon.Close()
		//Add exit code with Manual
		//os.Exit(0)
		return false, nil
	}

}
