package external

type ReadOnlyUTXOSetIterator interface {
	First() bool
	Next() bool
	Get() (outpoint *DomainOutpoint, utxoEntry UTXOEntry, err error)
	Close() error
}
