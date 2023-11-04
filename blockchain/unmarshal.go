package blockchain

import (
	"encoding/json"
)

func (b *MiniBlock) UnmarshalJSON(data []byte) error {
	v := &struct {
		height       uint64         `json:"height"`
		timestamp    int64          `json:"timestamp"`
		nonce        uint64         `json:"nonce"`
		previousHash [32]byte       `json:"previousHash"` //As per the Hash size
		megablock    [32]byte       `json:"megablock"`
		metblock     [32]byte       `json:"metblock"`
		transactions []*Transaction `json:"transactions"`
		currentHash  [32]byte       `json:"currentHash"`
		bits         uint64         `json:"bits"`
	}{
		height:       b.height,
		timestamp:    b.timestamp,
		nonce:        b.nonce,
		previousHash: b.previousHash,
		megablock:    b.megablock,
		metblock:     b.metblock,
		transactions: b.transactions,
		currentHash:  b.currentHash,
		bits:         b.bits,
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	b.bits = v.bits
	return nil
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

func (t *JsonTransaction) UnmarshalJSON(data []byte) error {
	v := struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
		Txhash    string  `json:"txhash"`
		Timestamp int64   `json:"timestamp"`
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
	t.value = v.Value
	t.senderBlockchainAddress = v.Sender
	t.txhash = v.Txhash
	t.recipientBlockchainAddress = v.Recipient
	t.timestamp = v.Timestamp

	return nil
}
