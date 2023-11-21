package blockchain

import "encoding/json"

type AmountRespose struct {
	Amount float32 `json:"amount"`
}

func (ar *AmountRespose) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Amount float32 `json:"amount"`
	}{
		Amount: ar.Amount,
	})
}
