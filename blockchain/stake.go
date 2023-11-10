package blockchain

import (
	"fmt"
	pb "github.com/Metchain/Metblock/proto"
	"github.com/btcsuite/goleveldb/leveldb"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"strconv"
	"time"
)

type NFTStaking struct {
	NFT           int
	MetCoin       float64
	WalletAddress string
	LockTime      int64
	UnLockTime    int64
	Status        bool
	LockingTXHash []byte
}

func (bc *Blockchain) StakeNFTMet(sender string, seed string, value string, nft string, stakeunlock string) bool {
	if bc.verifyNFt(sender, nft) {
		var dur string
		if stakeunlock == "3" {
			dur = "131400"
		} else if stakeunlock == "6" {
			dur = "262800"
		} else if stakeunlock == "9" {
			dur = "350400"
		} else if stakeunlock == "12" {
			dur = "525600"
		}
		Unlocktime, _ := time.ParseDuration(dur + "m")
		Stake := new(NFTStaking)

		f, _ := strconv.ParseFloat(value, 64)
		Stake.MetCoin = f
		Stake.WalletAddress = sender
		Stake.LockTime = time.Now().UnixMilli()
		Stake.UnLockTime = time.Now().Add(Unlocktime).UnixMilli()
		NFTid, _ := strconv.ParseInt(nft, 10, 64)
		Stake.NFT = int(NFTid)
		bc.NFTPool = append(bc.NFTPool, Stake)

		return true
	}
	return false
}

func (bc *Blockchain) StakedNFT(b *MiniBlock, pool []*NFTStaking) []*Transaction {
	batch := new(leveldb.Batch)
	txs := []*Transaction{}
	for _, z := range pool {
		sn := new(pb.StakeNFT)
		nftid := fmt.Sprintf("%v", z.NFT)
		t := NewTransaction(z.WalletAddress, STAKING_SENDER+"( MET NFT #"+nftid+")", float32(z.MetCoin))

		tempkey := "staked-" + fmt.Sprintf("%v", z.NFT)
		bc.Txpool = append(bc.Txpool, t)
		sn.Txhash = fmt.Sprintf("%x", t.txhash)
		sn.Blockhash = fmt.Sprintf("%x", b.currentHash)
		sn.LockTime = fmt.Sprintf("%v", z.LockTime)
		sn.UnlockTime = fmt.Sprintf("%v", z.UnLockTime)
		sn.Blockid = fmt.Sprintf("%v", b.height)
		sn.NFTid = fmt.Sprintf("%v", z.NFT)
		sn.NFTSender = fmt.Sprintf("%v", z.WalletAddress)
		sn.StakeAmount = fmt.Sprintf("%v", z.MetCoin)
		log.Println(sn.StakeAmount)
		staker, err := proto.Marshal(sn)
		txs = append(txs, t)
		if err != nil {
			log.Println("Error. This shouldn't happen")
			os.Exit(8888)
		}
		batch.Put([]byte(tempkey), staker)
		bc.NFTStake = append(bc.NFTStake, sn)
	}
	db := bc.Dbcon
	err := db.Write(batch, nil)

	if err != nil {
		fmt.Println("Error: ", err)

	}
	return txs
}

func (bc *Blockchain) verifyNFt(sender string, nft string) bool {
	return true
}
func (bc *Blockchain) GetStakedNFT() []*pb.StakeNFT {
	i := 1
	db := bc.Dbcon
	Staked := []*pb.StakeNFT{}
	for i <= 500 {
		kp := "staked-" + fmt.Sprintf("%v", i)
		m, err := db.Get([]byte(kp), nil)
		if err != nil {
			//log.Printf("NFT %v is not staked", i)

		} else {
			dbnft := new(pb.StakeNFT)
			proto.Unmarshal(m, dbnft)
			Staked = append(Staked, dbnft)

		}

		i = i + 1
	}
	log.Println("Checking for Staked NFT's")

	return Staked
}

