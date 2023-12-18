package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Metchain/Metblock/db/database"
	"github.com/Metchain/Metblock/domain"
	"github.com/Metchain/Metblock/heavyhash"
	"github.com/Metchain/Metblock/mconfig"
	pb "github.com/Metchain/Metblock/protoserver/grpcserver/protowire"
	"google.golang.org/protobuf/proto"

	"os"
	"strconv"
	"strings"
	"time"
)

type MiniBlock struct {
	height       uint64
	timestamp    int64
	nonce        uint64
	previousHash [32]byte //As per the Hash size
	megablock    [32]byte
	metblock     [32]byte
	transactions []*Transaction
	currentHash  [32]byte
	bits         uint64
}

type JsonMiniBlock struct {
	height       uint64
	timestamp    int64
	nonce        uint64
	previousHash [32]byte //As per the Hash size
	megablock    [32]byte
	metblock     [32]byte
	transactions []*JsonTransaction
	currentHash  [32]byte
	bits         uint64
}

func GetBlockDB(height int) ([]byte, []byte) {
	kln := 64 - len(strconv.Itoa(height))
	key := "bk-" + strings.Repeat("0", kln) + strconv.Itoa(height)

	return domain.BlockByKey(key)

}

var MBlockkey = database.MakeBucket([]byte("block"))
var Utxokey = database.MakeBucket([]byte("utxo_tx"))

func (b *MiniBlock) BlockToDB(db database.Database) error {

	key := strconv.Itoa(int(b.height))
	b.MarshalJSON()
	m, _ := json.Marshal(b)
	log.Infof("Block accepted. Current Block Height:", b.height)
	batch, _ := db.Begin()
	batch.Put(MBlockkey.Key([]byte(key)), m)

	for _, tx := range b.transactions {

		//tx := transaction.UnMarshalFor()
		//log.Println(transaction.senderBlockchainAddress)

		utxo := new(pb.UTXOResponse)
		utxo.Txhash = fmt.Sprintf("%x", tx.txhash)
		utxo.Timestamp = fmt.Sprintf("%v", tx.timestamp)
		utxo.Amount = fmt.Sprintf("%f", tx.value)
		utxo.Fromwallet = tx.senderBlockchainAddress
		utxo.Towallet = tx.recipientBlockchainAddress
		// Update the block info
		utxo.Blockid = fmt.Sprintf("%v", b.height)
		utxo.Txtype = fmt.Sprintf("%v", tx.txtype)
		utxo.Blockhash = fmt.Sprintf("%x", b.currentHash)

		jsonUtxo, err := proto.Marshal(utxo)
		if err != nil {
			log.Criticalf("This shouldn't happen. Please check version")
			os.Exit(36)
		}
		key := fmt.Sprintf("%x", tx.txhash)

		batch.Put(Utxokey.Key([]byte(key)), jsonUtxo)

		if tx.txtype == 3 {
			nftpb := new(pb.NFTResponse)
			nftkey := fmt.Sprintf("%v", tx.value)
			nftsd, err := db.Get(nftLayerkey.Key([]byte(nftkey)))
			if err != nil {
				log.Infof("Error Locating NFT :", nftkey, nftsd)
			}

			nftpb.NFTWallet = tx.recipientBlockchainAddress
			nftpb.NFTSender = tx.senderBlockchainAddress
			nftpb.NFTid = fmt.Sprintf("%v", tx.value)
			nftpb.Txhash = fmt.Sprintf("%x", tx.txhash)
			nftpb.Blockid = fmt.Sprintf("%v", b.height)
			nftpb.Blockhash = fmt.Sprintf("%x", b.currentHash)
			jsonnft, err := proto.Marshal(nftpb)
			if err != nil {
				log.Criticalf("This should happen at all. :", err)
				os.Exit(123456789)
			}
			batch.Put(nftLayerkey.Key([]byte(nftkey)), jsonnft)
		}

	}
	wallets := make(map[string]*pb.UTXOWalletBalanceRespose, 0)
	wallets = UpdateWalletList(b, db, wallets)
	for tempkey, val := range wallets {
		jsonWalletSender, err := proto.Marshal(val)
		if err != nil {
			log.Infof("Error. This shouldn't happen")
			os.Exit(8888)
		}
		batch.Put(utxo_walletkey.Key([]byte(tempkey)), jsonWalletSender)
		log.Infof("Added Wallet: " + val.WalletAddress + "\n Balance: " + val.Amount + "\n\n")
	}
	err := batch.Commit()
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	return nil
}

