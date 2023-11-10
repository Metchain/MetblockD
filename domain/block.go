package domain

import (
	"github.com/Metchain/Metblock/mconfig"
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/util"
	"log"
)

func (mc *Metchain) LastBlock() ([]byte, []byte) {

	db := mc.Dbcon
	key := []byte{}
	value := []byte{}
	iter := db.NewIterator(util.BytesPrefix([]byte("bk-")), nil)
	/**/

	ok := iter.Last()
	if ok {
		key = iter.Key()
		value = iter.Value()

	}
	ok = iter.Next()
	if ok {
		key = iter.Key()
		value = iter.Value()

	}

	iter.Release()
	//log.Println(fmt.Sprintf("%s", key))
	return key, value
}

func BlockByKey(height string) ([]byte, []byte) {
	d := mconfig.GetDBDir()

	//log.Println("Confirming Genesis Block")
	db, err := leveldb.OpenFile(d, nil)
	if err != nil {
		log.Println("Error Opening DB: ", err)

	}
	key := []byte{}
	value := []byte{}
	iter := db.NewIterator(nil, nil)
	iter.Seek([]byte(height))
	key = iter.Key()
	value = iter.Value()
	//log.Printf("key: %s | value: %s\n", key, value)
	defer db.Close()
	return key, value
}

func LastBlockRPC(db *leveldb.DB) ([]byte, []byte) {

	key := []byte{}
	value := []byte{}
	iter := db.NewIterator(util.BytesPrefix([]byte("bk-")), nil)
	/**/

	ok := iter.Last()
	if ok {
		key = iter.Key()
		value = iter.Value()

	}
	iter.Release()

	return key, value
}
