package wallets

type TransactionRequest struct {
	SenderPrivateKey           *string `json:"senderPrivateKey"`
	SenderBlockchainAddress    *string `json:"senderBlockchainAddress"`
	RecipientBlockchainAddress *string `json:"recipientBlockchainAddress"`
	SenderPublicKey            *string `json:"senderPublicKey"`
	Value                      *string `json:"value"`
}