func Hash(previousHash []byte, timestamp int64, nonce uint64) [32]byte {
	return [32]byte(heavyhash.HeavyHash(previousHash, timestamp, nonce))
}

func (b *MiniBlock) Hash() [32]byte {

	//previousHash := []byte(fmt.Sprintf("%x", LastMiniBlock().previousHash))
	previousHash := []byte(fmt.Sprintf("%x", [32]byte{}))
	return [32]byte(heavyhash.HeavyHash(previousHash, b.timestamp, uint64(b.nonce)))
}

func (b *MiniBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Height       uint64         `json:"height"`
		Timestamp    int64          `json:"timestamp"`
		Nonce        uint64         `json:"nonce"`
		PreviousHash [32]byte       `json:"previousHash"`
		Metblock     [32]byte       `json:"metblock"`
		Megablock    [32]byte       `json:"megablock"`
		CurrentHash  [32]byte       `json:"currentHash"`
		Transaction  []*Transaction `json:"transaction"`
		Bits         uint64         `json:"bits"`
	}{
		Height:       b.height,
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Metblock:     b.metblock,
		Megablock:    b.megablock,
		CurrentHash:  b.currentHash,
		Transaction:  b.transactions,
		Bits:         b.bits,
	})

}

func CreateMiniBlock(b *pb.RpcBlock, db database.Database, bc *Blockchain) ([32]byte, error, float64) {
	Temptimestamp := time.Now().UnixMilli()
	Tempoldntime := bc.LastRPCBlock.Timestamp + 6000
	if bc.LastRPCBlock.Timestamp > Temptimestamp || Tempoldntime >= Temptimestamp {
		//log.Printf("This is trigered")
		return [32]byte{}, errors.New("Block is too early"), 0
	}

	bc.Mux.Lock()
	defer bc.Mux.Unlock()

	nb := new(MiniBlock)

	var oldtimestamp int64
	nb.bits = uint64(b.Header.Bits)
	//log.Println(bc.MiniBlock)

	nb.nonce = b.Header.Nonce
	nb.height = bc.LastRPCBlock.Height + 1
	nb.previousHash = bc.LastRPCBlock.CurrentHash
	oldtimestamp = bc.LastRPCBlock.Timestamp

	nb.timestamp = time.Now().UnixMilli()

	oldntime := oldtimestamp + 6000

	if oldtimestamp > nb.timestamp || oldntime >= nb.timestamp {
		//log.Printf("This is trigered")
		nb.bits = nb.bits + 1
		return [32]byte{}, errors.New("Block is too early"), 0
	} else if nb.timestamp-oldntime > 4000 && nb.bits > 569658475 {
		nb.bits = nb.bits - 1
	} else {
		nb.bits = nb.bits + 1
	}

	ntx := bc.Txpool
	nfttx := bc.NFTPool

	nftre := []*Transaction{}
	if len(bc.NFTStake) >= 1 {
		nftre = bc.CheckNFTRewards()

	}
	bc.ClearTransactionPool()
	transactions := make([]*Transaction, 0)
	nb.megablock = StringTo32Byte(b.Header.Parents[0].ParentHashes[0])
	nb.metblock = StringTo32Byte(b.Header.Parents[0].ParentHashes[1])

	nb.currentHash = Hash([]byte(fmt.Sprintf("%x", nb.previousHash)), nb.timestamp, nb.nonce)

	MinBC := len(bc.MiniBlock)

	MegBC := len(bc.MegaBlock)
	MetBC := len(bc.MetBlock)

	var ts []*Transaction
	var Reward float64
	if MinBC <= 11 {
		bc.MiniBlock = append(bc.MiniBlock, fmt.Sprintf("%x", nb.currentHash))
		ts = append(transactions, NewTransactionMiner(MINING_SENDER, b.Header.UtxoCommitment, MINING_REWARD))
		Reward = MINING_REWARD
	} else if MinBC == 12 && MegBC < 5 {
		nb.megablock = nb.currentHash
		bc.MiniBlock = []string{}
		bc.MegaBlock = append(bc.MegaBlock, fmt.Sprintf("%x", nb.currentHash))

		ts = append(transactions, NewTransactionMiner(MINING_SENDER, b.Header.UtxoCommitment, MINING_REWARD_MEGA))
		Reward = MINING_REWARD_MEGA
	} else {
		nb.metblock = nb.currentHash
		bc.MegaBlock = []string{}
		bc.MiniBlock = []string{}
		bc.MetBlock = append(bc.MetBlock, fmt.Sprintf("%x", nb.currentHash))

		ts = append(transactions, NewTransactionMiner(MINING_SENDER, b.Header.UtxoCommitment, MINING_REWARD_MET))
		Reward = MINING_REWARD_MET
	}
	if MetBC >= 1 {
		bc.MetBlock = []string{}
	}

	for _, tx := range ntx {
		if tx.senderBlockchainAddress != mconfig.DeadWallet && (tx.senderBlockchainAddress != MINING_SENDER || tx.txtype == 3) {
			ts = append(ts, &Transaction{tx.senderBlockchainAddress, tx.recipientBlockchainAddress, tx.txtype, tx.value, tx.txhash, tx.timestamp, 1})
		}
	}

	for _, tx := range nftre {

		ts = append(ts, &Transaction{tx.senderBlockchainAddress, tx.recipientBlockchainAddress, tx.txtype, tx.value, tx.txhash, tx.timestamp, 1})

	}
	nb.transactions = ts
	nb.BlockToDB(db)
	bc.LastRPCBlock = bc.LastRPCBlock.BlockToRPCLastBlock(nb)

	if len(nfttx) >= 1 {

		bc.StakedNFT(nb, nfttx)

	}
	bc.WalletToDB(db, nb)

	return nb.currentHash, nil, Reward
}

