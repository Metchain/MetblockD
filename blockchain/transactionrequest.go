package blockchain

type TransactionRequest struct {
	SenderBlockchainAddress    *string  `json:"senderBlockchainAddress"`
	RecipientBlockchainAddress *string  `json:"recipientBlockchainAddress"`
	SenderPublicKey            *string  `json:"senderPublicKey"`
	Value                      *float32 `json:"value"`
	Signature                  *string  `json:"signature"`
}

/*func (tr *TransactionRequest) Validate() bool {
	if tr.SenderPublicKey == nil || tr.SenderBlockchainAddress == nil || tr.Value == nil || tr.Signature == nil {
		return false
	}
	return true
}*/
