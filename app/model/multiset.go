package model

import "github.com/Metchain/Metblock/external"

// Multiset represents a secp256k1 multiset
type Multiset interface {
	Add(data []byte)
	Remove(data []byte)
	Hash() *external.DomainHash
	Serialize() []byte
	Clone() Multiset
}

// MultisetStore represents a store of Multisets
type MultisetStore interface {
	Stage(stagingArea *StagingArea, blockHash *external.DomainHash, multiset Multiset)
	IsStaged(stagingArea *StagingArea) bool
	Get(dbContext DBReader, stagingArea *StagingArea, blockHash *external.DomainHash) (Multiset, error)
	Delete(stagingArea *StagingArea, blockHash *external.DomainHash)
}
