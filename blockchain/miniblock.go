package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Metchain/Metblock/domain"
	"github.com/Metchain/Metblock/heavyhash"
	"github.com/Metchain/Metblock/mconfig"
	pb "github.com/Metchain/Metblock/proto"
	"github.com/btcsuite/goleveldb/leveldb"
	"google.golang.org/protobuf/proto"
	"log"
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

func (b *MiniBlock) BlockToDB(db *leveldb.DB) error {
	kln := 64 - len(strconv.Itoa(int(b.height)))
	key := "bk-" + strings.Repeat("0", kln) + strconv.Itoa(int(b.height))
	b.MarshalJSON()
	m, _ := json.Marshal(b)
	log.Printf("Block accepted. Current Block Height:", b.height)
	batch := new(leveldb.Batch)
	batch.Put([]byte(key), m)

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
			log.Println("This shouldn't happen. Please check version")
			os.Exit(36)
		}
		key := fmt.Sprintf("%x", tx.txhash)

		batch.Put([]byte(key), jsonUtxo)

		//log.Println("Added Transactions to UTXO DATABASE:", tx.txhash, " :: ", fmt.Sprintf("%s", jsonUtxo))

		if tx.txtype == 3 {
			nftpb := new(pb.NFTResponse)
			nftkey := "NFT-" + fmt.Sprintf("%v", tx.value)
			nftsd, err := db.Get([]byte(nftkey), nil)
			if err != nil {
				log.Println("Error Locating NFT :", nftkey, nftsd)
			}

			nftpb.NFTWallet = tx.recipientBlockchainAddress
			nftpb.NFTSender = tx.senderBlockchainAddress
			nftpb.NFTid = fmt.Sprintf("%v", tx.value)
			nftpb.Txhash = fmt.Sprintf("%x", tx.txhash)
			nftpb.Blockid = fmt.Sprintf("%v", b.height)
			nftpb.Blockhash = fmt.Sprintf("%x", b.currentHash)
			jsonnft, err := proto.Marshal(nftpb)
			if err != nil {
				log.Println("This should happen at all. :", err)
				os.Exit(123456789)
			}
			batch.Put([]byte(nftkey), jsonnft)
		}

		// Subtract balance from senders wallet.

		//Adding it to wallet db

	}
	wallets := make(map[string]*pb.UTXOWalletBalanceRespose, 0)
	wallets = UpdateWalletList(b, db, wallets)
	for tempkey, val := range wallets {
		jsonWalletSender, err := proto.Marshal(val)
		if err != nil {
			log.Println("Error. This shouldn't happen")
			os.Exit(8888)
		}
		batch.Put([]byte(tempkey), jsonWalletSender)
		log.Printf("Added Wallet: " + val.WalletAddress + "\n Balance: " + val.Amount + "\n\n")
	}
	err := db.Write(batch, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	return nil
}

