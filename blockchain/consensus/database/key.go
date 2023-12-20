package database

import (
	"github.com/Metchain/MetblockD/app/model"
	"github.com/Metchain/MetblockD/db/database"
)

type dbKey struct {
	key *database.Key
}

func (d dbKey) Bucket() model.DBBucket {

	return d.Bucket()
}

func (d dbKey) Bytes() []byte {
	return d.key.Bytes()
}

func (d dbKey) Suffix() []byte {
	return d.key.Suffix()
}

func newDBKey(key *database.Key) model.DBKey {
	return dbKey{key: key}
}
