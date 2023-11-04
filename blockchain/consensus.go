package blockchain

import (
	"encoding/json"
	"github.com/Metchain/Metblock/blockchain/consensus"
	"github.com/Metchain/Metblock/mconfig"
	"log"
)

type GenesisConsensus struct {
	Timestamp    int64
	Nonce        int
	PreviousHash [32]byte //As per the Hash size
	Merkleroot   []byte
	Height       int64
	Transaction  mconfig.Txtransaction
}

func GenesisGenrate() *GenesisConsensus {
	g := *mconfig.NewGen()
	Gc := new(GenesisConsensus)
	Gc.Nonce = g.Nonce
	Gc.Timestamp = g.Timestamp
	Gc.Merkleroot = g.Merkleroot
	Gc.PreviousHash = g.PreviousHash
	Gc.Height = g.Height
	Gc.Transaction = g.Transaction

	return Gc
}

func (gc *GenesisConsensus) VerifyGenesis() (string, bool) {

	msg, err := consensus.VerfiyNonce(gc.Nonce)
	if err {
		log.Println(msg)
	} else {
		log.Println(msg)
	}

	msg, err = consensus.VerifyTimestamp(int(gc.Timestamp))
	if err {
		log.Println(msg)
	} else {
		log.Println(msg)
	}
	msg, err = consensus.VerifyMessage()
	if err {
		log.Println(msg)
	} else {
		log.Println(msg)
	}
	msg, err = consensus.VerifyMerkleRoot()
	if err {
		log.Println(msg)
	} else {
		log.Println(msg)
	}
	msg, err = consensus.VerifyPreviousHash()
	if err {
		log.Println(msg)
	} else {
		log.Println(msg)
	}

	msg, err = consensus.VerifyTransaction()
	if err {
		log.Println(msg)
	} else {
		log.Println(msg)
	}
	if err {
		return msg, err
	} else {
		return "Consensus Verification Complete", err
	}
}

func (gc *GenesisConsensus) GensisCompile() []byte {
	m, _ := json.Marshal(gc)
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

func (g *GenesisConsensus) GenesisUncompile(m []byte) {
	v := &struct {
		Timestamp    int64                 `json:"timestamp"`
		Nonce        int                   `json:"nonce"`
		PreviousHash [32]byte              `json:"previousHash"`
		Merkleroot   []byte                `json:"merkleroot"`
		Height       int64                 `json:"height"`
		Transaction  mconfig.Txtransaction `json:"transaction"`
	}{
		Timestamp:    g.Timestamp,
		Nonce:        g.Nonce,
		Merkleroot:   g.Merkleroot,
		PreviousHash: g.PreviousHash,
		Height:       g.Height,
		Transaction:  g.Transaction,
	}
	json.Unmarshal(m, v)
	log.Printf("%v", (v.Timestamp))
}
