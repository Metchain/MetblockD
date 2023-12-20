package stagingarea

import (
	"encoding/json"
	"fmt"
	"github.com/Metchain/MetblockD/db/database"
	"github.com/Metchain/MetblockD/external"
	pb "github.com/Metchain/MetblockD/protoserver/grpcserver/protowire"
	"github.com/Metchain/MetblockD/utils/logger"
	"google.golang.org/protobuf/proto"
	"math/big"
	"os"
	"strconv"
	"strings"
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
	Childblocks        []external.BlockLevelChildern
	Parents            []external.BlockLevelParents
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
		log.Criticalf("%v", err)

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

		Blockhash:    external.NewDomainHashFromByteArray(&b.CurrentHash),
		Blockheight:  int64(b.Height),
		Previoushash: external.NewDomainHashFromByteArray(&b.PreviousHash),

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

func (s *StagingArea) AddBlock(tempblock *external.TempBlock) (*external.TempBlock, error) {
	return s.addBlock(tempblock)
}

var Utxokey = database.MakeBucket([]byte("utxo_tx"))
var log = logger.RegisterSubSystem("MET- StagingArea")
var nftLayerkey = database.MakeBucket([]byte("layer2_nft"))
var utxo_walletkey = database.MakeBucket([]byte("utxo_wallet"))

func (s *StagingArea) addBlock(b *external.TempBlock) (*external.TempBlock, error) {
	db := s.databaseContext
	key := strconv.Itoa(int(b.Height))
	b.MarshalJSON()
	m, _ := json.Marshal(b)

	log.Infof("Block accepted. Current Block Height:", b.Height)

	batch, _ := db.Begin()
	batch.Put(blockkey.Key([]byte(key)), m)

	for _, tx := range b.Transactions {

		//tx := transaction.UnMarshalFor()
		//log.Println(transaction.senderBlockchainAddress)

		utxo := new(pb.UTXOResponse)
		utxo.Txhash = fmt.Sprintf("%x", tx.Txhash)
		utxo.Timestamp = fmt.Sprintf("%v", tx.Timestamp)
		utxo.Amount = fmt.Sprintf("%f", tx.Value)
		utxo.Fromwallet = tx.SenderBlockchainAddress
		utxo.Towallet = tx.RecipientBlockchainAddress
		// Update the block info
		utxo.Blockid = fmt.Sprintf("%v", b.Height)
		utxo.Txtype = fmt.Sprintf("%v", tx.Txtype)
		utxo.Blockhash = fmt.Sprintf("%x", b.CurrentHash)

		jsonUtxo, err := proto.Marshal(utxo)
		if err != nil {
			log.Criticalf("This shouldn't happen. Please check version")
			os.Exit(36)
		}
		key := fmt.Sprintf("%x", tx.Txhash)

		batch.Put(Utxokey.Key([]byte(key)), jsonUtxo)

		if tx.Txtype == 3 {
			nftpb := new(pb.NFTResponse)
			nftkey := fmt.Sprintf("%v", tx.Value)
			nftsd, err := db.Get(nftLayerkey.Key([]byte(nftkey)))
			if err != nil {
				log.Infof("Error Locating NFT :", nftkey, nftsd)
			}

			nftpb.NFTWallet = tx.RecipientBlockchainAddress
			nftpb.NFTSender = tx.SenderBlockchainAddress
			nftpb.NFTid = fmt.Sprintf("%v", tx.Value)
			nftpb.Txhash = fmt.Sprintf("%x", tx.Txhash)
			nftpb.Blockid = fmt.Sprintf("%v", b.Height)
			nftpb.Blockhash = fmt.Sprintf("%x", b.CurrentHash)
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
		return nil, err
	}
	return b, err

}

func UpdateWalletList(b *external.TempBlock, db database.Database, wallets map[string]*pb.UTXOWalletBalanceRespose) map[string]*pb.UTXOWalletBalanceRespose {

	for _, tx := range b.Transactions {

		sender := VerifyAddressPrefix(tx.SenderBlockchainAddress)
		reciver := VerifyAddressPrefix(tx.RecipientBlockchainAddress)

		// Get senders Wallet INFO

		// Get recivers Wallet INFO
		if wallets[sender] != nil {

			bfs, _ := strconv.ParseFloat(wallets[sender].Amount, 64)
			if (tx.Txtype == 1 || tx.Txtype == 0) && tx.Txstatus == 1 {
				wallets[sender].Amount = fmt.Sprintf("%.6f", SubBig(tx.Value, bfs))
			}
			if tx.Txtype == 3 && sender != "metchain:METCHAIN_Blockchain" {
				nftid, _ := strconv.ParseInt(fmt.Sprintf("%v", tx.Value), 10, 64)
				if containsint(wallets[reciver].NFT, nftid) {
					nfts := []int64{}
					for _, nft := range wallets[sender].NFT {
						if nftid != nft {
							nfts = append(nfts, nft)
						}

					}
				}

			}
			txinfo := &pb.Wallettx{Fromwallet: sender, Towallet: reciver, Value: fmt.Sprintf("%.6f", tx.Value), Txhash: fmt.Sprintf("%x", tx.Txhash), Timestamp: fmt.Sprintf("%v", tx.Timestamp), Txtype: fmt.Sprintf("%v", tx.Txtype)}
			var utxoinfo []*pb.Wallettx

			utxoinfo = append(utxoinfo, txinfo)
			counttx := len(wallets[sender].Wallettx)
			inlen := counttx
			if counttx >= 19 {
				inlen = 19
			}
			for _, sendertxval := range wallets[sender].Wallettx[:inlen] {
				utxoinfo = append(utxoinfo, sendertxval)

			}
			wallets[sender].Wallettx = utxoinfo

		} else {
			sdb, err := db.Get(utxo_walletkey.Key([]byte(sender)))
			if err == nil {
				sinfo := new(pb.UTXOWalletBalanceRespose)
				proto.Unmarshal(sdb, sinfo)
				wallets[sender] = &pb.UTXOWalletBalanceRespose{
					WalletAddress: sender,
					Amount:        sinfo.Amount,
					Wallettx:      sinfo.Wallettx,
					NFT:           sinfo.NFT,
				}
			}
			if wallets[sender] == nil {
				wallets[sender] = &pb.UTXOWalletBalanceRespose{
					WalletAddress: sender,
				}
			}

			if err == nil {
				sinfo := new(pb.UTXOWalletBalanceRespose)
				proto.Unmarshal(sdb, sinfo)
				wallets[sender] = &pb.UTXOWalletBalanceRespose{
					WalletAddress: sender,
					Amount:        sinfo.Amount,
					Wallettx:      sinfo.Wallettx,
					NFT:           sinfo.NFT,
				}
			}
			bfs, _ := strconv.ParseFloat(wallets[sender].Amount, 64)
			if (tx.Txtype == 1 || tx.Txtype == 0) && tx.Txstatus == 1 {
				wallets[sender].Amount = fmt.Sprintf("%.6f", SubBig(tx.Value, bfs))
			}
			if tx.Txtype == 3 {
				nftid, _ := strconv.ParseInt(fmt.Sprintf("%v", tx.Value), 10, 64)
				nfts := wallets[sender].NFT[:0]

				for _, i := range wallets[sender].NFT {
					if i != nftid {
						nfts = append(nfts, i)
					}
				}
				wallets[sender].NFT = nfts

			}
			txinfo := &pb.Wallettx{Fromwallet: sender, Towallet: reciver, Value: fmt.Sprintf("%.6f", tx.Value), Txhash: fmt.Sprintf("%x", tx.Txhash), Timestamp: fmt.Sprintf("%v", tx.Timestamp), Txtype: fmt.Sprintf("%v", tx.Txtype)}
			var utxoinfo []*pb.Wallettx
			utxoinfo = append(utxoinfo, txinfo)
			counttx := len(wallets[sender].Wallettx)
			inlen := counttx
			if counttx >= 19 {
				inlen = 19
			}
			for _, sendertxval := range wallets[sender].Wallettx[:inlen] {
				utxoinfo = append(utxoinfo, sendertxval)
			}
			wallets[sender].Wallettx = utxoinfo

		}
		if wallets[reciver] != nil {

			bfr, _ := strconv.ParseFloat(wallets[reciver].Amount, 64)
			if (tx.Txtype == 1 || tx.Txtype == 0) && tx.Txstatus == 1 {
				wallets[reciver].Amount = fmt.Sprintf("%.6f", AddBig(tx.Value, bfr))
			}
			if tx.Txtype == 3 {
				nftid, _ := strconv.ParseInt(fmt.Sprintf("%v", tx.Value), 10, 64)
				if !containsint(wallets[reciver].NFT, nftid) {

					wallets[reciver].NFT = append(wallets[reciver].NFT, nftid)
				}

			}
			rtxinfo := &pb.Wallettx{Fromwallet: sender, Towallet: reciver, Value: fmt.Sprintf("%.6f", tx.Value), Txhash: fmt.Sprintf("%x", tx.Txhash), Timestamp: fmt.Sprintf("%v", tx.Timestamp), Txtype: fmt.Sprintf("%v", tx.Txtype)}
			var rutxoinfo []*pb.Wallettx
			rutxoinfo = append(rutxoinfo, rtxinfo)
			counttx := len(wallets[reciver].Wallettx)
			inlen := counttx
			if counttx >= 19 {
				inlen = 19
			}
			for _, recivertxval := range wallets[reciver].Wallettx[:inlen] {

				rutxoinfo = append(rutxoinfo, recivertxval)
			}
			wallets[reciver].Wallettx = rutxoinfo
		} else {
			if !strings.Contains(tx.RecipientBlockchainAddress, "metchain:") {
				tx.RecipientBlockchainAddress = "metchain:" + tx.RecipientBlockchainAddress
			}
			rdb, err := db.Get(utxo_walletkey.Key([]byte(tx.RecipientBlockchainAddress)))

			if err == nil {
				rinfo := new(pb.UTXOWalletBalanceRespose)
				proto.Unmarshal(rdb, rinfo)
				wallets[reciver] = &pb.UTXOWalletBalanceRespose{
					WalletAddress: reciver,
					Amount:        rinfo.Amount,
					Wallettx:      rinfo.Wallettx,
					NFT:           rinfo.NFT,
				}
			}
			if wallets[reciver] == nil {
				wallets[reciver] = &pb.UTXOWalletBalanceRespose{
					WalletAddress: reciver,
				}
			}
			bfr, _ := strconv.ParseFloat(wallets[reciver].Amount, 64)
			if (tx.Txtype == 1 || tx.Txtype == 0) && tx.Txstatus == 1 {
				wallets[reciver].Amount = fmt.Sprintf("%.6f", AddBig(tx.Value, bfr))
			}
			if tx.Txtype == 3 {
				nftid, _ := strconv.ParseInt(fmt.Sprintf("%v", tx.Value), 10, 64)
				if !containsint(wallets[reciver].NFT, nftid) {

					wallets[reciver].NFT = append(wallets[reciver].NFT, nftid)
				}

			}
			rtxinfo := &pb.Wallettx{Fromwallet: sender, Towallet: reciver, Value: fmt.Sprintf("%.6f", tx.Value), Txhash: fmt.Sprintf("%x", tx.Txhash), Timestamp: fmt.Sprintf("%v", tx.Timestamp), Txtype: fmt.Sprintf("%v", tx.Txtype)}
			var rutxoinfo []*pb.Wallettx
			rutxoinfo = append(rutxoinfo, rtxinfo)
			counttx := len(wallets[reciver].Wallettx)
			inlen := counttx
			if counttx >= 19 {
				inlen = 19
			}
			for _, recivertxval := range wallets[reciver].Wallettx[:inlen] {
				rutxoinfo = append(rutxoinfo, recivertxval)
			}
			wallets[reciver].Wallettx = rutxoinfo
		}
	}
	return wallets
}

func containsint(s []int64, str int64) bool {
	for _, val := range s {

		if val == str {
			return true
		}
	}

	return false
}

func AddBig(x float32, y float64) *big.Float {
	z := new(big.Float)
	num := float64(x)
	newnum := new(big.Float)
	old := new(big.Float)

	newnum.SetFloat64(y)
	old.SetFloat64(num)
	z.Add(old, newnum)
	return z

}

func SubBig(x float32, y float64) *big.Float {
	z := new(big.Float)
	num, _ := strconv.ParseFloat(fmt.Sprintf("%v", x), 64)
	newnum := new(big.Float)
	old := new(big.Float)

	newnum.SetFloat64(y)
	old.SetFloat64(num)
	z.Sub(newnum, old)
	return z

}

func VerifyAddressPrefix(str string) string {
	if !strings.Contains(str, "metchain:") {
		str = "metchain:" + str

	}
	return str
}
