package domain

import (
	"encoding/json"
	"log"
)

type domaintx struct {
	Timestamp    int64    `json:"timestamp"`
	Nonce        uint64   `json:"nonce"`
	PreviousHash [32]byte `json:"previousHash"`
	Merkleroot   []byte   `json:"merkleroot"`
	Height       uint64   `json:"height"`
	Transaction  []byte   `json:"transaction"`
	BlockType    string   `json:"blockType"`
}

func ReadTx(m []byte) domaintx {
	n := new(domaintx)
	v := &struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        uint64   `json:"nonce"`
		PreviousHash [32]byte `json:"previousHash"`
		Merkleroot   []byte   `json:"merkleroot"`
		Height       uint64   `json:"height"`
		Transaction  []byte   `json:"transaction"`
		BlockType    string   `json:"blockType"`
	}{
		Timestamp:    n.Timestamp,
		Nonce:        n.Nonce,
		Merkleroot:   n.Merkleroot,
		PreviousHash: n.PreviousHash,
		Height:       n.Height,
		Transaction:  n.Transaction,
		BlockType:    n.BlockType,
	}
	json.Unmarshal(m, v)
	n.Height = v.Height
	n.Timestamp = v.Timestamp
	n.Merkleroot = v.Merkleroot
	n.Nonce = v.Nonce
	n.PreviousHash = v.PreviousHash
	n.Transaction = v.Transaction
	n.BlockType = v.BlockType

	return *n
}

func ReadTxn(m []byte) *domaintx {
	n := new(domaintx)
	v := &struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        uint64   `json:"nonce"`
		PreviousHash [32]byte `json:"previousHash"`
		Merkleroot   []byte   `json:"merkleroot"`
		Height       uint64   `json:"height"`
		Transaction  []byte   `json:"transaction"`
		BlockType    string   `json:"blockType"`
	}{
		Timestamp:    n.Timestamp,
		Nonce:        n.Nonce,
		Merkleroot:   n.Merkleroot,
		PreviousHash: n.PreviousHash,
		Height:       n.Height,
		Transaction:  n.Transaction,
		BlockType:    n.BlockType,
	}
	json.Unmarshal(m, v)
	n.Height = v.Height
	n.Timestamp = v.Timestamp
	n.Merkleroot = v.Merkleroot
	n.Nonce = v.Nonce
	n.PreviousHash = v.PreviousHash
	n.Transaction = v.Transaction
	log.Printf("Check message: %s", m)
	n.BlockType = v.BlockType

	return n
}