func CheckNFT(nv []int, nft string) bool {
	nftid, _ := strconv.ParseInt(nft, 10, 64)
	for _, id := range nv {
		if id == int(nftid) {
			return true
		}
	}
	return false
}
func (bc *Blockchain) CheckNFTRewards() []*Transaction {
	ntxpool := []*Transaction{}
	nfts := bc.GetNFT()
	for _, s := range bc.NFTStake {

		unlock, _ := strconv.ParseInt(s.UnlockTime, 10, 64)
		lock, _ := strconv.ParseInt(s.LockTime, 10, 64)

		locktime := (unlock - lock) / (60 * 1000)
		tn := time.Now().UnixMilli()
		if tn > unlock {

			var reward float32
			if CheckNFT(nfts.Superrare, s.NFTid) {
				if locktime == 131400 {

					reward = bc.calculatereward(2, s.StakeAmount)
				} else if locktime == 262800 {
					reward = bc.calculatereward(5, s.StakeAmount)
				} else if locktime == 350400 {
					reward = bc.calculatereward(8, s.StakeAmount)
				} else if locktime == 525600 {
					reward = bc.calculatereward(11, s.StakeAmount)
				} else {
					log.Printf("Unexpect error in %v NFT \n", s.NFTid)
					return nil
				}
			} else if CheckNFT(nfts.Rare, s.NFTid) {
				if locktime == 131400 {

					reward = bc.calculatereward(1.7, s.StakeAmount)
				} else if locktime == 262800 {
					reward = bc.calculatereward(4.25, s.StakeAmount)
				} else if locktime == 350400 {
					reward = bc.calculatereward(6.8, s.StakeAmount)
				} else if locktime == 525600 {
					reward = bc.calculatereward(9.35, s.StakeAmount)
				} else {
					log.Printf("Unexpect error in %v NFT \n", s.NFTid)
					return nil
				}
			} else if CheckNFT(nfts.LessCommon, s.NFTid) {
				if locktime == 131400 {

					reward = bc.calculatereward(1.3, s.StakeAmount)
				} else if locktime == 262800 {
					reward = bc.calculatereward(3.25, s.StakeAmount)
				} else if locktime == 350400 {
					reward = bc.calculatereward(5.2, s.StakeAmount)
				} else if locktime == 525600 {
					reward = bc.calculatereward(7.15, s.StakeAmount)
				} else {
					log.Printf("Unexpect error in %v NFT \n", s.NFTid)
					return nil
				}
			} else if CheckNFT(nfts.Common, s.NFTid) {
				if locktime == 131400 {

					reward = bc.calculatereward(1, s.StakeAmount)
				} else if locktime == 262800 {
					reward = bc.calculatereward(2.5, s.StakeAmount)
				} else if locktime == 350400 {
					reward = bc.calculatereward(4, s.StakeAmount)
				} else if locktime == 525600 {
					reward = bc.calculatereward(5.5, s.StakeAmount)
				} else {
					log.Printf("Unexpect error in %v NFT \n", s.NFTid)
					return nil
				}
			}

			tx := NewTransaction(MINING_SENDER+"( MET NFT #"+s.NFTid+")", s.NFTSender, reward)
			ntxpool = append(ntxpool, tx)
			am, _ := strconv.ParseFloat(s.StakeAmount, 64)
			oldtx := NewTransaction(STAKING_SENDER+"( MET NFT #"+s.NFTid+")", s.NFTSender, float32(am))
			ntxpool = append(ntxpool, oldtx)

			log.Printf("Clear NFT %v \n", s.NFTid)
			kp := "staked-" + fmt.Sprintf("%v", s.NFTid)
			bc.Dbcon.Delete([]byte(kp), nil)

		}

	}
	bc.NFTStake = bc.GetStakedNFT()
	return ntxpool

}

func (bc *Blockchain) calculatereward(lt float32, met string) float32 {
	fm, _ := strconv.ParseFloat(met, 64)
	return ((lt * float32(fm)) / 100)
}

func (bc *Blockchain) ClearStake(ls []*int64) {
	newpool := []*pb.StakeNFT{}

	for _, nft := range bc.NFTStake {

		for _, v := range ls {
			if nft.NFTid != fmt.Sprintf("%v", *v) {
				newpool = append(newpool, nft)
			}
		}
	}
	if len(ls) >= 1 {
		bc.NFTStake = newpool
	}
	for _, v := range ls {
		log.Printf("Clear NFT %v \n", *v)
		kp := "staked-" + fmt.Sprintf("%v", *v)
		bc.Dbcon.Delete([]byte(kp), nil)
	}
}
