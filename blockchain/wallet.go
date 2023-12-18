package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/Metchain/Metblock/db/database"

	"math/big"
	"strconv"
	"strings"

	pb "github.com/Metchain/Metblock/protoserver/grpcserver/protowire"
	"google.golang.org/protobuf/proto"
)

func (bc *Blockchain) AddWallet(wallet string) bool {

	bc.Wallets = append(bc.Wallets, &WalletCreated{WalletAddress: wallet})
	log.Infof(wallet)
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

func (bc *Blockchain) WalletToDB(db database.Database, nb *MiniBlock) error {

	_, value := LastWallet(db)
	vwall := new(pb.VWalletResponse)
	proto.Unmarshal(value, vwall)
	var key *database.Key
	i := 1

	if vwall != nil {

		i = int(vwall.Walletid) + 1

	}
	batch, _ := db.Begin()
	defer batch.RollbackUnlessClosed()
	for _, vw := range bc.Wallets {
		key = verifiedWalletidKey.Key([]byte(strconv.Itoa(i)))

		wkey := utxo_walletkey.Key([]byte(vw.WalletAddress))
		utxo := new(pb.VWalletResponse)
		utxo.Walletid = uint64(i)
		utxo.WalletAddress = vw.WalletAddress
		utxo.Blockhash = fmt.Sprintf("%x", nb.currentHash)
		utxo.Lockhash = fmt.Sprintf("%x", Hash([]byte(fmt.Sprintf("%x", vw.WalletAddress)), nb.timestamp, nb.nonce))
		m, _ := proto.Marshal(utxo)

		batch.Put(key, m)
		batch.Put(wkey, m)
		i = i + 1
	}
	err := batch.Commit()
	if err != nil {
		log.Infof("Error: ", err)
		return err
	}
	bc.Wallets = bc.Wallets[:0]
	return nil
}

var verifiedWalletidKey = database.MakeBucket([]byte("verified_wallets_by_id"))

func LastWallet(db database.Database) ([]byte, []byte) {

	key := []byte{}
	value := []byte{}

	cursor, err := db.Cursor(verifiedWalletidKey)
	if err != nil {
		log.Error(err)
	}

	for ok := cursor.Last(); ok; ok = cursor.Next() {
		dbkey, _ := cursor.Key()
		value, _ = cursor.Value()
		key = dbkey.Bytes()

	}

	utxo := new(pb.VWalletResponse)
	proto.Unmarshal(value, utxo)

	return key, value
}

func UpdateWalletList(b *MiniBlock, db database.Database, wallets map[string]*pb.UTXOWalletBalanceRespose) map[string]*pb.UTXOWalletBalanceRespose {

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
			rdb, err := db.Get(utxo_walletkey.Key([]byte(tx.recipientBlockchainAddress)))

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
