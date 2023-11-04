package blockchain

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	pb "github.com/Metchain/Metblock/proto"
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/util"
	"google.golang.org/protobuf/proto"
)

func (bc *Blockchain) AddWallet(wallet string) bool {

	bc.Wallets = append(bc.Wallets, &WalletCreated{WalletAddress: wallet})
	log.Println(wallet)
	return true
}

func (wc *WalletCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		WalletAddress string `json:"walletAddress"`
		BlockHash     string `json:"blockHash"`
		LockHash      string `json:"lockHash"`
	}{
		WalletAddress: wc.WalletAddress,
		BlockHash:     wc.BlockHash,
		LockHash:      wc.LockHash,
	})
}

func (bc *Blockchain) WalletToDB(db *leveldb.DB, nb *MiniBlock) error {

	_, value := LastWallet(db)
	vwall := new(pb.VWalletResponse)
	proto.Unmarshal(value, vwall)
	var key string
	i := 1
	kln := 63
	if vwall != nil {
		wid := int(vwall.Walletid) + 1
		i = int(vwall.Walletid) + 1
		kln = 64 - len(strconv.Itoa(wid))

	}
	batch := new(leveldb.Batch)
	for _, vw := range bc.Wallets {
		key = "vwallet-" + strings.Repeat("0", kln) + strconv.Itoa(i)
		wkey := "vwallethash-" + vw.WalletAddress
		utxo := new(pb.VWalletResponse)
		utxo.Walletid = uint64(i)
		utxo.WalletAddress = vw.WalletAddress
		utxo.Blockhash = fmt.Sprintf("%x", nb.currentHash)
		utxo.Lockhash = fmt.Sprintf("%x", Hash([]byte(fmt.Sprintf("%x", vw.WalletAddress)), nb.timestamp, nb.nonce))
		m, _ := proto.Marshal(utxo)

		batch.Put([]byte(key), m)
		batch.Put([]byte(wkey), m)
		i = i + 1
	}
	err := db.Write(batch, nil)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	bc.Wallets = bc.Wallets[:0]
	return nil
}

func LastWallet(db *leveldb.DB) ([]byte, []byte) {

	key := []byte{}
	value := []byte{}
	iter := db.NewIterator(util.BytesPrefix([]byte("vwallet-")), nil)
	/**/

	ok := iter.Last()
	if ok {
		key = iter.Key()
		value = iter.Value()

	}
	ok = iter.Next()
	if ok {
		key = iter.Key()
		value = iter.Value()

	}

	iter.Release()
	utxo := new(pb.VWalletResponse)
	proto.Unmarshal(value, utxo)

	log.Printf("%v", string(key))
	//os.Exit(1000)
	return key, value
}

func UpdateWalletList(b *MiniBlock, db *leveldb.DB, wallets map[string]*pb.UTXOWalletBalanceRespose) map[string]*pb.UTXOWalletBalanceRespose {

	for _, tx := range b.transactions {

		sender := VerifyAddressPrefix(tx.senderBlockchainAddress)
		reciver := VerifyAddressPrefix(tx.recipientBlockchainAddress)

		// Get senders Wallet INFO

		// Get recivers Wallet INFO
		if wallets[sender] != nil {

			bfs, _ := strconv.ParseFloat(wallets[sender].Amount, 64)
			if (tx.txtype == 1 || tx.txtype == 0) && tx.txstatus == 1 {
				wallets[sender].Amount = fmt.Sprintf("%.6f", SubBig(tx.value, bfs))
			}
			if tx.txtype == 3 && sender != "metchain:METCHAIN_Blockchain" {
				nftid, _ := strconv.ParseInt(fmt.Sprintf("%v", tx.value), 10, 64)
				if containsint(wallets[reciver].NFT, nftid) {
					nfts := []int64{}
					for _, nft := range wallets[sender].NFT {
						if nftid != nft {
							nfts = append(nfts, nft)
						}

					}
				}

			}
			txinfo := &pb.Wallettx{Fromwallet: sender, Towallet: reciver, Value: fmt.Sprintf("%.6f", tx.value), Txhash: fmt.Sprintf("%x", tx.txhash), Timestamp: fmt.Sprintf("%v", tx.timestamp), Txtype: fmt.Sprintf("%v", tx.txtype)}
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
			sdb, err := db.Get([]byte(sender), nil)
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
			if (tx.txtype == 1 || tx.txtype == 0) && tx.txstatus == 1 {
				wallets[sender].Amount = fmt.Sprintf("%.6f", SubBig(tx.value, bfs))
			}
			if tx.txtype == 3 {
				nftid, _ := strconv.ParseInt(fmt.Sprintf("%v", tx.value), 10, 64)
				nfts := wallets[sender].NFT[:0]

				for _, i := range wallets[sender].NFT {
					if i != nftid {
						nfts = append(nfts, i)
					}
				}
				wallets[sender].NFT = nfts

			}
			txinfo := &pb.Wallettx{Fromwallet: sender, Towallet: reciver, Value: fmt.Sprintf("%.6f", tx.value), Txhash: fmt.Sprintf("%x", tx.txhash), Timestamp: fmt.Sprintf("%v", tx.timestamp), Txtype: fmt.Sprintf("%v", tx.txtype)}
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
			if (tx.txtype == 1 || tx.txtype == 0) && tx.txstatus == 1 {
				wallets[reciver].Amount = fmt.Sprintf("%.6f", AddBig(tx.value, bfr))
			}
			if tx.txtype == 3 {
				nftid, _ := strconv.ParseInt(fmt.Sprintf("%v", tx.value), 10, 64)
				if !containsint(wallets[reciver].NFT, nftid) {

					wallets[reciver].NFT = append(wallets[reciver].NFT, nftid)
				}

			}
			rtxinfo := &pb.Wallettx{Fromwallet: sender, Towallet: reciver, Value: fmt.Sprintf("%.6f", tx.value), Txhash: fmt.Sprintf("%x", tx.txhash), Timestamp: fmt.Sprintf("%v", tx.timestamp), Txtype: fmt.Sprintf("%v", tx.txtype)}
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
			if !strings.Contains(tx.recipientBlockchainAddress, "metchain:") {
				tx.recipientBlockchainAddress = "metchain:" + tx.recipientBlockchainAddress
			}
			rdb, err := db.Get([]byte(tx.recipientBlockchainAddress), nil)

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
			if (tx.txtype == 1 || tx.txtype == 0) && tx.txstatus == 1 {
				wallets[reciver].Amount = fmt.Sprintf("%.6f", AddBig(tx.value, bfr))
			}
			if tx.txtype == 3 {
				nftid, _ := strconv.ParseInt(fmt.Sprintf("%v", tx.value), 10, 64)
				if !containsint(wallets[reciver].NFT, nftid) {

					wallets[reciver].NFT = append(wallets[reciver].NFT, nftid)
				}

			}
			rtxinfo := &pb.Wallettx{Fromwallet: sender, Towallet: reciver, Value: fmt.Sprintf("%.6f", tx.value), Txhash: fmt.Sprintf("%x", tx.txhash), Timestamp: fmt.Sprintf("%v", tx.timestamp), Txtype: fmt.Sprintf("%v", tx.txtype)}
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

func VerifyAddressPrefix(str string) string {
	if !strings.Contains(str, "metchain:") {
		str = "metchain:" + str

	}
	return str
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

func containsint(s []int64, str int64) bool {
	for _, val := range s {

		if val == str {
			return true
		}
	}

	return false
}
