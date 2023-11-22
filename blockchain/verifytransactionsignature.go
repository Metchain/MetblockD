package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/Metchain/Metblock/utils"
)

type TxVerify struct {
	pubkey    *ecdsa.PublicKey
	signature *utils.Signature
	tx        *Transaction
	hash      [32]byte
}

// for verification if the transaction is legit
func (bc *Blockchain) VerifyTransactionSignature(senderPublickey *ecdsa.PublicKey, s *utils.Signature, tx *Transaction) error {
	m, err := json.Marshal(tx)
	if err != nil {
		return err
	}

	h := sha256.Sum256(m)

	TxVerify := tx.CreateTxVerify(senderPublickey, s, h)
	err = TxVerify.VerifySignature()
	if err != nil {
		return err
	}

	return nil
}

func (Tx *TxVerify) VerifySignature() error {
	if ecdsa.Verify(Tx.pubkey, Tx.hash[:], Tx.signature.R, Tx.signature.S) {
		return nil
	}
	return fmt.Errorf("Error verifying signature")
}

func (Tx *Transaction) CreateTxVerify(pubkey *ecdsa.PublicKey, s *utils.Signature, hash [32]byte) *TxVerify {
	return &TxVerify{
		pubkey:    pubkey,
		signature: s,
		tx:        Tx,
		hash:      hash,
	}
}
