package external

// DomainBlock represents a Metchain block
type DomainBlock struct {
	Header       BlockHeader
	Transactions []*DomainTransaction
}

// Clone returns a clone of DomainBlock
func (block *DomainBlock) Clone() *DomainBlock {
	transactionClone := make([]*DomainTransaction, len(block.Transactions))
	/*for i, tx := range block.Transactions {
		transactionClone[i] = tx.Clone()
	}*/

	return &DomainBlock{
		Header:       block.Header,
		Transactions: transactionClone,
	}
}

type BlockHeader interface {
	BaseBlockHeader
	ToMutable() MutableBlockHeader
}

// BaseBlockHeader represents the header part of a MET block
type BaseBlockHeader interface {
	Version() uint16
	Blockheight() uint64
	BlockHash() *DomainHash
	Previoushash() *DomainHash
	Merkleroot() *DomainHash
	DirectParents() []*BlockLevelParents
	MetBlock() *DomainHash
	MegaBlock() *DomainHash
	ChildBlocks() []*BlockLevelChildern
	UTXOCommitment() *DomainHash
	TimeInMilliseconds() int64
	BlockLevel(maxBlockLevel int) int
	ParentByteToString() []string
	Bits() uint64
	Nonce() uint64
	UtxoCommitment() []byte
	Btype() int
}

type MutableBlockHeader interface {
	BaseBlockHeader
	ToImmutable() BlockHeader
	SetNonce(nonce uint64)
	SetTimeInMilliseconds(timeInMilliseconds int64)
	SetHashMerkleRoot(hashMerkleRoot *DomainHash)
}

// DomainTransaction represents a Metchain transaction
type DomainTransaction struct {
	Version uint16
	Inputs  []*DomainTransactionInput

	LockTime uint64

	Payload []byte

	Fee uint64

	ID *DomainTransactionID
}

// DomainTransactionInput represents a Metchain transaction input
type DomainTransactionInput struct {
	PreviousOutpoint DomainOutpoint
	SignatureScript  []byte
	Sequence         uint64
	SigOpCount       byte

	UTXOEntry UTXOEntry
}

// represents a Metchain ScriptPublicKey
type ScriptPublicKey struct {
	Script  []byte
	Version []byte
}

// DomainTransactionID represents the ID of a Metchain transaction
type DomainTransactionID DomainHash

// BlockAdded is an event raised by consensus when a block was added to the dag
type BlockAdded struct {
	Block *DomainBlock
}

// BlockAdded is an event raised by consensus when a block was added to the dag
type BlockSent struct {
	Block *DomainBlock
}

func (*BlockAdded) isConsensusEvent() {}

func (*BlockSent) isConsensusEvent() {}
