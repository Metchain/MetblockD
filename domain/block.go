package domain

import (
	"github.com/Metchain/Metblock/db/database"
	"github.com/Metchain/Metblock/mconfig"
	"github.com/btcsuite/goleveldb/leveldb"
)

var blockkey = database.MakeBucket([]byte("block"))

func (mc *Metchain) LastBlock() ([]byte, []byte) {

	db := mc.Db
	key := []byte{}
	value := []byte{}
	cursor, err := db.Cursor(blockkey)
	if err != nil {
		log.Error(err)
	}
	/**/

	for ok := cursor.Last(); ok; ok = cursor.Next() {
		dbkey, _ := cursor.Key()
		value, _ = cursor.Value()
		key = dbkey.Bytes()

	}

	//log.Println(fmt.Sprintf("%s", key))
	return key, value
}

func BlockByKey(height string) ([]byte, []byte) {
	d := mconfig.GetDBDir()

	//log.Println("Confirming Genesis Block")
	db, err := leveldb.OpenFile(d, nil)
	if err != nil {
		log.Infof("Error Opening DB: ", err)

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

func LastBlockRPC(db database.Database) ([]byte, []byte) {

	cursor, err := db.Cursor(blockkey)
	if err != nil {
		log.Info(err)
	}
	key := []byte{}
	value := []byte{}
	for ok := cursor.Last(); ok; ok = cursor.Next() {
		dbkey, _ := cursor.Key()
		value, _ = cursor.Value()
		key = dbkey.Bytes()

	}

	return key, value
}
