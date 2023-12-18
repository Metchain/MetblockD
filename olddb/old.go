package main

import (
	"github.com/Metchain/Metblock/external"
	"github.com/Metchain/Metblock/utils/pow"
	"log"
	"math"
	"os"
	"time"
)

/*
	func main() {
		dbdir := mconfig.GetDatadir() + "/data"
		newdir := infraconfig.DefaultAppDir + "/metchain-mainnet/datadir"

		olddb, err := leveldb.OpenFile(dbdir, nil)
		if err != nil {
			panic(err)
		}
		newdb, err := ldb.NewLevelDB(newdir, 256)
		if err != nil {
			panic(err)
		}
		log.Println(dbdir)
		log.Println(newdir)
		log.Println("Starting")
		// Sync Blocks
		blockcount := sycnBlocks(olddb, newdb, "block", "bk-")
		vwalletsbyhashcount := sycnBlocks(olddb, newdb, "verified_wallets", "vwallethash-")
		vwalletsbyidcount := sycnBlocks(olddb, newdb, "verified_wallets_by_id", "vwallet-")
		layer2nftcount := sycnBlocks(olddb, newdb, "layer2_nft", "NFT-")
		layer2nftstaked := sycnBlocks(olddb, newdb, "layer2_nft_staked", "staked-")
		uxtwalletcount := sycnBlocks(olddb, newdb, "utxo_wallet", "metchain:")
		log.Printf("Blocks updated: %v \n", blockcount)
		log.Printf("Verifed wallets by hash updated: %v \n", vwalletsbyhashcount)
		log.Printf("Verified wallets by id updated: %v \n", vwalletsbyidcount)
		log.Printf("Layer 2 NFTs updated: %v \n", layer2nftcount)
		log.Printf("Layer 2 NFT Staked updated: %v \n", layer2nftstaked)
		log.Printf("Utxo Wallets updated: %v \n", uxtwalletcount)
	}

	func sycnBlocks(olddb *leveldb.DB, newdb database.Database, bucket string, prefix string) int {
		iter := olddb.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
		infokey := database.MakeBucket([]byte(bucket))
		i := 0
		for iter.Next() {
			i++
			val := iter.Value()
			key := iter.Key()

			keystring := []byte(strings.Replace(string(key), prefix, "", 1))
			if prefix == "bk-" {
				st, _ := strconv.ParseInt(string(keystring), 10, 64)
				keystring = []byte(fmt.Sprintf("%v", st))

			}
			if prefix == "metchain:" {
				keystring = key
			}
			log.Println(string(key))

			newdb.Put(infokey.Key(keystring), val)
		}

		iter.Release()

		return i
	}
*/
var hash = "0b407c525c49e7269ac0116238efaf447d340d65d009161f450375a1ab351aad"

const eps float64 = 1e-9

type matrix [64][64]uint16

func main() {
	h, err := external.NewDomainHashFromByteSlice([]byte(hash[:32]))
	if err != nil {
		panic(err)
	}
	var mat matrix
	generator := pow.NewxoShiRo256PlusPlus(h)
	log.Printf("Generator", generator)
	ct := time.Now()
	kl := 0
	jk := 4
	valid := 0
	invalid := 0
	for {
		for i := range mat {
			for j := 0; j < 64; j += 16 {
				val := generator.Uint64()
				for shift := 0; shift < 16; shift++ {
					mat[i][j+shift] = uint16(val >> (jk * shift) & 0x0F)

				}

			}
		}
		kl++
		if mat.computeRank() == 64 {
			valid++
			//log.Println(mat)

		} else {
			invalid++
		}
		if time.Since(ct).Seconds() >= 6 {
			log.Printf("Time: %v        Hashes: %v             Shift Val X: %v              valid : %v \\ %v", time.Since(ct).Seconds(), kl, jk, valid, invalid)
			os.Exit(5)
		}
	}

}

func (mat *matrix) computeRank() int {
	var B [64][64]float64
	for i := range B {
		for j := range B[0] {
			B[i][j] = float64(mat[i][j])
		}
	}
	var rank int
	var rowSelected [64]bool
	for i := 0; i < 64; i++ {
		var j int
		for j = 0; j < 64; j++ {
			if !rowSelected[j] && math.Abs(B[j][i]) > eps {
				break
			}
		}
		if j != 64 {
			rank++
			rowSelected[j] = true
			for p := i + 1; p < 64; p++ {
				B[j][p] /= B[j][i]
			}
			for k := 0; k < 64; k++ {
				if k != j && math.Abs(B[k][i]) > eps {
					for p := i + 1; p < 64; p++ {
						B[k][p] -= B[j][p] * B[k][i]
					}
				}
			}
		}
	}
	return rank
}
