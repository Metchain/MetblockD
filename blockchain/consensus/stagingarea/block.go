package stagingarea

import (
	"encoding/json"
	"github.com/Metchain/Metblock/db/database"
	"github.com/Metchain/Metblock/external"
	"log"
)

var blockkey = database.MakeBucket([]byte("block"))

type StagingArea struct {
	databaseContext database.Database
	stagingKey      stagingKey
	stagingBlock    stagingBlock
}
type stagingBlock []byte
type stagingKey []byte
type BlockHeader struct {
	Version            uint16
	Blockhash          *external.DomainHash
	Blockheight        int64
	Previoushash       *external.DomainHash
	Merkleroot         *external.DomainHash
	Metblock           *external.DomainHash
	Megablock          *external.DomainHash
	Childblocks        []*external.BlockLevelChildern
	Parents            []*external.BlockLevelParents
	TimeInMilliseconds int64
	Btype              int
	Bits               uint64
	Nonce              uint64
	BlockLevel         int
	VerificationPoint  *external.DomainHash
}

func NewStagingArea(db database.Database) *StagingArea {
	return &StagingArea{
		databaseContext: db,
	}
}

func (s *StagingArea) StagingBlock() *BlockHeader {

	err := s.getStagingBlock()

	if err != nil {
		log.Fatalf("%v", err)

	}
	var bt int
	b := s.convertDomainBlockToRPCBlock()
	s.getBlockHash(b.CurrentHash)
	if (b.CurrentHash) == (b.Metblock) {
		bt = 0x04
	}
	if (b.CurrentHash) == (b.Megablock) {
		bt = 0x03
	} else {
		bt = 0x02
	}

	return &BlockHeader{

		Blockhash:          external.NewDomainHashFromByteArray(&b.CurrentHash),
		Blockheight:        int64(b.Height),
		Previoushash:       external.NewDomainHashFromByteArray(&b.PreviousHash),
		Merkleroot:         external.NewDomainHashFromByteArray(&b.PreviousHash),
		Metblock:           external.NewDomainHashFromByteArray(&b.Metblock),
		Megablock:          external.NewDomainHashFromByteArray(&b.Megablock),
		Btype:              bt,
		TimeInMilliseconds: b.Timestamp,
		Bits:               b.Bits,
		Nonce:              b.Nonce,
		BlockLevel:         1,
		VerificationPoint:  external.NewDomainHashFromByteArray(&b.PreviousHash),
	}
}

func (s *StagingArea) getBlockHash(hash [32]byte) *external.DomainHash {
	d := external.NewDomainHashFromByteArray(&hash)

	return d
}

func (s *StagingArea) convertDomainBlockToRPCBlock() *mBlock {
	block := &mBlock{}

	block.UnmarshalJSON(s.stagingBlock[:])
	return block
}

func (s *StagingArea) getStagingBlock() error {

	key := []byte{}
	value := []byte{}
	cursor, err := s.databaseContext.Cursor(blockkey)

	if err != nil {
		panic(err)
	}
	for ok := cursor.Last(); ok; ok = cursor.Next() {
		dbkey, _ := cursor.Key()
		value, _ = cursor.Value()
		key = dbkey.Bytes()

	}
	s.stagingKey = key
	s.stagingBlock = value

	//log.Println(fmt.Sprintf("%s", key))
	return nil

}

type mBlock struct {
	Height       uint64
	Timestamp    int64
	Nonce        uint64
	PreviousHash [32]byte //As per the Hash size
	Megablock    [32]byte
	Metblock     [32]byte
	Transactions []*mTransaction
	CurrentHash  [32]byte
	Bits         uint64
}

func (b *mBlock) UnmarshalJSON(data []byte) error {

	v := &struct {
		Height       uint64          `json:"height"`
		Timestamp    int64           `json:"timestamp"`
		Nonce        uint64          `json:"nonce"`
		PreviousHash [32]byte        `json:"previousHash"` //As per the Hash size
		Megablock    [32]byte        `json:"megablock"`
		Metblock     [32]byte        `json:"metblock"`
		Transactions []*mTransaction `json:"transactions"`
		CurrentHash  [32]byte        `json:"currentHash"`
		Bits         uint64          `json:"bits"`
	}{
		Height:       b.Height,
		Timestamp:    b.Timestamp,
		Nonce:        b.Height,
		PreviousHash: b.PreviousHash,
		Megablock:    b.Megablock,
		Metblock:     b.Metblock,
		Transactions: b.Transactions,
		CurrentHash:  b.CurrentHash,
		Bits:         b.Bits,
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	b.Height = v.Height
	b.Timestamp = v.Timestamp
	b.Nonce = v.Height
	b.PreviousHash = v.PreviousHash
	b.Megablock = v.Megablock
	b.Metblock = v.Metblock
	b.Transactions = v.Transactions
	b.CurrentHash = v.CurrentHash
	b.Bits = v.Bits

	return nil
}

type mTransaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	txtype                     int8
	value                      float32
	txhash                     [32]byte
	timestamp                  int64
	txstatus                   int8
}

func (t *mTransaction) UnmarshalJSON(data []byte) error {
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
