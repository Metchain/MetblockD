package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"github.com/Metchain/Metblock/utils"
	"log"
)

// for verification if the transaction is legit
func (bc *Blockchain) VerifyTransactionSignature(senderPublickey *ecdsa.PublicKey, s *utils.Signature, t *Transaction) bool {
	m, err := json.Marshal(t)
	if err != nil {
		log.Println("Here is the erro: ", err)
	}
	log.Println(string(m))

	//
	h := sha256.Sum256(m)

	//os.Exit(555)
	return ecdsa.Verify(senderPublickey, h[:], s.R, s.S)
	return ecdsa.Verify(senderPublickey, h[:], s.R, s.S)
}
