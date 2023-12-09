package domain

import (
	"github.com/Metchain/Metblock/db/database"
)

func (mc *Metchain) Put(key *database.Key, value []byte) interface{} {
	//TODO implement me
	return mc.Db.Put(key, value)
}

func (mc *Metchain) Get(key *database.Key) ([]byte, interface{}) {
	//TODO implement me
	return mc.Db.Get(key)
}

func (mc *Metchain) Has(key *database.Key) (bool, interface{}) {
	//TODO implement me
	return mc.Db.Has(key)
}

func (mc *Metchain) Delete(key *database.Key) interface{} {
	//TODO implement me
	return mc.Db.Delete(key)
}

func (mc *Metchain) Close() interface{} {
	//TODO implement me
	return mc.Db.Close()
}

var genkey = database.MakeBucket([]byte("block")).Key([]byte("0"))
