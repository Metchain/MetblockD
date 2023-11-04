package blockchain

import "encoding/json"

type AmountRespose struct {
	Amount float32 `json:"amount"`
}

func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0
	/*for _, b := range bc.chain {
		/*for _, t := range b.Transactions {
			value := t.value

			if blockchainAddress == t.recipientBlockchainAddress {
				totalAmount += value
			}
			if blockchainAddress == t.senderBlockchainAddress {
				totalAmount -= value
			}
		}
	}*/
	return totalAmount
}

func (ar *AmountRespose) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Amount float32 `json:"amount"`
	}{
		Amount: ar.Amount,
	})
}
