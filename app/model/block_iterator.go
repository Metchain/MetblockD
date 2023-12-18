package model

import "github.com/Metchain/Metblock/external"

type BlockIterator interface {
	First() bool
	Next() bool
	Get() (*external.DomainHash, error)
	Close() error
}
