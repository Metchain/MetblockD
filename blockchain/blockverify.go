package blockchain

import (
	"encoding/json"
	"fmt"
	pb "github.com/Metchain/MetblockD/protoserver/grpcserver/protowire"
	"github.com/Metchain/MetblockD/utils"
	"net"
)

func (b *MiniBlock) PreviousHash() [32]byte {
	return b.previousHash
}

func (b *MiniBlock) Nonce() uint64 {
	return b.nonce
}

func (b *MiniBlock) Transactions() []*Transaction {
	return b.transactions
}

func (bc *Blockchain) ValidChain(chain []*MiniBlock) bool {
	preBlock := chain[0]
	currentIndex := 1
	for currentIndex < len(chain) {
		b := chain[currentIndex]
		if b.previousHash != preBlock.Hash() {
			return false
		}

		/*if !bc.ValidProof(b.Nonce(), b.PreviousHash(), b.Transactions(), MINING_DIFFICULTY) {
			return false
		}*/

		preBlock = b
		currentIndex += 1
	}
	return true
}

func (mbc *Blockchain) ConvertBlockToDomainBlock() *pb.BlockInfo {
	_, lb := mbc.LastBlock()
	n := new(MBlock)
	n.UnmarshalJSON(lb)

	ProtoBlock := new(pb.BlockInfo)

	ProtoBlock.Blockhash = utils.ConvertDomainByteToProtoByte(n.CurrentHash)
	ProtoBlock.PreviousHash = utils.ConvertDomainByteToProtoByte(n.PreviousHash)
	ProtoBlock.Megablock = utils.ConvertDomainByteToProtoByte(n.Megablock)
	ProtoBlock.Metblock = utils.ConvertDomainByteToProtoByte(n.Metblock)
	ProtoBlock.Nonce = n.Nonce
	ProtoBlock.Blockheight = n.Height
	ProtoBlock.CurrentHash = utils.ConvertDomainByteToProtoByte(n.CurrentHash)
	ProtoBlock.Timestamp = n.Timestamp
	ProtoBlock.Bits = n.Bits
	for _, tx := range n.Transactions {
		ntx := new(pb.BlockTransactions)
		ntx.Sender = tx.senderBlockchainAddress
		ntx.Recipient = tx.recipientBlockchainAddress
		ntx.Txtype = utils.ConvertDomainInt8ToProtoInt64(tx.txtype)
		ntx.Value = utils.ConvertValToString(tx.value)
		ntx.Txhash = utils.ConvertDomainByteToProtoByte(tx.txhash)
		ntx.Timestamp = tx.timestamp
		ntx.Txstatus = utils.ConvertDomainInt8ToProtoInt64(tx.txstatus)
	}
	return ProtoBlock
}

func (bc *Blockchain) ConvertGroupToSecureDomainInfo(bks []*pb.BlockInfo) *pb.P2PBlockWithTrustedDataResponseMessage {

	return &pb.P2PBlockWithTrustedDataResponseMessage{
		BlockInfo:  bks,
		P2PMessage: nil,
	}

}
func (bc *Blockchain) ConvertBlockToDomainGroupBlock(b *pb.BlockInfo) []*pb.BlockInfo {
	block := make([]*pb.BlockInfo, 0)
	block = append(block, b)
	return block
}
func (bc *Blockchain) P2PBlockUngroup(bk *pb.P2PBlockWithTrustedDataRequestMessage) *pb.BlockInfo {
	for _, val := range bk.BlockInfo {
		return val
	}
	return nil
}

func (bc *Blockchain) MatchDomainBlockToP2PBlock(bk *pb.P2PBlockWithTrustedDataRequestMessage, domainIP net.Addr) error {
	ov := bc.ConvertBlockToDomainBlock()
	nbc := bc.P2PBlockUngroup(bk)

	//Match blockheight. If blockheight match but block hash doesn't match block the IP.
	if MatchBlockHeight(ov.Blockheight, nbc.Blockheight) == nil {
		// Verify Blockhash
		err := MatchDomainToP2PBytes(ov.Blockhash, nbc.Blockhash)
		if err != nil {
			return err
		}

		// Verify Previous Hash
		err = MatchDomainToP2PBytes(ov.PreviousHash, nbc.PreviousHash)
		if err != nil {
			return err
		}

		// Verify Mega Block

		err = MatchDomainToP2PBytes(ov.Metblock, nbc.Metblock)
		if err != nil {
			return err
		}

		// Verify Met Block

		err = MatchDomainToP2PBytes(ov.Megablock, nbc.Megablock)
		if err != nil {
			return err

		}
		// Verify Mega Block
		err = MatchDomainToP2PBytes(ov.CurrentHash, nbc.CurrentHash)
		if err != nil {
			return err
		}

	}
	return nil

}
func MatchDomainToP2PBytes(ov []byte, nv []byte) error {
	if len(ov) != len(nv) {
		return fmt.Errorf("Error recorded while processing the request. Consesus failure. Please make sure blockchain is synced.")
	}
	for key, val := range ov {
		if val != nv[key] {
			return fmt.Errorf("Error recorded while processing the request. Consesus failure. Please make sure blockchain is synced.")
		}
	}
	return nil
}

func MatchBlockHeight(i uint64, j uint64) error {
	if i != j {
		return fmt.Errorf("Error recorded while processing the request. Consesus failure. Please make sure blockchain is synced.")
	}
	return nil
}

type MBlock struct {
	Height       uint64
	Timestamp    int64
	Nonce        uint64
	PreviousHash [32]byte //As per the Hash size
	Megablock    [32]byte
	Metblock     [32]byte
	Transactions []*Transaction
	CurrentHash  [32]byte
	Bits         uint64
}

func (b *MBlock) UnmarshalJSON(data []byte) error {

	v := &struct {
		Height       uint64         `json:"height"`
		Timestamp    int64          `json:"timestamp"`
		Nonce        uint64         `json:"nonce"`
		PreviousHash [32]byte       `json:"previousHash"` //As per the Hash size
		Megablock    [32]byte       `json:"megablock"`
		Metblock     [32]byte       `json:"metblock"`
		Transactions []*Transaction `json:"transactions"`
		CurrentHash  [32]byte       `json:"currentHash"`
		Bits         uint64         `json:"bits"`
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
