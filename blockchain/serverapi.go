package blockchain

import (
	"crypto/ecdsa"
	"github.com/Metchain/Metblock/utils"
)

// created for blockchain node
func (bc *Blockchain) CreateTransaction(sender string, recipient string, value float32,
	senderPublicKey *ecdsa.PublicKey, s *utils.Signature) bool {

	isTransacted := bc.AddTransaction(sender, recipient, value, senderPublicKey, s)

	return isTransacted
}
