package main

import "encoding/json"

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	txtype                     int8
	value                      float32
	txhash                     [32]byte
	timestamp                  int64
	txstatus                   int8
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	v := struct {
		Sender    string   `json:"sender_blockchain_address"`
		Recipient string   `json:"recipient_blockchain_address"`
		Value     float32  `json:"value"`
		Txhash    [32]byte `json:"txhash"`
		Timestamp int64    `json:"timestamp"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
		Txhash:    t.txhash,
		Timestamp: t.timestamp,
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}
