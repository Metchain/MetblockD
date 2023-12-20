package blockchain

import (
	"encoding/json"
	"github.com/Metchain/MetblockD/blockchain/consensus"
	"github.com/Metchain/MetblockD/mconfig"
)

type GenesisConsensus struct {
	Timestamp    int64
	Nonce        int
	PreviousHash [32]byte //As per the Hash size
	Merkleroot   []byte
	Height       int64
	Transaction  *mconfig.Txtransaction
}

func GenesisGenrate() *GenesisConsensus {
	g := *mconfig.NewGen()
	return &GenesisConsensus{
		Nonce:        g.Nonce,
		Timestamp:    g.Timestamp,
		Merkleroot:   g.Merkleroot,
		PreviousHash: g.PreviousHash,
		Height:       g.Height,
		Transaction:  g.Transaction,
	}
}

func (gc *GenesisConsensus) VerifyNonce() {
	msg, err := consensus.VerfiyNonce(gc.Nonce)
	if err {
		log.Criticalf("%s : %s", msg, err)
	}

}

func (gc *GenesisConsensus) VerifyTimestamp() {
	msg, err := consensus.VerifyTimestamp(int(gc.Timestamp))
	if err {
		log.Criticalf("%s : %s", msg, err)
	}

}

func (gc *GenesisConsensus) VerifyMessage() {
	msg, err := consensus.VerifyMessage()
	if err {
		log.Criticalf("%s : %s", msg, err)
	}

}

func (gc *GenesisConsensus) VerifyMerkleRoot() {
	msg, err := consensus.VerifyMerkleRoot()
	if err {
		log.Criticalf("%s : %s", msg, err)
	}
}

func (gc *GenesisConsensus) VerifyPreviousHash() {
	msg, err := consensus.VerifyPreviousHash()
	if err {
		log.Criticalf("%s : %s", msg, err)
	}
}

func (gc *GenesisConsensus) VerifyTransaction() {
	msg, err := consensus.VerifyTransaction()
	if err {
		log.Criticalf("%s : %s", msg, err)
	}

}

func (gc *GenesisConsensus) VerifyGenesis() (string, bool) {

	// Verify the Gensis Nonce
	gc.VerifyNonce()

	// Verify the Gensis Timestamp
	gc.VerifyTimestamp()

	// Verify the Gensis Message
	gc.VerifyMessage()

	// Verify the Gensis MerkleRoot
	gc.VerifyMerkleRoot()

	// Verify the Gensis PreviousHash
	gc.VerifyPreviousHash()

	// Verify the Gensis Transaction
	gc.VerifyTransaction()
	return "Gensis verification Complete.", false
}

func (gc *GenesisConsensus) GensisCompile() []byte {
	m, _ := gc.MarshalJSON()
	return m
}

type Txtransaction struct {
	Txhex      []byte
	Txcoinbase []byte
}

func (tx *Txtransaction) MarshalJSON() ([]byte, error) {

	n, _ := json.Marshal(struct {
		Txhex      []byte `json:"txhex"`
		Txcoinbase []byte `json:"txcoinbase"`
	}{
		Txhex:      tx.Txhex,
		Txcoinbase: tx.Txcoinbase,
	})

	return n, nil
}

func (gc *GenesisConsensus) MarshalJSON() ([]byte, error) {
	tn, _ := json.Marshal(struct {
		Txhex      []byte `json:"txhex"`
		Txcoinbase []byte `json:"txcoinbase"`
	}{
		Txhex:      gc.Transaction.Txhex,
		Txcoinbase: gc.Transaction.Txhex,
	})

	n, _ := json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previousHash"`
		Merkleroot   []byte   `json:"merkleroot"`
		Height       int64    `json:"height"`
		Transaction  []byte   `json:"transaction"`
	}{
		Timestamp:    gc.Timestamp,
		Nonce:        gc.Nonce,
		Merkleroot:   gc.Merkleroot,
		PreviousHash: gc.PreviousHash,
		Height:       gc.Height,
		Transaction:  tn,
	})

	return n, nil
}
