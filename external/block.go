package external

import (
	"encoding/json"
	"fmt"
)

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
	Merkleroothash() *DomainHash
	DirectParents() BlockLevelParents
	MetBlock() *DomainHash
	MegaBlock() *DomainHash
	ChildBlocks() []BlockLevelChildern
	UTXOCommitment() *DomainHash
	TimeInMilliseconds() int64
	BlockLevel(maxBlockLevel int) int
	Parents() []BlockLevelParents
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

type TempBlock struct {
	Height       uint64
	Timestamp    int64
	Nonce        uint64
	PreviousHash [32]byte //As per the Hash size
	Megablock    [32]byte
	Metblock     [32]byte
	Transactions []*TempTransaction
	CurrentHash  [32]byte
	Bits         uint64
}

type TempTransaction struct {
	SenderBlockchainAddress    string
	RecipientBlockchainAddress string
	Txtype                     int8
	Value                      float32
	Txhash                     [32]byte
	Timestamp                  int64
	Txstatus                   int8
}

func (b *TempBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Height       uint64             `json:"height"`
		Timestamp    int64              `json:"timestamp"`
		Nonce        uint64             `json:"nonce"`
		PreviousHash [32]byte           `json:"previousHash"`
		Metblock     [32]byte           `json:"metblock"`
		Megablock    [32]byte           `json:"megablock"`
		CurrentHash  [32]byte           `json:"currentHash"`
		Transaction  []*TempTransaction `json:"transaction"`
		Bits         uint64             `json:"bits"`
	}{
		Height:       b.Height,
		Timestamp:    b.Timestamp,
		Nonce:        b.Nonce,
		PreviousHash: b.PreviousHash,
		Metblock:     b.Metblock,
		Megablock:    b.Megablock,
		CurrentHash:  b.CurrentHash,
		Transaction:  b.Transactions,
		Bits:         b.Bits,
	})

}

func (t *TempTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
		Txtype    int8    `json:"txtype"`
		Txhash    string  `json:"txhash"`
		Timestamp int64   `json:"timestamp"`
		Txstatus  int8    `json:"txstatus"`
	}{
		Sender:    t.SenderBlockchainAddress,
		Recipient: t.RecipientBlockchainAddress,
		Value:     t.Value,
		Txtype:    t.Txtype,

		Txhash:    fmt.Sprintf("%x", t.Txhash),
		Timestamp: t.Timestamp,
		Txstatus:  t.Txstatus,
	})
}