func (lb *LastRPCBlock) BlockToRPCLastBlock(nb *MiniBlock) *LastRPCBlock {
	return &LastRPCBlock{
		Height:       nb.height,
		PreviousHash: nb.previousHash,
		Timestamp:    nb.timestamp,
		Bits:         nb.bits,
		Nonce:        nb.nonce,
		Megablock:    nb.megablock,
		Metblock:     nb.metblock,
		Transactions: nb.transactions,
		CurrentHash:  nb.currentHash,
	}

}
func (bc *Blockchain) LastMiniBlockRPC() error {
	lbk, lbv := domain.LastBlockRPC(bc.Db)
	genkey := ("bk-" + strings.Repeat("0", 64))
	nbk := fmt.Sprintf("%v", string(lbk))
	if nbk == genkey {
		n := domain.ReadTx(lbv)

		b := &LastRPCBlock{
			Height:       n.Height,
			PreviousHash: n.PreviousHash,
			Timestamp:    n.Timestamp,
		}
		bc.LastRPCBlock = b

	} else {
		block := new(MBlock)
		err := block.UnmarshalJSON(lbv)
		if err != nil {
			return err
		}
		b := &LastRPCBlock{
			Height:       block.Height,
			PreviousHash: block.PreviousHash,
			Timestamp:    block.Timestamp,
			Bits:         block.Bits,
			Nonce:        block.Nonce,
			Megablock:    block.Megablock,
			Metblock:     block.Metblock,
			Transactions: block.Transactions,
			CurrentHash:  block.CurrentHash,
		}
		bc.LastRPCBlock = b
	}

	return nil

}

func (B *MBlock) DomainblockToRPCBlock(val []byte) error {
	return B.UnmarshalJSON(val)
}

func StringTo32Byte(e string) [32]byte {
	var a [32]byte
	copy(a[:], []byte(e))
	return a
}