func (b *MiniBlock) Print() {
	log.Printf("Height 				%d\n", b.height)
	log.Printf("Timestamp			%d\n", b.timestamp)
	log.Printf("Nonce				%d\n", b.nonce)
	log.Printf("Megablock				%d\n", b.nonce)
	log.Printf("Metblock				%d\n", b.nonce)
	log.Printf("Prev_Hash			%x\n", b.previousHash)
	log.Printf("Current_Hash          %x\n", b.currentHash)
	//fmt.Printf("Transactions			%s\n", b.Transactions)
	for _, t := range b.transactions {
		t.Print()
	}
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

func LastMiniBlock(mc *domain.Metchain) *MBlock {
	lbk, lbv := mc.LastBlock()
	genkey := ("bk-" + strings.Repeat("0", 64))
	nbk := fmt.Sprintf("%v", string(lbk))
	block := new(MBlock)
	if nbk == genkey {
		n := domain.ReadTx(lbv)

		block.Timestamp = n.Timestamp
		block.PreviousHash = n.PreviousHash
		block.Nonce = n.Nonce
		block.Height = n.Height
		block.Bits = 539658475
		if n.Height == 0 {
			g := string(n.Transaction)
			vt := new(mconfig.Txtransaction)

			json.Unmarshal([]byte(g), vt)
		}
	} else {
		block.UnmarshalJSON(lbv)

	}

	mc.Cdiff = block.Bits
	mc.Blockheight = block.Height
	//log.Println(block)
	return block
}

func CreateMiniBlock(b *pb.RpcBlock, db *leveldb.DB, bc *Blockchain) ([32]byte, error) {
	Temptimestamp := time.Now().UnixMilli()
	Tempoldntime := bc.LastBlockTime + 10000
	if bc.LastBlockTime > Temptimestamp || Tempoldntime >= Temptimestamp {
		//log.Printf("This is trigered")
		return [32]byte{}, errors.New("Block is too early")
	}

	bc.Mux.Lock()
	defer bc.Mux.Unlock()

	nb := new(MiniBlock)

	var oldtimestamp int64
	nb.bits = uint64(b.Header.Bits)
	//log.Println(bc.MiniBlock)

	nb.nonce = b.Header.Nonce
	nb.height = bc.LastBlockHeight
	nb.previousHash = bc.LastBlockHash
	oldtimestamp = bc.LastBlockTime

	nb.timestamp = time.Now().UnixMilli()

	oldntime := oldtimestamp + 10000

	if oldtimestamp > nb.timestamp || oldntime >= nb.timestamp {
		//log.Printf("This is trigered")
		nb.bits = nb.bits + 1
		return [32]byte{}, errors.New("Block is too early")
	} else if nb.timestamp-oldntime > 2000 && nb.bits > 569658475 {
		nb.bits = nb.bits - 1
	} else {
		nb.bits = nb.bits + 1
	}

	ntx := bc.Txpool
	nfttx := bc.NFTPool

	nftre := []*Transaction{}
	if len(bc.NFTStake) >= 1 {
		log.Println("Entered Staking")
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

	if MinBC <= 11 {
		bc.MiniBlock = append(bc.MiniBlock, fmt.Sprintf("%x", nb.currentHash))
		ts = append(transactions, NewTransactionMiner(MINING_SENDER, b.Header.UtxoCommitment, MINING_REWARD))
	} else if MinBC == 12 && MegBC < 5 {
		nb.megablock = nb.currentHash
		bc.MiniBlock = []string{}
		bc.MegaBlock = append(bc.MegaBlock, fmt.Sprintf("%x", nb.currentHash))

		ts = append(transactions, NewTransactionMiner(MINING_SENDER, b.Header.UtxoCommitment, MINING_REWARD_MEGA))
	} else {
		nb.metblock = nb.currentHash
		bc.MegaBlock = []string{}
		bc.MiniBlock = []string{}
		bc.MetBlock = append(bc.MetBlock, fmt.Sprintf("%x", nb.currentHash))

		ts = append(transactions, NewTransactionMiner(MINING_SENDER, b.Header.UtxoCommitment, MINING_REWARD_MET))
	}
	if MetBC >= 1 {
		bc.MetBlock = []string{}
	}

	for _, tx := range ntx {

		ts = append(ts, &Transaction{tx.senderBlockchainAddress, tx.recipientBlockchainAddress, tx.txtype, tx.value, tx.txhash, tx.timestamp, 1})

	}

	for _, tx := range nftre {

		ts = append(ts, &Transaction{tx.senderBlockchainAddress, tx.recipientBlockchainAddress, tx.txtype, tx.value, tx.txhash, tx.timestamp, 1})

	}
	nb.transactions = ts
	nb.BlockToDB(db)

	bc.LastBlockTime = nb.timestamp
	bc.LastBlockHeight = nb.height + 1
	bc.LastBlockHash = nb.currentHash

	if len(nfttx) >= 1 {

		bc.StakedNFT(nb, nfttx)

	}
	bc.WalletToDB(db, nb)

	return nb.currentHash, nil
}

func LastMiniBlockRPC(db *leveldb.DB) (uint64, [32]byte, int64) {
	lbk, lbv := domain.LastBlockRPC(db)
	genkey := ("bk-" + strings.Repeat("0", 64))
	nbk := fmt.Sprintf("%v", string(lbk))
	if nbk == genkey {
		n := domain.ReadTx(lbv)
		return n.Height + 1, n.PreviousHash, n.Timestamp

	} else {
		block := new(MBlock)
		block.UnmarshalJSON(lbv)
		return block.Height + 1, block.CurrentHash, block.Timestamp
	}

}

func StringTo32Byte(e string) [32]byte {
	var a [32]byte
	copy(a[:], []byte(e))
	return a
}
