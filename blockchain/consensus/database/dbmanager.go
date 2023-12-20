package database

import (
	"github.com/Metchain/MetblockD/db/database"
)

type dbManager struct {
	db database.Database
}

func (d dbManager) Put(key *database.Key, value []byte) error {
	//TODO implement me
	return d.db.Put(key, value)
}

func (d dbManager) Get(key *database.Key) ([]byte, error) {
	//TODO implement me
	return d.db.Get(key)
}

func (d dbManager) Has(key *database.Key) (bool, error) {
	//TODO implement me
	return d.db.Has(key)
}

func (d dbManager) Delete(key *database.Key) error {
	//TODO implement me

	return d.db.Delete(key)
}

func (d dbManager) Cursor(bucket *database.Bucket) (database.Cursor, error) {
	//TODO implement me
	return d.db.Cursor(bucket)
}

func (d dbManager) Begin() (database.Transaction, error) {
	//TODO implement me
	return d.db.Begin()
}

func (d dbManager) Compact() error {
	//TODO implement me

	return d.db.Compact()
}

func (d dbManager) Close() error {
	//TODO implement me
	return d.db.Close()
}

// New returns wraps the given database as an instance of model.DBManager
func New(db database.Database) database.Database {
	return &dbManager{db: db}
}
