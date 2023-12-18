package blockchain

type DomainBlock struct {
	Height         uint64
	Timestamp      int64
	Nonce          uint64
	PreviousHash   [32]byte //As per the Hash size
	Megablock      [32]byte
	Metblock       [32]byte
	Transactions   []*Transaction
	CurrentHash    [32]byte
	UtxoCommitment string
	Bits           uint64
}
